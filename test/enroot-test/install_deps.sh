#!/bin/bash

set -e

export DEBIAN_FRONTEND=noninteractive

apt update
apt install -y iputils-ping dnsutils telnet strace vim tree lsof git wget curl gnupg2 build-essential parallel jq squashfs-tools zstd
