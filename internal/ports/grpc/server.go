package grpc

import "github.com/amirhnajafiz/starling/pkg/raftpb"

// RaftServerImpl implements the RaftServer interface defined in raftpb.
type RaftServerImpl struct {
	raftpb.UnimplementedRaftServer
}

// AppendEntries handles incoming AppendEntries requests from other Raft nodes.
func (r *RaftServerImpl) AppendEntries(req *raftpb.AppendEntriesRequest, stream *raftpb.AppendEntriesResponse) error {
	return nil
}

// RequestVote handles incoming RequestVote requests from other Raft nodes.
func (r *RaftServerImpl) RequestVote(req *raftpb.RequestVoteRequest, stream *raftpb.RequestVoteResponse) error {
	return nil
}
