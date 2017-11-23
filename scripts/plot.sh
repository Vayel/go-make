#!/bin/sh

if [[ -z "$1" ]] || [[ -z "$2" ]]
then
    echo "Usage:"
    echo -e "\t./plot.sh measures output_dir"
    exit
fi

python3 convert_to_gnuplot.py "$1" "$2"
./plot_one.sh "$2/times.png" "$2/times.txt" "Time (ms)"
./plot_one.sh "$2/speedups.png" "$2/speedups.txt" "Speedup"
./plot_one.sh "$2/efficiencies.png" "$2/efficiencies.txt" "Efficiency"
