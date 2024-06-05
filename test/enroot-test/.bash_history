    1  ll
    2  cd enroot-test/
    3  ll
    4  ./fix_perms.sh 
    5  ll
    6  ./install_deps.sh
    7  ll
    8  ./install_toolkit.sh 
    9  ./install_enroot.sh 
   10  ll
   11  source nvidia_env.sh 
   12  ./download_check.sh 
   13  ll
   14  ./enroot-check_3.5.0_x86_64.run --verify
   15  ./enroot-check_3.5.0_x86_64.run --keep
   16  ./enroot-check_3.5.0_x86_64.run
   17  mount -o remount,rw /proc
   18  ./enroot-check_3.5.0_x86_64.run
   19  ls
   20  cd enroot-check_3.5.0_x86_64
   21  ls
   22  cd .enroot/
   23  ll
   24  cd etc
   25  ll
   26  cd system/
   27  ll
   28  cd mounts.d/
   29  ll
   30  vim 10-system.fstab 
   31  cd ../../../..
   32  cd ..
   33  ./enroot-check_3.5.0_x86_64.run 
   34  ll
   35  cat enroot-check_3.5.0_x86_64/.enroot/etc/system/mounts.d/10-system.fstab 
   36  vim /usr/local/etc/enroot/mounts.d/10-system.fstab 
   37  ./enroot-check_3.5.0_x86_64
   38  ./enroot-check_3.5.0_x86_64.run
   39  cat enroot-check_3.5.0_x86_64/.enroot/etc/system/mounts.d/10-system.fstab 
   40  vim /usr/local/etc/enroot/mounts.d/10-system.fstab 
   41  rm -rf enroot-check_3.5.0_x86_64
   42  ./enroot-check_3.5.0_x86_64.run 
   43  ./enroot-check_3.5.0_x86_64.run --keep
   44  ls
   45  cat enroot-check_3.5.0_x86_64/.enroot/etc/system/mounts.d/10-system.fstab 
   46  rm -rf enroot-check_3.5.0_x86_64
   47  vim /usr/local/etc/enroot/mounts.d/10-system.fstab 
   48  ./enroot-check_3.5.0_x86_64.run 
   49  ll /tmp
   50  ll
   51  cd /etc
   52  ll
   53  cd /
   54  cd /usr/local/
   55  ls
   56  cd etc
   57  ll
   58  cd enroot/
   59  ll
   60  vim enroot.conf
   61  ll enroot.conf.d/
   62  cat environ.d/
   63  cat environ.d
   64  printenv | grep NV
   65  export NVIDIA_DRIVER_CAPABILITIES=""
   66  printenv | grep NV
   67  unset NVIDIA_DRIVER_CAPABILITIES
   68  ll
   69  cd /enroot-test/
   70  ll
   71  cat nvidia_env.sh 
   72  export NVIDIA_VISIBLE_DEVICES=all
   73  printenv
   74  nvidia-smi
   75  echo $NVIDIA_VISIBLE_DEVICES
   76  unset NVIDIA_VISIBLE_DEVICES
   77  ll
   78  nvidia-smi
   79  ll
   80  printenv | grep NVI
   81  NVIDIAll
   82  ll
   83  nvidia-smi --list-gpus
   84  export NVIDIA_VISIBLE_DEVICES=GPU-47ce1a79-c8cc-a439-2346-1ccf3f2dfb03
   85  echo $NVIDIA_VISIBLE_DEVICES 
   86  ll
   87  ./enroot-check_3.5.0_x86_64.run 
   88  unset NVIDIA_VISIBLE_DEVICES
   89  printenv | grep NV
   90  cat /usr/local/etc/enroot/mounts.d/10-system.fstab
   91  vim /usr/local/etc/enroot/mounts.d/10-system.fstab 
   92  ./enroot-check_3.5.0_x86_64.run 
   93  printenv | grep NV
   94  ll
   95  ./enroot-check_3.5.0_x86_64.run 
   96  history > .bash_history
