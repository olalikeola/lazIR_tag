// Binary hammer sends requests to your Raft cluster as fast as it can.
// It sends the written out version of the Dutch numbers up to 2000.
// In the end it asks the Raft cluster what the longest three words were.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Darkhunter9/lazIR_tag/src/proto"
	_ "github.com/Jille/grpc-multi-resolver"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/health"
)

func main() {

	serviceConfig := `{"healthCheckConfig": {"serviceName": "Score"}, "loadBalancingConfig": [ { "round_robin": {} } ]}`
	retryOpts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(100 * time.Millisecond)),
		grpc_retry.WithMax(5),
	}

	conn, err := grpc.Dial("multi:///localhost:5001,localhost:5002,localhost:5003,localhost:5004,localhost:5005,localhost:5006,localhost:5007,localhost:5008,localhost:5009,localhost:5010,localhost:5011",
		grpc.WithDefaultServiceConfig(serviceConfig), grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(retryOpts...)))

	if err != nil {
		log.Fatalf("dialing failed: %v", err)
	}

	ctx := context.Background()

	defer conn.Close()
	c := pb.NewScoreClient(conn)

	resp, err := c.GetRecord(ctx, &pb.GetRecordRequest{})

	if err != nil {
		log.Fatalf("failed to get record: %v", err)
	}

	score_map := resp.Score

	score_node_name_1 := score_map["node_name1"]
	score_node_name_2 := score_map["node_name2"]

	target_score_node_name_1 := score_node_name_1 + 1
	target_score_node_name_2 := score_node_name_2 - 1

	fmt.Println("score_node_name_1: ", score_node_name_1)
	fmt.Println("score_node_name_2: ", score_node_name_2)
	fmt.Println("target_score_node_name_1: ", target_score_node_name_1)
	fmt.Println("target_score_node_name_2: ", target_score_node_name_2)

	// get current time
	current_time := time.Now()
	c.AddRecord(ctx, &pb.AddRecordRequest{Shooter: "node_name1", Victim: "node_name2"})

	//submit request for getting shot

	for {
		resp, err := c.GetRecord(ctx, &pb.GetRecordRequest{})

		if err == nil {
			score_map = resp.Score
			if score_map["node_name1"] == target_score_node_name_1 && score_map["node_name2"] == target_score_node_name_2 {
				result_time := time.Now()
				fmt.Println("Time taken to get shot: ", result_time.Sub(current_time))
				break
			} else {
				fmt.Println("Waiting for shot...")
			}
		} else {
			fmt.Println(err)
		}
	}

}
