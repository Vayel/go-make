#!/bin/sh

if [[ -z "$1" ]] || [[ -z "$2" ]]
then
    echo "Usage:"
    echo -e "\t./plot.sh measures output_dir"
    exit
fi

python3 convert_to_gnuplot.py "$1" "$2"
./plot_one.sh times.png times.txt "Time (ms)" "$2"
./plot_one.sh speedups.png speedups.txt "Speedup" "$2"
./plot_one.sh efficiencies.png efficiencies.txt "Efficiency" "$2"
