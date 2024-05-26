package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	pb "github.com/Darkhunter9/lazIR_tag/src/proto"

	"github.com/Jille/raft-grpc-leader-rpc/leaderhealth"
	transport "github.com/Jille/raft-grpc-transport"
	"github.com/Jille/raftadmin"
	"github.com/hashicorp/raft"
	boltdb "github.com/hashicorp/raft-boltdb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	config = flag.String("config", "", "config file path")
)

func main() {
	flag.Parse()

	if *config == "" {
		log.Fatalf("flag --config is required")
	}
	c := getConf(*config)

	if c.Address == "localhost" {
		c.Address = GetOutboundIP().String()
		log.Printf("%s", c.Address)
	}
	c.Address = fmt.Sprintf("%s:%d", c.Address, c.Port)
	homeDir, _ := os.UserHomeDir()
	c.RaftDir = filepath.Join(homeDir, c.RaftDir, c.RaftId)
	err := os.MkdirAll(c.RaftDir, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to mkdir %s", c.RaftDir)
	}

	ctx := context.Background()
	_, port, err := net.SplitHostPort(c.Address)
	if err != nil {
		log.Fatalf("failed to parse local address (%s): %v", c.Address, err)
	}
	sock, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	st := NewScoreTracker(c)

	r, tm, err := NewRaft(ctx, c, st)
	if err != nil {
		log.Fatalf("failed to start raft: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterScoreServer(s, &rpcInterface{
		scoreTracker: st,
		raft:         r,
	})
	tm.Register(s)
	leaderhealth.Setup(r, s, []string{"Score"})
	raftadmin.Register(s, r)
	reflection.Register(s)
	if err := s.Serve(sock); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func NewRaft(ctx context.Context, config *conf, fsm raft.FSM) (*raft.Raft, *transport.Manager, error) {
	c := raft.DefaultConfig()
	c.LocalID = raft.ServerID(config.RaftId)

	ldb, err := boltdb.NewBoltStore(filepath.Join(config.RaftDir, "logs.dat"))
	if err != nil {
		return nil, nil, fmt.Errorf(`boltdb.NewBoltStore(%q): %v`, filepath.Join(config.RaftDir, "logs.dat"), err)
	}

	sdb, err := boltdb.NewBoltStore(filepath.Join(config.RaftDir, "stable.dat"))
	if err != nil {
		return nil, nil, fmt.Errorf(`boltdb.NewBoltStore(%q): %v`, filepath.Join(config.RaftDir, "stable.dat"), err)
	}

	fss, err := raft.NewFileSnapshotStore(config.RaftDir, 3, os.Stderr)
	if err != nil {
		return nil, nil, fmt.Errorf(`raft.NewFileSnapshotStore(%q, ...): %v`, config.RaftDir, err)
	}

	tm := transport.New(raft.ServerAddress(config.Address), []grpc.DialOption{grpc.WithInsecure()})

	r, err := raft.NewRaft(c, fsm, ldb, sdb, fss, tm.Transport())
	if err != nil {
		return nil, nil, fmt.Errorf("raft.NewRaft: %v", err)
	}

	if config.RaftBootstrap {
		cfg := raft.Configuration{
			Servers: []raft.Server{
				{
					Suffrage: raft.Voter,
					ID:       raft.ServerID(config.RaftId),
					Address:  raft.ServerAddress(config.Address),
				},
			},
		}
		f := r.BootstrapCluster(cfg)
		if err := f.Error(); err != nil {
			return nil, nil, fmt.Errorf("raft.Raft.BootstrapCluster: %v", err)
		}
	}

	return r, tm, nil
}

func NewScoreTracker(config *conf) *scoreTracker {
	s := make(map[string]int32)
	for _, player := range config.Players {
		s[player] = int32(config.Life)
	}

	return &scoreTracker{node: config.RaftId, score: s}
}
