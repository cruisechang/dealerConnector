package rpc

import (
	"context"
	"errors"
	"time"

	pb "github.com/cruisechang/liveServer/protobuf"
)

//dragonTiger
func (r *clientManager) RoundProcess1(clientName string, data *pb.RoundProcess1Data) (*pb.Empty, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.RoundProcess1(ctx, data)
	return res, error
}
func (r *clientManager) RoundResultType1(clientName string, data *pb.RoundResultType1Data) (*pb.Empty, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.RoundResultType1(ctx, data)
	return res, error
}
func (r *clientManager) UpdateResultType1(clientName string, data *pb.UpdateResultType1Data) (*pb.Empty, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.UpdateResultType1(ctx, data)
	return res, error
}
func (r *clientManager) HistoryResultType1(clientName string, data *pb.HistoryResultType1Data) (*pb.HistoryResultType1Res, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.HistoryResultType1(ctx, data)
	return res, error
}
