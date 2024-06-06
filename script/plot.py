import matplotlib.pyplot as plt

arr = [(3, 21.3905166), (4, 25.8317541), (5, 25.4648125), (6, 26.184533199999997), (7, 28.828495799999995), (8, 28.3030125), (9, 30.6698999), (10, 33.2847751)]

x = [i[0] for i in arr]
y = [i[1] for i in arr]

plt.figure(figsize=(10, 6))
plt.plot(x, y, marker='o', linestyle='-', color='b', markersize=8, linewidth=2, markerfacecolor='r', markeredgewidth=2)

plt.title('Data Trend', fontsize=16)
plt.xlabel('Clusters', fontsize=14)
plt.ylabel('Latency', fontsize=14)

plt.grid(True, which='both', linestyle='--', linewidth=0.5)
plt.minorticks_on()

plt.xticks(fontsize=12)
plt.yticks(fontsize=12)

plt.show()
