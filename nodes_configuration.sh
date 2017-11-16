#!/bin/sh

###Prerequesit###

#start by connecting to grid5000
#ssh <login>@access.grid5000.fr
#ssh <site>
#reserve nodes as in  https://www.grid5000.fr/mediawiki/index.php/Getting_Started
#exemple: oarsub -I -l nodes=4,walltime=0:45 -t deploy

#launch the script

kadeploy3 -f $OAR_FILE_NODES -e wheezy-x64-nfs -k

uniq $OAR_FILE_NODES | taktuk -l root -s -o connector -o status -o output='"$host: $line\n"' -f - broadcast exec [ "apt-get update" ]

uniq $OAR_FILE_NODES | taktuk -l root -s -o connector -o status -o output='"$host: $line\n"' -f - broadcast exec [ "apt-get install golang-go unzip git -y" ]

uniq $OAR_FILE_NODES | taktuk -s -o connector -o status -o output='"$host: $line\n"' -f - broadcast exec [ "wget https://github.com/Vayel/go-make/archive/master.zip" ]

uniq $OAR_FILE_NODES | taktuk -s -o connector -o status -o output='"$host: $line\n"' -f - broadcast exec [ "unzip master.zip" ]

uniq $OAR_FILE_NODES | taktuk -s -o connector -o status -o output='"$host: $line\n"' -f - broadcast exec [ "cd go-make-master && make" ]
