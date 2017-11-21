#!/bin/sh

##Prerequisit##
#./nodes_configuration.sh
#UNIQ_FILE_NODES exported by nodes_configuration.sh

if [[ -z "$1" ]]
then
    echo "Usage: ./launch_sequential.sh makefile_path"
    echo "Ex: ./launch_master.sh /tmp/go-make/makefiles/xxx"
    exit
fi

export UNIQ_FILE_NODES=~/grid5000_nodes.txt
node=$(head -1 $UNIQ_FILE_NODES)
# TODO: change Makefile
taktuk -m $node broadcast exec [ "/tmp/go-make/bin/sequential $1 all ~/go-make/logs/time_seq.json" ] 
