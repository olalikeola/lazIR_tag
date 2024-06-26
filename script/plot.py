import matplotlib.pyplot as plt

arr = [(3, 21.3905166), (4, 25.8317541), (5, 25.4648125), (6, 26.184533199999997), (7, 28.828495799999995), (8, 28.3030125), (9, 30.6698999), (10, 33.2847751)]

arr_10ply = [(3, 22.081173319999998), (4, 25.758135820000003), (5, 26.21421249), (6, 28.580066189999997), (7, 29.504255809999997), (8, 32.44659962000001), (9, 34.09108044), (10, 36.823127099999994)]

arr_25ply = [(3, 32.18032268723533), (4, 29.357745390198513), (5, 29.703881861189803), (6, 32.85647831199502), (7, 36.035921268762024), (8, 35.73410809615385), (9, 36.88802092961373), (10, 42.99792754715909)]

x = [i[0] for i in arr]
y = [i[1] for i in arr]

x_10ply = [i[0] for i in arr_10ply]
y_10ply = [i[1] for i in arr_10ply]

x_25ply = [i[0] for i in arr_25ply]
y_25ply = [i[1] for i in arr_25ply]

plt.figure(figsize=(10, 6))

# Plot for arr
plt.plot(x, y, marker='o', linestyle='-', color='b', markersize=8, linewidth=2, markerfacecolor='r', markeredgewidth=2, label='1 player')

# Plot for arr_10ply with a different style
plt.plot(x_10ply, y_10ply, marker='s', linestyle='--', color='g', markersize=8, linewidth=2, markerfacecolor='y', markeredgewidth=2, label='10 players')

plt.plot(x_25ply, y_25ply, marker='s', linestyle='--', color='r', markersize=8, linewidth=2, markerfacecolor='y', markeredgewidth=2, label='25 players')

plt.title('Data Trend', fontsize=16)
plt.xlabel('Clusters', fontsize=14)
plt.ylabel('Latency', fontsize=14)

plt.grid(True, which='both', linestyle='--', linewidth=0.5)
plt.minorticks_on()

plt.xticks(fontsize=12)
plt.yticks(fontsize=12)

plt.legend(fontsize=12)
plt.show()
