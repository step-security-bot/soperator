ARG BASE_IMAGE=ubuntu:focal

FROM $BASE_IMAGE AS controller_slurmctld

ARG SLURM_VERSION=23.11.6

ARG DEBIAN_FRONTEND=noninteractive

# TODO: Install only those dependencies that are required for running slurmctld + useful utilities
# Install dependencies
RUN apt-get update && \
    apt -y install \
        wget \
        curl \
        git \
        build-essential \
        bc \
        python3  \
        autoconf \
        pkg-config \
        libssl-dev \
        libpam0g-dev \
        libtool \
        libjansson-dev \
        libjson-c-dev \
        libmunge-dev \
        libhwloc-dev \
        liblz4-dev \
        flex \
        libevent-dev \
        jq \
        squashfs-tools \
        zstd \
        software-properties-common \
        iputils-ping \
        dnsutils \
        telnet \
        strace \
        vim \
        tree \
        lsof \
        daemontools

# Install PMIx
COPY common/scripts/install_pmix.sh /opt/bin/
RUN chmod +x /opt/bin/install_pmix.sh && \
    /opt/bin/install_pmix.sh && \
    rm /opt/bin/install_pmix.sh

# TODO: Install only necessary packages
# Download and install Slurm packages
RUN for pkg in slurm-smd-client slurm-smd-dev slurm-smd-libnss-slurm slurm-smd-libpmi0 slurm-smd-libpmi2-0 slurm-smd-libslurm-perl slurm-smd-slurmctld slurm-smd; do \
        wget -P /tmp https://github.com/nebius/slurm-deb-packages/releases/download/v$SLURM_VERSION/${pkg}_$SLURM_VERSION-1_amd64.deb; \
    done

RUN apt install -y /tmp/*.deb && rm -rf /tmp/*.deb

# Install slurm plugins
COPY common/chroot-plugin/chroot.c /usr/src/chroot-plugin/
COPY common/scripts/install_slurm_plugins.sh /opt/bin/
RUN chmod +x /opt/bin/install_slurm_plugins.sh && \
    /opt/bin/install_slurm_plugins.sh && \
    rm /opt/bin/install_slurm_plugins.sh

# Update linker cache
RUN ldconfig

# Delete users & home because they will be linked from jail
RUN rm /etc/passwd* /etc/group* /etc/shadow* /etc/gshadow*
RUN rm -rf /home

# Expose the port used for accessing slurmctld
EXPOSE 6817

# Create dir and file for multilog hack
RUN mkdir -p /var/log/slurm/multilog && \
    touch /var/log/slurm/multilog/current && \
    ln -s /var/log/slurm/multilog/current /var/log/slurm/slurmctld.log

# Copy & run the entrypoint script
COPY controller/slurmctld_entrypoint.sh /opt/bin/slurm/
RUN chmod +x /opt/bin/slurm/slurmctld_entrypoint.sh
ENTRYPOINT ["/opt/bin/slurm/slurmctld_entrypoint.sh"]