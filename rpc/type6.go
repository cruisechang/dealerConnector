package rpc

import (
	"context"
	"errors"
	"time"

	pb "github.com/cruisechang/liveServer/protobuf"
)

func (r *clientManager) RoundResultType6(clientName string, data *pb.RoundResultType6Data) (*pb.Empty, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.RoundResultType6(ctx, data)
	return res, error
}
func (r *clientManager) UpdateResultType6(clientName string, data *pb.UpdateResultType6Data) (*pb.Empty, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.UpdateResultType6(ctx, data)
	return res, error
}

func (r *clientManager) HistoryResultType6(clientName string, data *pb.HistoryResultType6Data) (*pb.HistoryResultType6Res, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.HistoryResultType6(ctx, data)
	return res, error
}

func (r *clientManager) RerollDice(clientName string, data *pb.RerollDiceData) (*pb.Empty, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.RerollDice(ctx, data)
	return res, error
}
