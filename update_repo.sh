export UNIQ_FILE_NODES=~/grid5000_nodes.txt

uniq $OAR_FILE_NODES > $UNIQ_FILE_NODES

taktuk -s -o connector -o status -o output='"$host: $line\n"' -f $UNIQ_FILE_NODES broadcast exec [ "cd /tmp/go-make; git pull" ]
