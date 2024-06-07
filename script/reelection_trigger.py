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



os.chdir('config')
leader_config = "config_bootstrap.yaml"
config_files.remove(leader_config)
other_config = config_files[0]
config_files.remove(other_config)

addr_list = []
names = []
addr_list += [get_addr_port(leader_config), get_addr_port(other_config)]
names += [get_name(leader_config), get_name(other_config)]
first_cmd = "raftadmin " + addr_list[0] + " add_voter " + names[1] + " " + addr_list[1] + " 0"

while (len(sys.argv) > 2):
    time.sleep(reelection_int)
    print('********REELECTION**********')
    
    cmd = "raftadmin --leader multi:///"+ ",".join(addr_list) + " leadership_transfer"
    print(cmd)
    os.system(cmd)
