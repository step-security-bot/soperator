unset NVIDIA_VISIBLE_DEVICES

++++++++++ remove "@" from the beggining of --ldconfig in 98-* hook

We also need "mount --rbind /mnt/jail /mnt/jail" AT THE VERY BEGINNING OF complement_jail.fs

# After that, enroot import & enroot create work

# BUT:
# $ enroot start --root --rw ubuntu sh -c 'apt update && apt install -y wget build-essential'
# enroot-nsenter: failed to create user namespace: Operation not permitted

# REASON:
# EPERM (since Linux 3.9)
# CLONE_NEWUSER was specified in flags and the caller is in
# a chroot environment (i.e., the caller's root directory
# does not match the root directory of the mount namespace
# in which it resides).

# unshare --mount /bin/bash
# chroot /mnt/jail /bin/bash
# mount --make-rslave / # don't propagate mount changes back to the host


# unshare --mount /bin/bash
# mount --rbind /mnt/jail /mnt/jail
# mkdir -p /mnt/jail/mnt/host
# pivot_root /mnt/jail /mnt/jail/mnt/host
# cd /
# (?) umount /mnt/host
