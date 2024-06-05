#!/bin/bash

set -e

enroot import 'docker://nvcr.io#nvidia/cuda:11.4.0-devel-ubuntu20.04'
enroot create --name cuda nvidia+cuda+11.4.0-devel-ubuntu20.04.sqsh

enroot start --root --rw cuda sh -c 'apt update && apt install -y wget build-essential'
enroot start --root --rw cuda sh -c 'wget https://github.com/NVIDIA/cuda-samples/archive/refs/tags/v11.4.tar.gz'
enroot start --root --rw cuda sh -c 'tar -xzvf v11.4.tar.gz'
enroot start --root --rw cuda sh -c 'cd cuda-samples-11.4 && make -j'
enroot start --root --rw cuda sh -c 'cd cuda-samples-11.4/Samples/deviceQuery && make run'
enroot start --root --rw cuda sh -c 'cd cuda-samples-11.4/Samples/bandwidthTest && make run'
#enroot start --rw cuda sh -c 'cd /usr/local/cuda/samples/5_Simulations/nbody && make -j'
