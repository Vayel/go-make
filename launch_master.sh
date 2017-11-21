#!/bin/sh

#It takes as first argument the makefile it is by default 10
#Exemple: ./launch_master 10

##Prerequisit##
#./nodes_configuration.sh
#UNIQ_FILE_NODES exported by nodes_configuration.sh

if [[ -z "$1" ]]
then
    echo "Usage: ./launch_master.sh makefile_path"
    echo "Ex: ./launch_master.sh /tmp/go-make/makefiles/xxx"
    exit
fi

masterport=10000
output=/tmp/go-make/outputfiles/
logdir=~/go-make/logs/
target=all

export UNIQ_FILE_NODES=~/grid5000_nodes.txt
master=$(head -1 $UNIQ_FILE_NODES)
taktuk -m $master broadcast exec [ "/tmp/go-make/bin/master $1 $target $masterport $output $logdir" ] 
sleep 3 # Wait for slaves to terminate
