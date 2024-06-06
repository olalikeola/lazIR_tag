import os
import sys
import re
import matplotlib.pyplot as plt
from matplotlib.animation import FuncAnimation
import matplotlib.cm as cm
import numpy as np

current_read_index = 0

ips = ["192.168.8.160", "192.168.8.167", "192.168.8.103", "192.168.8.151"]
ping_cmd = ["ping -c 1 -t 1 " + ip for ip in ips]

def get_alive_status():
    ping_resp = [os.popen(cmd).read() for cmd in ping_cmd]
    alive = ['1 packets received' in p for p in ping_resp]
    return alive

while(True):
    alive = get_alive_status()
    for i in range(len(ips)):
        print(f"{ips[i]} is {'alive' if alive[i] else 'dead'}")
    os.system('sleep 1')





