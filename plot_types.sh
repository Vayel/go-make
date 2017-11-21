#!/bin/sh

if [[ -z "$1" ]]
then
    echo "Usage: ./plot_types.sh measures_file"
    exit
fi

python convert_to_gnuplot.py "$1"
./plot_measure_type.sh times.png times.txt "Temps d'exécution (ms)"
./plot_measure_type.sh speedups.png speedups.txt "Accélération"
./plot_measure_type.sh efficiencies.png efficiencies.txt "Efficacité"
