package rpc

import (
	"context"
	"errors"
	"time"

	pb "github.com/cruisechang/liveServer/protobuf"
)

//niuniu
func (r *clientManager) RoundProcess2(clientName string, data *pb.RoundProcess2Data) (*pb.Empty, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.RoundProcess2(ctx, data)
	return res, error
}

func (r *clientManager) RoundResultType2(clientName string, data *pb.RoundResultType2Data) (*pb.Empty, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.RoundResultType2(ctx, data)
	return res, error
	return nil, nil
}

func (r *clientManager) UpdateResultType2(clientName string, data *pb.UpdateResultType2Data) (*pb.Empty, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.UpdateResultType2(ctx, data)
	return res, error
}
func (r *clientManager) HistoryResultType2(clientName string, data *pb.HistoryResultType2Data) (*pb.HistoryResultType2Res, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.HistoryResultType2(ctx, data)
	return res, error
}
