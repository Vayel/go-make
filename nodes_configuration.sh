#!/bin/sh

###Prerequesit###

#start by connecting to grid5000
#ssh <login>@access.grid5000.fr
#ssh <site>
#reserve nodes as in  https://www.grid5000.fr/mediawiki/index.php/Getting_Started
#exemple: oarsub -I -l nodes=4,walltime=0:15 -t deploy

#launch the script


UNIQ_FILE_NODES=$(uniq $OAR_FILE_NODES)

kadeploy3 -f $UNIQ_FILE_NODES -e wheezy-x64-nfs -k

taktuk -l root -s -o connector -o status -o output='"$host: $line\n"' -f $UNIQ_FILE_NODES broadcast exec [ "apt-get update" ]

taktuk -l root -s -o connector -o status -o output='"$host: $line\n"' -f $UNIQ_FILE_NODES broadcast exec [ "apt-get install golang-go unzip git -y" ]

taktuk -s -o connector -o status -o output='"$host: $line\n"' -f $UNIQ_FILE_NODES broadcast exec [ "wget https://github.com/Vayel/go-make/archive/master.zip" ]

taktuk -s -o connector -o status -o output='"$host: $line\n"' -f $UNIQ_FILE_NODES broadcast exec [ "unzip master.zip" ]

taktuk -s -o connector -o status -o output='"$host: $line\n"' -f $UNIQ_FILE_NODES broadcast exec [ "cd go-make-master && make" ]

#head -1 $UNIQ_FILE_NODES | taktuk -m - broadcast exec [ nohup ./go-make-master/bin/master makefiles/10 all 10000 outputfiles/ ]

#taktuk -m "$(head -1 $UNIQ_FILE_NODES)" broadcast exec [ "nohup ./go-make-master/bin/master go-make-master/makefiles/10 all 10000 go-make-master/outputfiles/" ] &

#master=$(head -1 $UNIQ_FILE_NODES)
#nodes_without_master=$(UNIQ_FILE_NODES | tail -n +2)
#for machine in $nodes_without_master
#do
    #taktuk -m $machine broadcast exec [ "nohup ./go-make-master/bin/slave $master 10000 $machine 40000 go-make-master/outputfiles/" ] &
#done

#To see the machines: cat $OAR_NODES_FILES
#To get the IP addresses: cat .ssh/known_hosts | grep <machine_name>
#To connect to a machine: ssh <machine_name>
#Run the program: ./bin/master <path-to-makefile> <rule-to-execute> <rpc-port> <result-dir>
#                 ./bin/slave  slave master-rpc-addr master-rpc-port slave-rpc-addr slave-rpc-port dependency-dir
