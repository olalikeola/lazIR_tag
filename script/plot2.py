import matplotlib.pyplot as plt
import pandas as pd
import numpy as np

# Read the CSV file
df = pd.read_csv("sim_results/node_Failure=True.csv")

# Get a continuous colormap
cmap = plt.colormaps['plasma']
normalize = plt.Normalize(vmin=0, vmax=len(df)-1)

# Set the figure size
plt.figure(figsize=(10,6))

# Iterate through each row in the dataframe and plot
for i, row in enumerate(df.iterrows()):
    n_players = row[1][0]
    data = row[1][1:]
    x = np.array(list(map(int, data.index)))  # Correct the x-axis values
    y = np.array(data)

    plt.plot(x, y, marker='o', linestyle='-',
        color=cmap(normalize(i)), markersize=8, linewidth=2,
        markerfacecolor=cmap(normalize(i)), markeredgewidth=2, label=f'{int(n_players)}'
    )
    print('plotted this!')

# Set the title and labels with bold font
plt.title('With Node Failures', fontsize=16, fontweight='bold')
plt.xlabel('Number of Raft Nodes', fontsize=14, fontweight='bold')
plt.ylabel('Latency (ms)', fontsize=14, fontweight='bold')

# Add grid lines
plt.grid(True , linewidth=0.1)
plt.minorticks_on()

# Set the tick font sizes
plt.xticks(fontsize=12)
plt.yticks(fontsize=12)

# Add a legend with a smaller font size
plt.legend(fontsize=10, title='Number of Players', title_fontsize='12', loc='best')

# Customize font to look more publication-ready
plt.rcParams.update({'font.family': 'serif', 'font.serif': 'Times New Roman'})

# Show the plot
plt.show()

