import sys
import os
import json
import subprocess
import glob
import threading
import queue
import time
import shutil


LOG_DIR = '~/go-make/logs'
NODES_DIR = '~/go-make/logs/nodes'
RESULT_PATH = os.path.expanduser('~/measures.json')
SEQ_LOG = os.path.join(LOG_DIR, 'sequential.json')
SEQ_NODE = os.path.join(NODES_DIR, 'sequential.txt')
MASTER_LOG = os.path.join(LOG_DIR, 'master.json')
SLAVE_LOG_PATTERN = os.path.join(LOG_DIR, 'slave_*.json')
FIRST_RULE = 'all'

os.makedirs(os.path.expanduser(LOG_DIR), exist_ok=True)

def help():
    print('Usage:')
    print('\tpython3 test.py min_n_slaves max_n_slaves n_slaves_step n_reps makefile-path')
    print('Examples:')
    print('\tpython3 test.py 14 29 5 5 ~/go-make/makefiles/xxx')


def launch_master(q, makefile_path, rule):
    ret = subprocess.call(["./launch_master.sh", makefile_path, rule, MASTER_LOG])
    if ret: # Error
        return
    with open(os.path.join(LOG_DIR, 'time_master.json')) as f:
        data = json.load(f)
        q.put(data['total'])


def run_para(n_slaves, makefile_path):
    print('Running parallel with {} slaves...'.format(n_slaves), end='\n\n')
    q = queue.Queue() # allows to get return value of the thread
    threading.Thread(target=launch_master, args=[q, makefile_path]).start()
    time.sleep(3) # wait for master to start
    subprocess.Popen(["./launch_slave.sh",  str(n_slaves)])
    return q.get()


def run_seq(makefile_path):
    print('Running sequential', end='\n\n')

    ret = subprocess.call([
        "./sequential.sh",
        makefile_path,
        FIRST_RULE,
        SEQ_LOG,
        SEQ_NODE,
    ])
    if ret: # Error
        return

    with open(os.path.expanduser(SEQ_LOG)) as f:
        mes = json.load(f)

    return mes['total']


def clean_logs():
    print('Cleaning log dir...')
    for f in glob.glob(os.path.expanduser(os.path.join(LOG_DIR, "*.json"))):
        os.remove(f)
    for f in glob.glob(os.path.expanduser(os.path.join(NODES_DIR, "*.txt"))):
        os.remove(f)


if __name__ == '__main__':
    try:
        MIN_N_SLAVES = int(sys.argv[1])
        MAX_N_SLAVES = int(sys.argv[2])
        assert MAX_N_SLAVES >= MIN_N_SLAVES
        assert MIN_N_SLAVES > 0

        N_SLAVES_STEP = int(sys.argv[3])
        assert N_SLAVES_STEP > 0

        N_REPS = int(sys.argv[4])
        assert N_REPS >= 1

        MAKEFILE = sys.argv[5]
    except Exception as e:
        print('Error with parameters {}'.format(sys.argv[1:]))
        help()
        print('')
        raise e

    measures = {}

    clean_logs()
    measures[0] = [run_seq(MAKEFILE) for _ in range(N_REPS)]
    # We save the results gradually in the case where an error appears during
    # the computations
    with open(RESULT_PATH, 'w') as f:
        json.dump(measures, f, indent=4)


    sys.exit(0) # TODO

    for n_slaves in range(MIN_N_SLAVES, MAX_N_SLAVES + 1, N_SLAVES_STEP):
        clean_logs()
        mes = [run_para(n_slaves, MAKEFILE) for _ in range(N_REPS)]
        n = len(glob.glob(os.path.expanduser(SLAVE_LOG_PATTERN)))
        if n != n_slaves:
            print('Invalid number of slave log files ({0} instead of {1})'.format(n, n_slaves))
        measures[n_slaves] = mes
        with open(RESULT_PATH, 'w') as f:
            json.dump(measures, f, indent=4)
