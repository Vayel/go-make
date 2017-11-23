#!/bin/sh

# ssh <login>@access.grid5000.fr
# ssh <site>
# Get nodes: https://www.grid5000.fr/mediawiki/index.php/Getting_Started
# oarsub -I -l nodes=TODO,walltime=1 -t deploy

. ./common.sh
uniq $OAR_FILE_NODES > $UNIQ_NODES_FILE
cat $UNIQ_NODES_FILE > $NODES_FILE

# TODO: deploy a custom image?

kadeploy3 -f $NODES_FILE -e jessie-x64-nfs -k
taktuk -l root -s -o connector -o status -o output='"$host: $line\n"' -f $NODES_FILE broadcast exec [ "apt-get update" ]
taktuk -l root -s -o connector -o status -o output='"$host: $line\n"' -f $NODES_FILE broadcast exec [ "apt-get install -y golang-go" ]
