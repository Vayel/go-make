#!/bin/sh

# TODO: remove non-working nodes

if [[ -z "$OAR_FILE_NODES" ]]
then
    echo "Cannot find OAR_FILE_NODES variable"
    exit 1
fi

BIN_DIR=~/
NODES_FILE=~/nodes.txt
uniq $OAR_FILE_NODES > $NODES_FILE
MASTER_RPC_PORT=10000
