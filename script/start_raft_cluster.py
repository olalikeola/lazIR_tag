import os
import sys
import yaml
import time
import random 
def get_dict_from_yaml(file):
    with open(file, 'r') as stream:
        return yaml.load(stream, Loader = yaml.FullLoader)

def get_addr_port(file):
    config = get_dict_from_yaml(file)
    return str(config['address']) + ":" + str(config['port'])

def get_name(file):
    config = get_dict_from_yaml(file)
    return str(config['raftId'])

def get_other_config(name):
    if(name == "node_name1"):
        return "config_bootsrap.yaml"
    elif name == "node_name2":
        return "config_node.yaml"
    else:
        number = int(name[-1])
        return f"config_node{number-1}.yaml"
    
config_files = os.listdir('config')

n_servers = 6
if(len(sys.argv) > 1):
    n_servers = int(sys.argv[1])

if(len(sys.argv) > 2):
    failure_prop = float(sys.argv[2])
    
config_files.remove("config_bootstrap.yaml")
config_files = ["config_bootstrap.yaml"] + config_files[:n_servers-1]

#clear old files
os.system('rm -rf ~/data/*')

#start raft cluster, one node at a time
os.chdir('src/raft_grpc')
for config_file in config_files:
    cmd = 'go run *.go --config ../../config/' + config_file + ' > ' + 'out/' + get_name('../../config/' + config_file) + '.out' + ' 2>&1 &'
    print("running: " + cmd)
    os.system(cmd)

#sleep for 10 seconds
time.sleep(5)
os.chdir('../../config')
leader_config = "config_bootstrap.yaml"
config_files.remove(leader_config)
other_config = config_files[0]
config_files.remove(other_config)

addr_list = []
names = []
addr_list += [get_addr_port(leader_config), get_addr_port(other_config)]
names += [get_name(leader_config), get_name(other_config)]
first_cmd = "raftadmin " + addr_list[0] + " add_voter " + names[1] + " " + addr_list[1] + " 0"
input()
print("running: " + first_cmd)
os.system(first_cmd)
while(len(config_files)):
    config_file = config_files[0]
    addr = get_addr_port(config_file)
    name = get_name(config_file)

    cmd = "raftadmin --leader multi:///" + ",".join(addr_list) + " add_voter " + name + " " + addr + " 0"
    print("running: " + cmd)
    os.system(cmd)
    addr_list.append(addr)
    names.append(name)
    config_files.remove(config_file)



# upon catch ctrl+c kill all children and exit
def signal_handler(sig, frame):
    os.system('kill $(pgrep -f \-\-config')
    print(f'{reelection_count=}')
    sys.exit(0)

alive_status = [True] * n_servers

os.chdir('../src/raft_grpc')

if(failure_prop != -1):
    while(True):
        time.sleep(0.1)


        for i in range(n_servers):
            coin_toss = random.random()
            n_failed = sum([1 for a in alive_status if alive_status[a]])

            if(coin_toss < failure_prop and alive_status[i]):
                if(2 * n_failed  + 1 < n_servers):
                    print(f"Killing {names[i]}")
                    cmd = f'kill $(pgrep -f "\-\-config .*{get_other_config(names[i])}")'
                    cmd_pgrep = f'pgrep -f "\-\-config .*{get_other_config(names[i])}"'
                    os.system(cmd_pgrep)
                    print("running: " + cmd )
                    os.system(cmd)
                    alive_status[i] = False

            elif(coin_toss > failure_prop and not alive_status[i]):
                print(f"Starting {names[i]}")
                os.system(f"go run *.go --config ../../config2/{get_other_config(names[i])} > {names[i]}.out 2>&1 &")

                alive_status[i] = True
