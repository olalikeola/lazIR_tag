n_nodes=10
for n_players in 3 4 5 6 7 8 9 10 20 25 50
do
    python script/run_experiments.py --n_players $n_players --n_nodes $n_nodes 
    python script/run_experiments.py --n_nodes $n_nodes --n_players $n_players --node_failure
done
