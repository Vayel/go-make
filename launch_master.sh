#!/bin/sh

#It takes as first argument the makefile it is by default 10
#Exemple: ./launch_master 10

##Prerequisit##
#./nodes_configuration.sh

if [ -z "$1" ]
then
    makefile="10"
else
    makefile=$1
fi

taktuk -m "$(head -1 $OAR_FILE_NODES)" broadcast exec [ "nohup ./go-make-master/bin/master go-make-master/makefiles/"$makefile" all 10000 go-make-master/outputfiles/" ] &
