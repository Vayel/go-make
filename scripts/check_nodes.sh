#!/bin/sh

. ./common.sh

nodes=$(cat $NODES_FILE)
> $NODES_FILE
for node in $nodes
do
    ret=$(taktuk -o status -m $node broadcast exec [ "ls" ])
    if [[ -z $ret ]]
    then
        echo $ret >> $NODES_FILE
    fi
done

echo "Available nodes:"
cat $NODES_FILE | wc -l
cat $NODES_FILE
