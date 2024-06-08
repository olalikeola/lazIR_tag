import numpy as np
import os
import pandas as pd
import time
import tqdm 
import multiprocessing
import argparse

parser = argparse.ArgumentParser()
parser.add_argument('n_players', type=int, help='num of players')
args = parser.parse_args()

def run_cmd_get_output(cmd):
    return os.popen(cmd).read()

n_players = args.n_players

op = {'n_players': n_players}

#assuming that the cluster is already running on the r/pi nodes
for num_players in range(1, n_players + 1):
    cmd = 'go run sim.go'
    
    t = []
    for _ in tqdm.tqdm(range(20)):
        #spawn threads which runs the cmd in parallel 20 times
        outputs = []

        for p in range(n_players):
            os.system(cmd + '> /tmp/raftexp' + str(p) + '.out 2>&1 &')
        
        
        os.system('sleep 1')
        while True:
            try:
                for p in range(n_players):
                    outputs += [run_cmd_get_output('cat /tmp/raftexp' + str(p) + '.out')]
                for output in outputs:
                    t += [float(output.split('\n')[-2].split(': ')[1].strip()[:-2])]
                break
            except Exception as e:
                outputs = []
                os.system('sleep 1')


        

    t = np.mean(np.array(t))
    op[str(i)] = t
    print((i, t))

print("results: ")
print(op)



