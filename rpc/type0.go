package rpc

import (
	"context"
	"errors"
	"time"

	pb "github.com/cruisechang/liveServer/protobuf"
)

//baccarat
//dragonTiger
func (r *clientManager) RoundProcess0(clientName string, data *pb.RoundProcess0Data) (*pb.Empty, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.RoundProcess0(ctx, data)
	return res, error
}
func (r *clientManager) RoundResultType0(clientName string, data *pb.RoundResultType0Data) (*pb.Empty, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.RoundResultType0(ctx, data)
	return res, error
}
func (r *clientManager) UpdateResultType0(clientName string, data *pb.UpdateResultType0Data) (*pb.Empty, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.UpdateResultType0(ctx, data)
	return res, error
}

func (r *clientManager) HistoryResultType0(clientName string, data *pb.HistoryResultType0Data) (*pb.HistoryResultType0Res, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.HistoryResultType0(ctx, data)
	return res, error
}
