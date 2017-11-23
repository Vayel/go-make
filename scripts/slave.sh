#!/bin/sh

if [[ -z "$1" ]] || [[ -z "$2" ]]
then
    echo "Usage:"
    echo -e "\t./slave.sh log_dir nodes_path"
    echo "Example:"
    echo -e "\t./slave.sh ~/logs ~/logs/nodes/slaves.txt"
    exit
fi

. ./common.sh

SLAVE_RPC_PORT=40000

master=$(head -1 $NODES_FILE)
slaves=$(cat $NODES_FILE | tail -n +2 | head -$1)
echo $slaves > $2

for slave in $slaves
do
    taktuk -m $slave broadcast exec [ "cd $GENERATED_FILES_DIR; $BIN_DIR/slave $master $MASTER_RPC_PORT $slave $SLAVE_RPC_PORT $1" ] &
done
