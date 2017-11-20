import sys
import os
import json
import subprocess
import glob
import threading
import queue


MIN_N_SLAVES = 14  # Le sujet dit: "Les tests devront etre realises sur un minimum de 15 machines"
LOG_DIR = os.path.expanduser('~/go-make/logs')
RESULT_PATH = os.path.join(LOG_DIR, 'time_measures.json')

def help():
    print('python3 test.py max_n_slaves n_slaves_step n_reps')
    print('Examples:')
    print('\tpython3 test.py 29 5 5')


def launch_master(q):
    proc = subprocess.run(["./launch_master.sh"])
    res = []
    for f in glob.glob(os.path.join(LOG_DIR, 'time_*.json')):
        with open(f) as logfile:
            res.append((f, json.load(logfile)))
    q.put(res)


def run_para(n_slaves):
    q = queue.Queue() # allows to get return value of the thread
    threading.Thread(target=launch_master, args=[q]).start()
    subprocess.Popen(["./launch_slave.sh",  str(n_slaves)])
    return q.get()


def run_seq():
    subprocess.run("./launch_sequential.sh")
    with open(os.path.join(LOG_DIR, 'time_seq.json')) as f:
        mes = json.load(f)
    return mes


if __name__ == '__main__':
    MAX_N_SLAVES = int(sys.argv[1])
    assert MAX_N_SLAVES >= MIN_N_SLAVES
    N_SLAVES_STEP = int(sys.argv[2])
    assert N_SLAVES_STEP > 0
    N_REPS = int(sys.argv[3])
    assert N_REPS > 1

    subprocess.run("./nodes_configuration.sh")
    measures = {}
    
    measures[0] = [run_seq() for _ in range(N_REPS)]

    for n_slaves in range(MIN_N_SLAVES, MAX_N_SLAVES + 1, N_SLAVES_STEP):
        measures[n_slaves] = [run_para(n_slaves) for _ in range(N_REPS)]

    with open(RESULT_PATH, 'w') as f:
        json.dump(measures, f)
