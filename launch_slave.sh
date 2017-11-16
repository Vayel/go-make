#!/bin/sh

##Prerequesit##
#./launch_master.sh

nodes=$(uniq $OAR_FILE_NODES)
master=$(head -1 $OAR_FILE_NODES)
nodes_without_master=$(uniq $OAR_FILE_NODES | tail -n +2)

for machine in $nodes_without_master
do
    taktuk -m $machine broadcast exec [ "nohup ./go-make-master/bin/slave $master 10000 $machine 40000 go-make-master/outputfiles/" ] &
done


