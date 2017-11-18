#!/bin/sh

python convert_to_gnuplot.py "$1"
./plot_measure_type.sh times.png times.txt "Temps d'exécution"
./plot_measure_type.sh speedups.png speedups.txt "Accélération"
./plot_measure_type.sh efficiencies.png efficiencies.txt "Efficacité"
