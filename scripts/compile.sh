#!/bin/sh

. ./common.sh

node=$(head -1 $NODES_FILE)
taktuk -l root -m $node broadcast exec [ "apt-get install make" ] 

for prog in "master" "slave" "sequential"
do
    taktuk -m $node broadcast exec [ "cd $PROJECT_DIR; make $prog; mv bin/$prog $BIN_DIR" ] 
done
