#!/bin/sh

#It takes as first argument the makefile it is by default 10
#Exemple: ./launch_master 10

##Prerequisit##
#./nodes_configuration.sh
#UNIQ_FILE_NODES exported by nodes_configuration.sh

slaveport = 40000
output = /tmp/go-make/outputfiles/
logdir = ~/go-make/logs/

export UNIQ_FILE_NODES=~/grid5000_nodes.txt
master=$(head -1 $UNIQ_FILE_NODES)
taktuk -m $master broadcast exec [ "/tmp/go-make/bin/master $1 $slaveport $output $logfile" ] 
#taktuk -m $master broadcast exec [ "make --directory /tmp/go-make/ run_master $1" ] 
#taktuk -m $master broadcast exec [ "make --directory /tmp/go-make/ run_master makefiles/4" ] 
