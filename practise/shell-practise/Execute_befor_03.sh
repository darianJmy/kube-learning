#!/usr/bin/env sh
#----------------------------------------------------
# (C) COPYRIGHT Beijing NewData Tech .LTD 2021
# All Rights Reserverd
#----------------------------------------------------

syntax() {
    exit 22
}

Users=root,oracle,grid

f_NoSecret() {
    array_user=($(echo ${Users} | sed "s/,/ /g"))
    for ((i = 0; i < ${#array_user[@]}; i++)); do
        for ((j = 1; j < 3; j++)); do
            if [[ ${array_user[i]} == "root" ]]; then
                cat $HOME/auth.keys | awk "NR==$j {print}" >> $HOME/.ssh/authorized_keys
                cat $HOME/auth_node.keys  | awk "NR==$j {print}" >> $HOME/.ssh/authorized_keys
            else
                cat $HOME/auth.keys | awk "NR==$j {print}" >> /home/${array_user[i]}/.ssh/authorized_keys  
                cat $HOME/auth_node.keys | awk "NR==$j {print}" >> /home/${array_user[i]}/.ssh/authorized_keys  
            fi
        done
    done
}
f_NoSecret