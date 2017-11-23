#!/bin/sh

if [[ -z "$1" ]] || [[ -z "$2" ]] || [[ -z "$3" ]]
then
    echo "Usage: ./sequential.sh makefile_path first_rule log_path"
    echo "Example: ./sequential.sh ~/go-make/makefiles/0 all ~/logs/seq.json"
    exit
fi

. ./common.sh

node=$(head -1 $NODES_FILE)
taktuk -m $node broadcast exec [ "cd $GENERATED_FILES_DIR; $BIN_DIR/sequential $1 $2 $3" ] 
