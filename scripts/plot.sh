#!/bin/sh

if [[ -z "$1" ]]
then
    echo "Usage: ./plot.sh measures_file"
    exit
fi

python3 convert_to_gnuplot.py "$1"
./plot_one.sh times.png times.txt "Time (ms)"
./plot_one.sh speedups.png speedups.txt "Speedup"
./plot_one.sh efficiencies.png efficiencies.txt "Efficiency"
