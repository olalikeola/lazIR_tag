import os
import numpy as np
import tqdm 
import multiprocessing

def run_cmd_get_output(cmd):
    return os.popen(cmd).read()

op = []
n_players = 25
for i in range(3, 11):
    #start raft cluster

    os.system('python3 script/start_raft_cluster.py ' + str(i) )

    os.system('python3 script/reelection_trigger.py ' + str(i) + ' 1 > /tmp/raft_reelec 2>&1 &')

    #run sim
    os.chdir('/Users/varund/Desktop/q3/cs244b/project/lazIR_tag_final_final/src/sim')

    #get output of cmd
    cmd = 'go run sim.go'
    
    t = []
    for _ in tqdm.tqdm(range(2)):
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
    op += [(i, t)]
    print((i, t))

    os.system('kill $(pgrep -f "\-\-config")')
    os.system('kill $(pgrep -f "reelection_trigger.py")')
    os.chdir('../..')
    os.system('sleep 1')

print("results: ")
print(op)

