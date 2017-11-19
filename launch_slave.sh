#!/bin/sh

##Prerequesit##
#./launch_master.sh

##Prerequisit##
#./nodes_configuration.sh
#UNIQ_FILE_NODES exported by nodes_configuration.sh

export UNIQ_FILE_NODES=/home/lcarre/grid5000_nodes.txt
master=$(head -1 $UNIQ_FILE_NODES)
nodes_without_master=$(cat $UNIQ_FILE_NODES | tail -n +2)
#machine=genepi-14.grenoble.grid5000.fr
for machine in $nodes_without_master
do
    taktuk -m $machine broadcast exec [ "/tmp/go-make/bin/slave $master 10000 $machine 40000 /tmp/go-make/outputfiles/" ]
	#taktuk -m $machine broadcast exec [ "/tmp/go-make/bin/slave $master 10000 $machine 40000 /tmp/go-make/outputfiles/" ]
done


