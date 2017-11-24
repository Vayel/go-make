#!/bin/sh

# Ex: ./plot_measure_type.sh output.png data.txt title y_label
# data.txt is the form of:
#
# 0 y01 y02 ...
# 1 y11 y12 ...
# 2 y21 y22 ...
# 3 y31 y32 ...
# ...
# yij is the measured value for i slave(s) at the jth execution
# y0j is the measured value for the jth sequential execution
# The last column is the average of all executions
# Each column is separated by exactly one space

N_COLS=`head -n1 "$2" | grep -o " " | wc -l`
N_COLS=$((N_COLS+1))
N_LINES=`wc -l < "$2"`
X_PADDING=0.5
MAX_SLAVES=`tail -n 1 "$2" | cut -d ' ' -f1`
plot_args="'$2' u 1:2 notitle"

# -1 because the last column is the average and is plotted differently
for n in $(seq 3 10)
do
  plot_args="$plot_args, '$2' u 1:$n notitle"
done
plot_args="$plot_args, '$2' u 1:$(($N_COLS-2)) w l title 'Mean'"
plot_args="$plot_args, '$2' u 1:$(($N_COLS-1)) w l title 'Low'"
plot_args="$plot_args, '$2' u 1:$N_COLS w l title 'High'"

gnuplot <<EOF
set terminal png
set output "$1"
set xlabel "Esclaves"
set xtics -1,1,$MAX_SLAVES
set xr [-1*$X_PADDING:$MAX_SLAVES+$X_PADDING]
set ylabel "$3"
plot $plot_args
EOF
