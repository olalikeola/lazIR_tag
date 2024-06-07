import numpy as np
import os
import pandas as pd
import time
import tqdm 
import multiprocessing
import argparse

parser = argparse.ArgumentParser()
parser.add_argument('--n_players', type=int, required=True, help='num of players')
parser.add_argument('--n_nodes', type=int, required=True, help='num nodes')
parser.add_argument('--node_failure', action='store_true')
args = parser.parse_args()

def run_cmd_get_output(cmd):
    return os.popen(cmd).read()

n_players = args.n_players
n_nodes = args.n_nodes
node_failure = args.node_failure

op = {'n_players': n_players}

for i in range(3, n_nodes + 1):
    #start raft cluster

    cmd = 'python3 script/start_raft_cluster.py ' + str(i)
    if node_failure:
        cmd += ' 0.1'
    else:
        cmd += ' -1'
    cmd += ' > /tmp/raftexp.out 2>&1 &'
    print("running: " + cmd)
    os.system(cmd)
    time.sleep(10)

    

    #run sim
    os.chdir('/Users/pinocholsaipant/Documents/CS244b/lazIR_tag/src/sim')

    #get output of cmd
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
    # os.system('pgrep -f "\-\-config"')
    print("first kill")
    os.system('kill $(pgrep -f "\-\-config")')
    print("second kill")
    os.system('kill $(pgrep -f "start_raft_cluster")')
    os.chdir('../..')
    os.system('sleep 1')

print("results: ")
print(op)

res = pd.read_csv(f'script/sim_results/{node_failure=}.csv', index_col=False)
res.loc[len(res)] = op
res.to_csv(f'script/sim_results/{node_failure=}.csv', index=False)


