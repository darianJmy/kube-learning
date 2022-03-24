#!/usr/bin/env sh
#----------------------------------------------------
# (C) COPYRIGHT Beijing NewData Tech .LTD 2021
# All Rights Reserverd
#----------------------------------------------------

syntax() {
    echo "bash $0 -PUB_IP1 10.10.33.36 -PUB_IP2 10.10.33.37 -PUB_HOSTNAME1 oracle-node01 -PUB_HOSTNAME2 oracle-node02"

    echo "Parameter description: "
    echo "[ PUB_IP1 ]   : Public IP of node 1"
    echo "[ PUB_IP2 ]   : Public IP of node 1"
    echo "[ PUB_HOSTNAME1]   : Public HOSTNAME of node 1 "
    echo "[ PUB_HOSTNAME2]   : Public HOSTNAME of node 2 "
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
PRI_HOSTNAME1=$PUB_HOSTNAME1"priv"
PRI_HOSTNAME2=$PUB_HOSTNAME2"priv"
Users=root,oracle,grid


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
if [ -z "${PUB_IP1}" -o -z "${PUB_IP2}" -o -z "${PUB_HOSTNAME1}" -o -z "${PUB_HOSTNAME2}" ]; then
    f_PrintLog "ERROR" "One or more parameters are not entered: PUB_IP1, PUB_IP2, PUB_HOSTNAME1, PUB_HOSTNAME2"
fi

chmod 755 ${WorkDir}/${ShellName}
if [ "$(id -u)" != '0' ]; then
    f_PrintLog "INFO" "Switch to root and re-execute the script."
    # exec sudo gosu root "${WorkDir}/${ShellName}" "$@"
    exec sudo "${WorkDir}/${ShellName}" "$@"
fi

# 清除所有sshkey
f_CleanSshKey() {
    array_user=($(echo ${Users} | sed "s/,/ /g"))
    for ((i = 0; i < ${#array_user[@]}; i++)); do
        if [[ ${array_user[i]} != "root" ]]; then
            rm -rf /home/${array_user[i]}/.ssh/*
        fi
    done
}
f_CleanSshKey

# 为交互式生成密钥文件
f_SSHkeygen() {
    expect << EOF
    spawn su oracle -c "ssh-keygen -f /home/oracle/.ssh/id_rsa -N '' -t rsa"
    expect "(y/n)?" {send "n\r"}
EOF

    expect << EOF
    spawn su grid -c "ssh-keygen -f /home/grid/.ssh/id_rsa -N '' -t rsa"
    expect "(y/n)?" {send "n\r"}
EOF

    expect << EOF
    spawn su root -c "ssh-keygen -f /root/.ssh/id_rsa -N '' -t rsa"
    expect "(y/n)?" {send "n\r"}
EOF
}
f_SSHkeygen

# 公钥
f_SSHkeyscan() {
    array_user=($(echo ${Users} | sed "s/,/ /g"))
    for ((i = 0; i < ${#array_user[@]}; i++)); do
        if [[ ${array_user[i]} == "root" ]]; then
            ssh-keyscan $PUB_HOSTNAME1 >> $HOME/.ssh/known_hosts
            ssh-keyscan $PUB_HOSTNAME2 >> $HOME/.ssh/known_hosts
            ssh-keyscan $PRI_HOSTNAME1 >> $HOME/.ssh/known_hosts
            ssh-keyscan $PRI_HOSTNAME2 >> $HOME/.ssh/known_hosts
        else
            su ${array_user[i]} -c  "ssh-keyscan $PUB_HOSTNAME1 >> /home/${array_user[i]}/.ssh/known_hosts"
            su ${array_user[i]} -c  "ssh-keyscan $PUB_HOSTNAME2 >> /home/${array_user[i]}/.ssh/known_hosts"
            su ${array_user[i]} -c  "ssh-keyscan $PRI_HOSTNAME1 >> /home/${array_user[i]}/.ssh/known_hosts"
            su ${array_user[i]} -c  "ssh-keyscan $PRI_HOSTNAME2 >> /home/${array_user[i]}/.ssh/known_hosts"
        fi
    done
}
f_SSHkeyscan

# 免密
f_SSHcopyid() {
    rm -rf $HOME/auth.keys
    array_user=($(echo ${Users} | sed "s/,/ /g"))
    for ((i = 0; i < ${#array_user[@]}; i++)); do
        if [[ ${array_user[i]} == "root" ]]; then
            cat $HOME/.ssh/id_rsa.pub > /root/auth.keys
        else
            cat /home/${array_user[i]}/.ssh/id_rsa.pub >> /root/auth.keys
        fi
    done

    ip a | grep $PUB_IP1
    if [ "$?" -eq 0 ]; then
        
        scp $HOME/auth.keys root@$PUB_IP2:/root/auth_node.keys
    else
        scp $HOME/auth.keys root@$PUB_IP1:/root/auth_node.keys
    fi
}
f_SSHcopyid  

