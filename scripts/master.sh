#!/bin/sh

if [[ -z "$1" ]] || [[ -z "$2" ]] || [[ -z "$3" ]]
then
    echo "Usage:"
    echo -e "\t./master.sh makefile_path first_rule log_path"
    echo "Example:"
    echo -e "\t./master.sh ~/go-make/makefiles/0 all ~/logs/master.json"
    exit
fi

. ./common.sh

master=$(head -1 $NODES_FILE)
taktuk -m $master broadcast exec [ "cd $GENERATED_FILES_DIR; $BIN_DIR/master $1 $2 $MASTER_RPC_PORT $3" ] 
