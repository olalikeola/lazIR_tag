import sys
import re
import matplotlib.pyplot as plt
from matplotlib.animation import FuncAnimation
import numpy as np

current_read_index = 0

# Create the figure and subplots only once
fig, ax = plt.subplots(2, 2)
fig.suptitle("LazIR tag", fontsize=25, fontweight='bold', fontname='Phosphate')

# Define background colors for each subplot
background_colors = ['#FFDDC1', '#FFC3A0', '#FFABAB', '#FF677D']

cmap = plt.colormaps['plasma']
cmap2 = plt.colormaps['inferno']
normalize = plt.Normalize(vmin=0, vmax=3)


def update(frame):
    global current_read_index
    line = sys.stdin.readline()
    if line.startswith("map"):
        scores = {l.split(":")[0]: int(l.split(":")[1]) for l in line[len("map["):-2].strip().split(" ")}
        print(scores)
        players = list(scores.keys())
        points = list(scores.values())
        players = ["Zihao", "Varun", "Ola", "Pino"]
        points = [10, 15, 7, 6] 
        for i in range(4):
            x, y = i // 2, i % 2
            ax[x][y].cla()  # Clear the specific subplot
            ax[x][y].set_title(players[i], fontsize=14, fontweight='bold', fontname='Menlo', color = cmap(normalize(0)))
            
            ax[x][y].text(0.5, 0.5, points[i], va='center', 
                          ha='center', fontsize=68, 
                          fontweight='bold', color = cmap2(1 - normalize(0))
            )
            ax[x][y].set_xticks([])
            ax[x][y].set_yticks([])
            ax[x][y].patch.set_facecolor(cmap2(normalize(0)))  # Set background color
        
        
        # Adjust layout
        fig.tight_layout(rect=[0, 0.03, 1, 0.95])

        # Add index text
        current_read_index = 104
        fig.text(0.95, 0.02, f"idx: {current_read_index}", ha='center', fontsize=8, style='italic')

    else:
        current_read_index = int(line.strip())
        print(current_read_index)

ani = FuncAnimation(fig, update, interval=10)
plt.show()

