#!/bin/sh

. ./common.sh

cd ..
make master
make slave
make sequential

mv bin/master bin/slave bin/sequential $BIN_DIR
