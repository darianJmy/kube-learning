#!/usr/bin/env sh
#----------------------------------------------------
# (C) COPYRIGHT Beijing NewData Tech .LTD 2021
# All Rights Reserverd
#----------------------------------------------------

syntax() {
    echo "bash $0 -PUB_IP1 10.10.33.36 -PUB_IP2 10.10.33.37 -PUB_HOSTNAME1 oracle-node01 -PUB_HOSTNAME2 oracle-node02 -disks /dev/sda1,/dev/sdb1,/dev/sdc1,/dev/sdd1,/dev/sde1 -installdir /u01"

    echo "Parameter description: "
    echo "[ PUB_IP1 ]   : Public IP of node 1"
    echo "[ PUB_IP2 ]   : Public IP of node 1"
    echo "[ PUB_HOSTNAME1]   : Public HOSTNAME of node 1 "
    echo "[ PUB_HOSTNAME2]   : Public HOSTNAME of node 2 "
    echo "[ disks ]   : Disk for storing RAC , multiple paths can be entered and separated by \",\""
    echo "[ INSTALLDIR ]   : install location for oracle"
    exit 22
}

# 博云初始化输入的参数
PUB_IP1=${{node1_pubIp}}
PUB_IP2=${{node2_pubIp}}
PUB_HOSTNAME1=${{node1_pubHostName}}
PUB_HOSTNAME2=${{node2_pubHostName}}

###
# 设置参数默认值
#PUB_IP1=10.10.33.36
#PUB_IP2=10.10.33.37
#PUB_HOSTNAME1=oracle-node01
#PUB_HOSTNAME2=oracle-node02#
###

DISKS=/dev/sda1,/dev/sdb1,/dev/sdc1,/dev/sdd1,/dev/sde1
INSTALLDIR=/u01
LogLevel=3

# 初始变量
export LC_ALL=C
export LANG=en_US
export LDR_CNTRL=MAXDATA=0x80000000
export TZ="BEIST-8"
ShellName="$(echo $0 | awk -F / '{print $NF}')"
WorkDir="$(echo $0 | sed s/${ShellName}//g)"
[ -z "${WorkDir}" ] && WorkDir=${PWD}
cd ${WorkDir}
WorkDir=${PWD}
LogDir="${WorkDir}/logs"
mkdir -p ${LogDir}
chmod ugo+rwx ${LogDir} 2>/dev/null
LogFile=${LogDir}/${ShellName}.log
typeset uname_a=$(uname -a)
typeset Platform=$(echo ${uname_a%% *} | tr a-z A-Z)
if [ "${Platform}" = "LINUX" ]; then
    export PATH="$PATH:/usr/local/sbin:/sbin:/bin:/usr/sbin:/usr/bin:/root/bin"
    typeset SysID=$(uname -m)
elif [ "${Platform}" = "AIX" ]; then
    export PATH="$PATH:/usr/bin:/etc:/usr/sbin:/sbin"
    typeset SysID=$(uname -u | awk -F, '{print substr($2,3)}')
fi
typeset UserName=$(whoami)
typeset HostName=$(hostname)
ShellOption="$@"
PShellName=$(echo "${ShellName}" | sed 's/^[0-9]*_//')

# 判断参数不能为空
if [ -z "${PUB_IP1}" -o -z "${PUB_IP2}" -o -z "${PUB_HOSTNAME1}" -o -z "${PUB_HOSTNAME2}" -o -z "${DISKS}" ]; then
    f_PrintLog "ERROR" "One or more parameters are not entered: PUB_IP1, PUB_IP2, PRI_IP1, PRI_IP2, VIP1, VIP2, SCANIP, DISKS"
fi

chmod 755 ${WorkDir}/${ShellName}
if [ "$(id -u)" != '0' ]; then
    f_PrintLog "INFO" "Switch to root and re-execute the script."
    # exec sudo gosu root "${WorkDir}/${ShellName}" "$@"
    exec sudo "${WorkDir}/${ShellName}" "$@"
fi

# 设置主机名
f_SetHostName() {
    ip a | grep $PUB_IP1 > /dev/null
    if [ "$?" -eq 0 ]; then
        hostnamectl set-hostname $PUB_HOSTNAME1
    else
        ip a | grep $PUB_IP2 > /dev/null
        if [ "$?" -eq 0 ]; then
            hostnamectl set-hostname $PUB_HOSTNAME2
        fi    
    fi
}
f_SetHostName

# 检查selinux状态
f_Selinuxchk() {
    cat /etc/selinux/config | grep ^SELINUX=enforcing > /dev/null
    if [ "$?" -eq 0 ]; then
        sed -i 's/^SELINUX=enforcing/SELINUX=disabled/g' /etc/selinux/config    
    fi

    getenforce  | grep Enforcing > /dev/null
    if [ "$?" -eq 0 ]; then
        setenforce 0
    fi    
}
f_Selinuxchk

# 格式化磁盘
f_DiskFormat() {
    ip a | grep $PUB_IP1 > /dev/null
    if [ "$?" -eq 0 ]; then
        array_disks=($(echo ${DISKS} | sed "s/,/ /g"))
        for ((i = 0; i < ${#array_disks[@]}; i++)); do
            if [[ "/dev/sda1" != ${array_disks[i]} ]]; then
                dd if=/dev/zero of=${array_disks[i]} bs=32768 count=32768 conv=notrunc oflag=direct
            fi
        done
    fi    
}
f_DiskFormat

# yum 安装
f_YumInstall() {
    yum install -y expect
}
f_YumInstall