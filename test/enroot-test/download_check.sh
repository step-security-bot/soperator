#!/bin/bash

set -e

curl -fSsL -O https://github.com/NVIDIA/enroot/releases/download/v3.5.0/enroot-check_3.5.0_$(uname -m).run
chmod +x enroot-check_*.run
