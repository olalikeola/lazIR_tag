import os
import numpy as np
import tqdm 

op = []
for i in range(3, 11):
    #start raft cluster

    os.system('python3 script/start_raft_cluster.py ' + str(i))

    #run sim
    os.chdir('src/sim')

    #get output of cmd
    cmd = 'go run sim.go'
    
    t = []
    for _ in tqdm.tqdm(range(10)):
        output = os.popen(cmd).read()
        t += [float(output.split('\n')[-2].split(': ')[1].strip()[:-2])]
    t = np.mean(np.array(t))
    op += [(i, t)]

    os.system('kill $(pgrep -f "\-\-config")')
    os.chdir('../..')
    os.system('sleep 1')

print("results: ")
print(op)

