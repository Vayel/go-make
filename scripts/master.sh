#!/bin/sh

if [[ -z "$1" ]] || [[ -z "$2" ]] || [[ -z "$3" ]] || [[ -z "$4" ]]
then
    echo "Usage:"
    echo -e "\t./master.sh makefile_path first_rule log_path node_path"
    echo "Example:"
    echo -e "\t./master.sh ~/go-make/makefiles/0 all ~/logs/master.json ~/logs/nodes/master.txt"
    exit
fi

. ./common.sh

master=$(head -1 $NODES_FILE)
echo $master > $4
taktuk -m $master broadcast exec [ "cd $GENERATED_FILES_DIR; $BIN_DIR/master $1 $2 $MASTER_RPC_PORT $3" ] 
