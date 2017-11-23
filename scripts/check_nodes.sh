#!/bin/sh

. ./common.sh

nodes=$(cat $UNIQ_NODES_FILE)
> $NODES_FILE
for node in $nodes
do
    output=$(taktuk -o taktuk -o status -o connector -o error='"Error: $line\n"' -m $node broadcast exec [ "go help" ])
    if [[ $output == Error:* ]]
    then
        echo $output
    else
        echo $node >> $NODES_FILE
    fi
done

echo -e "\nAvailable nodes: $(cat $NODES_FILE | wc -l)"
cat $NODES_FILE
