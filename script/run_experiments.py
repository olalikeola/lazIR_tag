import os
import numpy as np
import tqdm 
import multiprocessing

def run_cmd_get_output(cmd):
    return os.popen(cmd).read()

op = []
for i in range(3, 11):
    #start raft cluster

    os.system('python3 script/start_raft_cluster.py ' + str(i))

    #run sim
    os.chdir('/Users/varund/Desktop/q3/cs244b/project/lazIR_tag_final_final/src/sim')

    #get output of cmd
    cmd = 'go run sim.go'
    
    t = []
    for _ in tqdm.tqdm(range(10)):
        #spawn threads which runs the cmd in parallel 20 times
        outputs = []

        for p in range(25):
            os.system(cmd + '> /tmp/raftexp' + str(p) + '.out 2>&1 &')
        
        
        os.system('sleep 5')
        for p in range(25):
            outputs += [run_cmd_get_output('cat /tmp/raftexp' + str(p) + '.out')]        

        for output in outputs:
            print("output: " + output)
            t += [float(output.split('\n')[-2].split(': ')[1].strip()[:-2])]

    t = np.mean(np.array(t))
    op += [(i, t)]

    os.system('kill $(pgrep -f "\-\-config")')
    os.chdir('../..')
    os.system('sleep 1')

print("results: ")
print(op)

