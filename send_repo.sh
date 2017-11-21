#!/bin/sh

taktuk -l root -s -o connector -o status -o output='"$host: $line\n"' -f $UNIQ_FILE_NODES broadcast put { ~/go-make } { /tmp }
