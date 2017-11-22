#!/bin/sh

if [[ -z "$1" ]] || [[ -z "$2" ]] || [[ -z "$3" ]]
then
    echo "Usage: ./master.sh makefile_path first_rule log_path"
    echo "Example: ./master.sh ~/go-make/makefiles/0 all ~/logs/master.json"
    exit
fi

. ./common.sh

OUTPUT_DIR=/tmp

master=$(head -1 $NODES)
taktuk -m $master broadcast exec [ "~/master $1 $2 $MASTER_RPC_PORT $OUTPUT_DIR $3" ] 
