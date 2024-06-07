import os
import sys
import yaml
import time

def get_dict_from_yaml(file):
    with open(file, 'r') as stream:
        return yaml.load(stream, Loader = yaml.FullLoader)

def get_addr_port(file):
    config = get_dict_from_yaml(file)
    return str(config['address']) + ":" + str(config['port'])

def get_name(file):
    config = get_dict_from_yaml(file)
    return str(config['raftId'])
config_files = os.listdir('config')

n_servers = 6
if(len(sys.argv) > 1):
    n_servers = int(sys.argv[1])

if (len(sys.argv) > 2):
    reelection_int = float(sys.argv[2])
    
config_files.remove("config_bootstrap.yaml")
config_files = ["config_bootstrap.yaml"] + config_files[:n_servers-1]

#clear old files
os.system('rm -rf ~/data/*')

#start raft cluster, one node at a time
os.chdir('src/raft_grpc')
for config_file in config_files:
    cmd = 'go run *.go --config ../../config/' + config_file + ' > ' + get_name('../../config/' + config_file) + '.out' + ' 2>&1 &'
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
