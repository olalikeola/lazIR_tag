package main

import "C"

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

func client(dns string) (*pb.ScoreClient, *grpc.ClientConn, error) {
	serviceConfig := `{"healthCheckConfig": {"serviceName": "Score"}, "loadBalancingConfig": [ { "round_robin": {} } ]}`
	retryOpts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(100 * time.Millisecond)),
		grpc_retry.WithMax(5),
	}

	conn, err := grpc.Dial(dns,
		grpc.WithDefaultServiceConfig(serviceConfig), grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(retryOpts...)))

	if err != nil {
		log.Fatalf("dialing failed: %v", err)
	}

	c := pb.NewScoreClient(conn)

	return &c, conn, nil
}

//export addRecord
func addRecord(dns string, shooter string, victim string) error {
	c, conn, err1 := client(dns)
	defer conn.Close()

	if err1 != nil {
		return fmt.Errorf("client: %v", err1)
	}

	cli := *c
	_, err2 := cli.AddRecord(
		context.Background(),
		&pb.AddRecordRequest{Shooter: shooter, Victim: victim},
	)

	if err2 != nil {
		return fmt.Errorf("addRecord: %v", err2)
	}

	return nil
}

//export getRecord
func getRecord(dns string) (map[string]int32, error) {
	c, conn, err1 := client(dns)
	defer conn.Close()

	if err1 != nil {
		return nil, fmt.Errorf("client: %v", err1)
	}

	cli := *c
	res, err2 := cli.GetRecord(
		context.Background(),
		&pb.GetRecordRequest{},
	)

	if err2 != nil {
		return nil, fmt.Errorf("getRecord: %v", err2)
	}

	return res.Score, nil
}

func main() {}
