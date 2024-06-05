#!/bin/bash

chown 0:0 10-system.fstab
cp ./10-system.fstab /usr/local/etc/enroot/mounts.d/
