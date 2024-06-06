import os
import sys
import re
import matplotlib.pyplot as plt
from matplotlib.animation import FuncAnimation
import matplotlib.cm as cm
import numpy as np

current_read_index = 0


def update(frame):
    global current_read_index
    line = sys.stdin.readline()
    if line.startswith("map"):
        scores = { l.split(":")[0]: int(l.split(":")[1]) for l in line[len("map["):-2].strip().split(" ") }
        print(scores)

        plt.cla()
        
        # Use a color map
        colors = cm.viridis(np.linspace(0, 1, len(scores)))

        # Plot the bar chart
        bars = plt.bar(scores.keys(), scores.values(), color=colors, width=0.6)

        # Add value labels on top of the bars
        for bar in bars:
            yval = bar.get_height()
            plt.text(bar.get_x() + bar.get_width()/2.0, yval, int(yval), va='bottom', ha='center')

        plt.xticks(rotation=45, ha='right')
        plt.xlabel("Players", fontsize=12)
        plt.ylabel("Life Points", fontsize=12)
        plt.grid(True, linestyle='--', alpha=0.7)
        plt.tight_layout()

        # Add main title and subtext
        plt.suptitle("LazIR tag", fontsize=25, fontweight='bold', fontname='Phosphate')
        plt.figtext(0.95, 0.05, f"idx: {current_read_index}", ha='center', fontsize=8, style='italic')
    else:
        current_read_index = int(line.strip())
        print(current_read_index)

ani = FuncAnimation(plt.gcf(), update, interval=10)

plt.show()
