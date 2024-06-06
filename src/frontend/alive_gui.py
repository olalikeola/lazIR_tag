import os
import sys
import re
import matplotlib.pyplot as plt
from matplotlib.animation import FuncAnimation
import matplotlib.cm as cm
import numpy as np

# List of IP addresses to monitor
ips = ["192.168.8.160", "192.168.8.167", "192.168.8.103", "192.168.8.151"]
ping_cmd = ["ping -c 1 -t 5 " + ip for ip in ips]

# Function to check the alive status of IPs
def get_alive_status():
    ping_resp = [os.popen(cmd).read() for cmd in ping_cmd]
    alive = ['1 packets received' in p for p in ping_resp]
    return alive

# Setup plot
fig, ax = plt.subplots()
bars = ax.bar(ips, [0] * len(ips), color='blue')
ax.set_ylim(0, 1)
ax.set_ylabel('Alive Status')
ax.set_title('IP Alive Status Monitoring')

# Update function for animation
def update(frame):
    alive = get_alive_status()
    for i, bar in enumerate(bars):
        bar.set_height(1 if alive[i] else 0)
        bar.set_color('green' if alive[i] else 'red')
    return bars

# Create animation
ani = FuncAnimation(fig, update, frames=np.arange(0, 100), interval=1000, blit=False)

# Show plot
plt.show()
