package rpc

import (
	"context"
	"errors"
	"log"
	"time"

	pb "github.com/cruisechang/liveServer/protobuf"
	"github.com/cruisechang/nex"
)

type ClientManager interface {
	AddClient(addr, port string, name string) (pb.RPCClient, error)
	GetClientByName(key string) (pb.RPCClient, bool)

	DealerLogin(clientName string, data *pb.DealerLoginData) (*pb.DealerLoginRes, error)
	RoomLlogin(clientName string, data *pb.RoomLoginData) (*pb.Empty, error)
	OnlineNotify(clientName string, data *pb.OnlineNotifyData) (*pb.Empty, error)
	GetRoomInfo(clientName string, data *pb.GetRoomInfoData) (*pb.GetRoomInfoRes, error)

	Waiting(clientName string, data *pb.WaitingData) (*pb.Empty, error)
	BeginBetting(clientName string, data *pb.BeginBettingData) (*pb.BeginBettingRes, error)
	EndBetting(clientName string, data *pb.EndBettingData) (*pb.Empty, error)
	ChangeBoot(clientName string, data *pb.ChangeBootData) (*pb.ChangeBootRes, error)
	CancelRound(clientName string, data *pb.CancelRoundData) (*pb.Empty, error)

	//baccarat
	RoundProcess0(clientName string, data *pb.RoundProcess0Data) (*pb.Empty, error)
	RoundResultType0(clientName string, data *pb.RoundResultType0Data) (*pb.Empty, error)
	UpdateResultType0(clientName string, data *pb.UpdateResultType0Data) (*pb.Empty, error)
	HistoryResultType0(clientName string, data *pb.HistoryResultType0Data) (*pb.HistoryResultType0Res, error)

	//dragonTiger
	RoundProcess1(clientName string, data *pb.RoundProcess1Data) (*pb.Empty, error)
	RoundResultType1(clientName string, data *pb.RoundResultType1Data) (*pb.Empty, error)
	UpdateResultType1(clientName string, data *pb.UpdateResultType1Data) (*pb.Empty, error)
	HistoryResultType1(clientName string, data *pb.HistoryResultType1Data) (*pb.HistoryResultType1Res, error)

	//sicbo
	RoundResultType6(clientName string, data *pb.RoundResultType6Data) (*pb.Empty, error)
	UpdateResultType6(clientName string, data *pb.UpdateResultType6Data) (*pb.Empty, error)
	HistoryResultType6(clientName string, data *pb.HistoryResultType6Data) (*pb.HistoryResultType6Res, error)
	RerollDice(clientName string, data *pb.RerollDiceData) (*pb.Empty, error)

	//rolette
	RoundResultType7(clientName string, data *pb.RoundResultType7Data) (*pb.Empty, error)
	UpdateResultType7(clientName string, data *pb.UpdateResultType7Data) (*pb.Empty, error)
	HistoryResultType7(clientName string, data *pb.HistoryResultType7Data) (*pb.HistoryResultType7Res, error)
	Rethrow(clientName string, data *pb.RethrowData) (*pb.Empty, error)

	//niuniu
	RoundProcess2(clientName string, data *pb.RoundProcess2Data) (*pb.Empty, error)
	RoundResultType2(clientName string, data *pb.RoundResultType2Data) (*pb.Empty, error)
	UpdateResultType2(clientName string, data *pb.UpdateResultType2Data) (*pb.Empty, error)
	HistoryResultType2(clientName string, data *pb.HistoryResultType2Data) (*pb.HistoryResultType2Res, error)
}
type clientManager struct {
	nex         nex.Nex
	clientTable map[string]pb.RPCClient
}

func NewRPCManager(nex nex.Nex) ClientManager {
	return &clientManager{
		nex:         nex,
		clientTable: make(map[string]pb.RPCClient),
	}
}

func (r *clientManager) AddClient(addr, port string, name string) (pb.RPCClient, error) {

	c, err := r.nex.GetGRPCClient(addr, port, pb.NewRPCClient)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	client, ok := c.(pb.RPCClient)
	if !ok {
		log.Println("new client error")
		return nil, errors.New("nex grpcClient assertion error")
	}

	r.clientTable[name] = client

	return client, nil
}
func (r *clientManager) GetClientByName(name string) (pb.RPCClient, bool) {
	v, ok := r.clientTable[name]
	return v, ok
}
func (r *clientManager) DealerLogin(clientName string, data *pb.DealerLoginData) (*pb.DealerLoginRes, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.DealerLogin(ctx, data)
	return res, error
}
func (r *clientManager) RoomLlogin(clientName string, data *pb.RoomLoginData) (*pb.Empty, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.RoomLogin(ctx, data)
	return res, error
}
func (r *clientManager) OnlineNotify(clientName string, data *pb.OnlineNotifyData) (*pb.Empty, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.OnlineNotify(ctx, data)
	return res, error
}

func (r *clientManager) GetRoomInfo(clientName string, data *pb.GetRoomInfoData) (*pb.GetRoomInfoRes, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.GetRoomInfo(ctx, data)
	return res, error
}

func (r *clientManager) Waiting(clientName string, data *pb.WaitingData) (*pb.Empty, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.Waiting(ctx, data)
	return res, error
}

func (r *clientManager) BeginBetting(clientName string, data *pb.BeginBettingData) (*pb.BeginBettingRes, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.BeginBetting(ctx, data)
	return res, error
}
func (r *clientManager) EndBetting(clientName string, data *pb.EndBettingData) (*pb.Empty, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.EndBetting(ctx, data)
	return res, error
}

func (r *clientManager) ChangeBoot(clientName string, data *pb.ChangeBootData) (*pb.ChangeBootRes, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.ChangeBoot(ctx, data)
	return res, error
}
func (r *clientManager) CancelRound(clientName string, data *pb.CancelRoundData) (*pb.Empty, error) {
	c, ok := r.GetClientByName(clientName)
	if !ok {
		return nil, errors.New("rpc client not found by passed key")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, error := c.CancelRound(ctx, data)
	return res, error
}
