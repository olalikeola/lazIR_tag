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

	conn, err := grpc.Dial("multi:///localhost:5001,localhost:5002,localhost:5003",
		grpc.WithDefaultServiceConfig(serviceConfig), grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(retryOpts...)))

	if err != nil {
		log.Fatalf("dialing failed: %v", err)
	}

	ctx := context.Background()

	defer conn.Close()
	c := pb.NewScoreClient(conn)

	fmt.Println("Varun just shot Rahul!")
	c.AddRecord(ctx, &pb.AddRecordRequest{Shooter: "node_name1", Victim: "node_name2"})

	fmt.Println("Varun just shot Rahul, again!")
	c.AddRecord(ctx, &pb.AddRecordRequest{Shooter: "node_name1", Victim: "node_name3"})

	fmt.Println("Varun just shot Rahul, againx2!")
	c.AddRecord(ctx, &pb.AddRecordRequest{Shooter: "node_name2", Victim: "node_name3"})

	//sleep for 1 seconds
	time.Sleep(1 * time.Second)

	resp, err := c.GetRecord(ctx, &pb.GetRecordRequest{})
	fmt.Println(resp)
}
