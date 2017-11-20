#!/bin/sh

##Prerequesit##
#./launch_master.sh n_slaves

##Prerequisit##
#./nodes_configuration.sh
#UNIQ_FILE_NODES exported by nodes_configuration.sh

export UNIQ_FILE_NODES=~/grid5000_nodes.txt
master=$(head -1 $UNIQ_FILE_NODES)
nodes_without_master=$(cat $UNIQ_FILE_NODES | tail -n +2 | head -$1)
for machine in $nodes_without_master
do
    taktuk -m $machine broadcast exec [ "/tmp/go-make/bin/slave $master 10000 $machine 40000 /tmp/go-make/outputfiles/" ]
done


