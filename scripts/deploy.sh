#!/bin/sh

# ssh <login>@access.grid5000.fr
# ssh <site>
# Get nodes: https://www.grid5000.fr/mediawiki/index.php/Getting_Started
# oarsub -I -l nodes=20,walltime=1 -t deploy

export UNIQ_FILE_NODES=~/nodes.txt
uniq $OAR_FILE_NODES > $UNIQ_FILE_NODES

kadeploy3 -f $UNIQ_FILE_NODES -e jessie-x64-nfs -k
taktuk -l root -s -o connector -o status -o output='"$host: $line\n"' -f $UNIQ_FILE_NODES broadcast exec [ "apt-get update" ]
taktuk -l root -s -o connector -o status -o output='"$host: $line\n"' -f $UNIQ_FILE_NODES broadcast exec [ "apt-get install -y golang-go" ]
