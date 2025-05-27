package service

import (
	"context"

	"github.com/amirhnajafiz/starling/pkg/raftpb"
)

// RaftServerImpl implements the RaftServer interface defined in raftpb.
type RaftServerImpl struct {
	raftpb.UnimplementedRaftServer
}

// AppendEntries handles incoming AppendEntries requests from other Raft nodes.
func (r *RaftServerImpl) AppendEntries(ctx context.Context, req *raftpb.AppendEntriesRequest) (*raftpb.AppendEntriesResponse, error) {
	return nil, nil
}

// RequestVote handles incoming RequestVote requests from other Raft nodes.
func (r *RaftServerImpl) RequestVote(ctx context.Context, req *raftpb.RequestVoteRequest) (*raftpb.RequestVoteResponse, error) {
	return nil, nil
}
