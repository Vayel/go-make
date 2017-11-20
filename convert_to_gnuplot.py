import sys
import json


TIMES_FNAME = 'times.txt'
SPEEDUPS_FNAME = 'speedups.txt'
EFFICIENCIES_FNAME = 'efficiencies.txt'

def help():
    print('Usage: python convert_to_gnuplot.py <measures.json>')


def mean(l):
    return sum(l) / len(l)

def save_values(fname, values):
    with open(fname, 'w') as f:
        for n, line in sorted(values.items()):
            f.write(
                str(n) + ' ' +
                ' '.join((str(x) for x in line)) + ' ' +
                str(mean(line)) + '\n'
            )


def json_to_matrix(f):
    return {int(k): v for k, v in json.load(f).items()}


def build_speedups(measures, seq_time):
    values = {}
    for n, times in measures.items():
        if not n:
            continue
        values[n] = [seq_time / val for val in times]
    return values


def build_efficiencies(speedups):
    values = {}
    for n, times in speedups.items():
        if not n:
            continue
        values[n] = [val / n for val in times]
    return values


if __name__ == '__main__':
    try:
        path = sys.argv[1]
    except IndexError:
        help()
        sys.exit(1)

    try:
        f = open(path)
    except FileNotFoundError:
        help()
        sys.exit(1)
    else:
        times = json_to_matrix(f)
    finally:
        f.close()

    save_values(TIMES_FNAME, times)
    seq_time = mean(times[0])
    speedups = build_speedups(times, seq_time)
    save_values(SPEEDUPS_FNAME, speedups)
    save_values(EFFICIENCIES_FNAME, build_efficiencies(speedups))
