#!/bin/sh

##Prerequesit##
#./launch_master.sh

##Prerequisit##
#./nodes_configuration.sh
#UNIQ_FILE_NODES exported by nodes_configuration.sh

master=$(head -1 $UNIQ_FILE_NODES)
nodes_without_master=$(cat $UNIQ_FILE_NODES | tail -n +2)
machine=genepi-22.grenoble.grid5000.fr
for machine in $nodes_without_master
do
    taktuk -m $machine broadcast exec [ "nohup /tmp/go-make/bin/slave $master 10000 $machine 40000 /tmp/go-make/outputfiles/" ] &
	taktuk -m $machine broadcast exec [ "/tmp/go-make/bin/slave $master 10000 $machine 40000 outputfiles/" ]
done


