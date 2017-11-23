import sys
import json
import math
import statistics as stats


TIMES_FNAME = None
SPEEDUPS_FNAME = 'speedups.txt'
EFFICIENCIES_FNAME = 'efficiencies.txt'
STUDENT_TABLE = {3: 2.92, 5: 2.132}

class StudentTableError(ValueError): pass

def help():
    print('Usage:')
    print('\tpython3 convert_to_gnuplot.py <measures.json> <output_path>')


def error(l):
    try:
        return STUDENT_TABLE[len(l)] * stats.stdev(l) / math.sqrt(len(l))
    except KeyError:
        raise StudentTableError('The Student table has no entry for {} repetitions'.format(len(l)))


def save_values(fname, values):
    with open(fname, 'w') as f:
        for n, line in sorted(values.items()):
            err = error(line)
            m = stats.mean(line)
            f.write(
                str(n) + ' ' +
                ' '.join((str(x) for x in line)) + ' ' +
                str(m) + ' ' +
                str(m - err) + ' ' +
                str(m + err) + '\n'
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
        TIMES_FNAME = sys.argv[2]
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

    try:
        save_values(TIMES_FNAME, times)
    except StudentTableError as e:
        print('Error:', e)
        sys.exit(1)
    seq_time = stats.mean(times[0])
    speedups = build_speedups(times, seq_time)
    try:
        save_values(SPEEDUPS_FNAME, speedups)
        save_values(EFFICIENCIES_FNAME, build_efficiencies(speedups))
    except StudentTableError as e:
        print('Error:', e)
        sys.exit(1)
