#!/bin/sh

if [[ -z "$1" ]]
then
    echo "Usage: ./slave.sh log_dir"
    echo "Example: ./slave.sh ~/logs"
    exit
fi

. ./common.sh

SLAVE_RPC_PORT=40000

master=$(head -1 $NODES_FILE)
slaves=$(cat $NODES_FILE | tail -n +2 | head -$1)

for slave in $slaves
do
    taktuk -m $slave broadcast exec [ "cd $GENERATED_FILES_DIR; $BIN_DIR/slave $master $MASTER_RPC_PORT $slave $SLAVE_RPC_PORT $1" ] &
done
