package main

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync"
	"time"

	pb "github.com/Darkhunter9/lazIR_tag/src/proto"

	"github.com/Jille/raft-grpc-leader-rpc/rafterrors"
	"github.com/hashicorp/raft"
)

type scoreTracker struct {
	// TODO: consider using sync.Map
	mtx   sync.RWMutex
	node  string
	score map[string]int32
}

var _ raft.FSM = &scoreTracker{}

func (f *scoreTracker) Apply(l *raft.Log) interface{} {
	f.mtx.Lock()
	defer f.mtx.Unlock()
	log := strings.Split(string(l.Data), " ")
	a, b := log[0], log[1]
	fmt.Printf("Node %s: %s get 1 from %s", f.node, a, b)
	f.score[a]++
	f.score[b]--

	return nil
}

func (f *scoreTracker) Snapshot() (raft.FSMSnapshot, error) {
	// Make sure that any future calls to f.Apply() don't change the snapshot.
	score := make(map[string]int32)
	for k, v := range f.score {
		score[k] = v
	}
	return &snapshot{score}, nil
}

func (f *scoreTracker) Restore(r io.ReadCloser) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	scores := strings.Split(string(b), " ")
	for i := range len(scores) / 2 {
		if s, err := strconv.Atoi(scores[2*i+1]); err == nil {
			f.score[scores[2*i]] = s
		}
	}
	return nil
}

type snapshot struct {
	score map[string]int32
}

func (s *snapshot) Persist(sink raft.SnapshotSink) error {
	var score string
	for k, v := range s.score {
		score += fmt.Sprintf("%s %d ", k, v)
	}

	_, err := sink.Write([]byte(strings.TrimSuffix(score, " ")))
	if err != nil {
		sink.Cancel()
		return fmt.Errorf("sink.Write(): %v", err)
	}
	return sink.Close()
}

func (s *snapshot) Release() {
}

type rpcInterface struct {
	scoreTracker *scoreTracker
	raft         *raft.Raft
}

func (r rpcInterface) AddRecord(ctx context.Context, req *pb.AddRecordRequest) (*pb.AddRecordResponse, error) {
	shooter := req.GetShooter()
	victim := req.GetVictim()
	f := r.raft.Apply([]byte(shooter+" "+victim), time.Second)
	if err := f.Error(); err != nil {
		return nil, rafterrors.MarkRetriable(err)
	}
	return &pb.AddRecordResponse{
		CommitIndex: f.Index(),
	}, nil
}

func (r rpcInterface) GetRecord(ctx context.Context, req *pb.GetRecordRequest) (*pb.GetRecordResponse, error) {
	r.scoreTracker.mtx.RLock()
	defer r.scoreTracker.mtx.RUnlock()

	score := make(map[string]int32)
	for k, v := range r.scoreTracker.score {
		score[k] = v
	}

	return &pb.GetRecordResponse{
		Score:       score,
		ReadAtIndex: r.raft.AppliedIndex(),
	}, nil
}
