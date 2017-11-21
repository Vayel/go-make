import sys
import os
import json
import subprocess
import glob
import threading
import queue
import time


MIN_N_SLAVES_ = 1  # Le sujet dit: "Les tests devront etre realises sur un minimum de 15 machines"
LOG_DIR = os.path.expanduser('~/go-make/logs')
RESULT_PATH = os.path.join(LOG_DIR, 'time_measures.json')
SEQ_LOGS = os.path.join(LOG_DIR, 'seq.log')
MASTER_LOGS = os.path.join(LOG_DIR, 'master.log')
SLAVE_LOGS = os.path.join(LOG_DIR, 'slave.log')

def help():
    print('python3 test.py min_n_slaves max_n_slaves n_slaves_step n_reps makefile-path')
    print('Examples:')
    print('\tpython3 test.py 14 29 5 5 /tmp/go-make/makefiles/xxx')


def launch_master(q, makefile_path):
    ret = subprocess.call(["./launch_master.sh", makefile_path])
    if ret: # Error
        return
    with open(os.path.join(LOG_DIR, 'time_master.json')) as f:
        data = json.load(f)
        q.put(data['total'])


def run_para(n_slaves, makefile_path):
    print('Run parallel with {} slaves'.format(n_slaves), end='\n\n')
    q = queue.Queue() # allows to get return value of the thread
    threading.Thread(target=launch_master, args=[q, makefile_path]).start()
    time.sleep(3) # wait for master to start
    subprocess.Popen(["./launch_slave.sh",  str(n_slaves)])
    return q.get()


def run_seq(makefile_path):
    print('Run sequential', end='\n\n')
    ret = subprocess.call(["./launch_sequential.sh", makefile_path])
    if ret: # Error
        return
    with open(os.path.join(LOG_DIR, 'time_seq.json')) as f:
        mes = json.load(f)
    return mes['total']


if __name__ == '__main__':
    try:
        MIN_N_SLAVES = int(sys.argv[1])
        MAX_N_SLAVES = int(sys.argv[2])
        assert MAX_N_SLAVES >= MIN_N_SLAVES
        assert MIN_N_SLAVES >= MIN_N_SLAVES_
        N_SLAVES_STEP = int(sys.argv[3])
        assert N_SLAVES_STEP > 0
        N_REPS = int(sys.argv[4])
        assert N_REPS >= 1
        MAKEFILE = sys.argv[5]
    except:
        help()
        sys.exit(1)

    measures = {}
    
    measures[0] = [run_seq(MAKEFILE) for _ in range(N_REPS)]
    with open(RESULT_PATH, 'w') as f:
        json.dump(measures, f, indent=4)

    for n_slaves in range(MIN_N_SLAVES, MAX_N_SLAVES + 1, N_SLAVES_STEP):
        measures[n_slaves] = [run_para(n_slaves, MAKEFILE) for _ in range(N_REPS)]
        with open(RESULT_PATH, 'w') as f:
            json.dump(measures, f, indent=4)
