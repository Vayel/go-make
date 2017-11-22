#!/bin/sh

# TODO: remove non-working nodes

export NODES=`uniq $OAR_FILE_NODES`
export MASTER_RPC_PORT=10000
