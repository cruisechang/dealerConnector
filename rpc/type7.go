package rpc

import (
	pb "github.com/cruisechang/liveServer/protobuf"

	"context"
	"errors"
	"time"
)

func (r *clientManager) RoundResultType7(clientName string, data *pb.RoundResultType7Data) (*pb.Empty, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.RoundResultType7(ctx, data)
	return res, error
}
func (r *clientManager) UpdateResultType7(clientName string, data *pb.UpdateResultType7Data) (*pb.Empty, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.UpdateResultType7(ctx, data)
	return res, error
}

func (r *clientManager) HistoryResultType7(clientName string, data *pb.HistoryResultType7Data) (*pb.HistoryResultType7Res, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.HistoryResultType7(ctx, data)
	return res, error
}

func (r *clientManager) Rethrow(clientName string, data *pb.RethrowData) (*pb.Empty, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.Rethrow(ctx, data)
	return res, error
}
