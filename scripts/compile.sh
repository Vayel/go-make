#!/bin/sh

. ./common.sh

node=$(head -1 $NODES_FILE)

for prog in "master" "slave" "sequential"
do
taktuk -m $node broadcast exec [ "cd $PROJECT_DIR; make $prog; mv bin/$prog $BIN_DIR" ] 
done
