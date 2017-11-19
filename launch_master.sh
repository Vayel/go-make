#!/bin/sh

#It takes as first argument the makefile it is by default 10
#Exemple: ./launch_master 10

##Prerequisit##
#./nodes_configuration.sh
#UNIQ_FILE_NODES exported by nodes_configuration.sh

master=$(head -1 $UNIQ_FILE_NODES)
taktuk -m $master broadcast exec [ "nohup make --directory /tmp/go-make/ run_master $1" ] 
taktuk -m $master broadcast exec [ "nohup make --directory /tmp/go-make/ run_master makefiles/4 &" ] 
