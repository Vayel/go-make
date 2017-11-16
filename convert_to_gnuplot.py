import sys
import json


TIMES_FNAME = 'times.txt'
SPEEDUPS_FNAME = 'speedups.txt'
EFFICIENCIES_FNAME = 'efficiencies.txt'

def help():
    print('Usage: python convert_to_gnuplot.py <measures.json>')


def save_matrix(fname, matrix):
    with open(fname, 'w') as f:
        for n, line in enumerate(matrix):
            f.write(str(n) + ' ' + ' '.join((str(x) for x in line)) + '\n')


def json_to_matrix(f):
    data = json.load(f)
    return [d['measures'] for _, d in sorted(data.items())]


def build_speedups(times, seq_time):
    for line in times:
        for i, val in enumerate(line):
            line[i] = seq_time / val
    return times


def build_efficiencies(speedups):
    for n, line in enumerate(times):
        if not n:
            continue
        for i, val in enumerate(line):
            line[i] = val / n
    return speedups


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

    save_matrix(TIMES_FNAME, times)
    seq_time = sum(times[0]) / len(times[0])
    save_matrix(SPEEDUPS_FNAME, build_speedups(times, seq_time))
    save_matrix(EFFICIENCIES_FNAME, build_efficiencies(times))
