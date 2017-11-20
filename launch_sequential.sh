#!/bin/sh

##Prerequisit##
#./nodes_configuration.sh
#UNIQ_FILE_NODES exported by nodes_configuration.sh

export UNIQ_FILE_NODES=~/grid5000_nodes.txt
node=$(head -1 $UNIQ_FILE_NODES)
# TODO: change Makefile
taktuk -m $node broadcast exec [ "/tmp/go-make/bin/sequential /tmp/go-make/makefiles/examples/blender_2.59/Makefile /tmp/go-make/outputfiles/out.avi ~/go-make/logs/time_seq.json" ] 
