#!/bin/bash

set -e

pushd /usr/src
    echo "Clone git repository with enroot sources"
    git clone https://github.com/NVIDIA/enroot.git
    cd enroot
    git checkout v3.5.0

    echo "Build enroot"
    make -j$(nproc)

    echo "Install enroot"
    make install

    echo "Allow unprivileged users to import images"
    make setcap
popd
