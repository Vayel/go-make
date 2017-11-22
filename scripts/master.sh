#!/bin/sh

#It takes as first argument the makefile it is by default 10
#Exemple: ./launch_master 10

##Prerequisit##
#./nodes_configuration.sh
#UNIQ_FILE_NODES exported by nodes_configuration.sh

if [[ -z "$1" ]] || [[ -z "$2" ]] || [[ -z "$3" ]]
then
    echo "Usage: ./master.sh makefile_path first_rule log_path"
    echo "Example: ./master.sh ~/go-make/makefiles/0 all ~/logs/master.json"
    exit
fi

. ./common.sh

RPC_PORT=10000
OUTPUT_DIR=/tmp

master=$(head -1 $NODES)
taktuk -m $master broadcast exec [ "~/master $1 $2 $RPC_PORT $OUTPUT_DIR $3" ] 
