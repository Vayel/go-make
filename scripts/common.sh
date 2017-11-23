#!/bin/sh

# TODO: remove non-working nodes

if [[ -z "$OAR_FILE_NODES" ]]
then
    echo "Cannot find OAR_FILE_NODES variable"
    exit 1
fi

BIN_DIR=~/
PROJECT_DIR=~/go-make
NODES_FILE=~/nodes.txt
MASTER_RPC_PORT=10000
