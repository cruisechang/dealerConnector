package command

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cruisechang/dealerConnector/config"
	pb "github.com/cruisechang/liveServer/protobuf"

	"github.com/cruisechang/nex"
	nxLog "github.com/cruisechang/nex/log"
)

type beginBettingProcessor struct {
	BasicProcessor
}

func NewBeginBettingProcessor(processor BasicProcessor) (*beginBettingProcessor, error) {
	p := &beginBettingProcessor{
		BasicProcessor: processor,
	}

	return p, nil

}

func (p *beginBettingProcessor) Run(obj *nex.CommandObject) (resErr error) {
	logger := p.GetLogger()
	conf := p.GetConfigurer()
	user := obj.User
	logPrefix := "2005 beginBetting"

	defer func() {
		if r := recover(); r != nil {
			logger.LogFile(nxLog.LevelPanic, fmt.Sprintf("%s panic:%v", logPrefix, r))
			resErr = errors.New(fmt.Sprintf("%s panic %v", logPrefix, r))
		}
	}()

	//parsing command data
	deStr, err := base64.StdEncoding.DecodeString(obj.Cmd.Data)

	if err != nil {
		p.SendCommand(config.CodeBase64DecodeFailed, 0, conf.CmdBeginBetting(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s base64 decode cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	//取 data
	data := []config.BeginBettingCmdData{}

	if err := json.Unmarshal(deStr, &data); err != nil {
		p.SendCommand(config.CodeJsonUnmarshalJsonFailed, 0, conf.CmdBeginBetting(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json unmarshal cmd data error,user:%s,error:%s", logPrefix, obj.User.Name(), err.Error()))
		resErr = err
		return
	}

	//data[0].Boot
	//data[0].Inning

	//rpc server
	rpc := p.GetRPCManager()
	rpcRes, err := rpc.BeginBetting("c0", &pb.BeginBettingData{RoomID: int64(user.RoomID()), Boot: int64(data[0].Boot), Round: int64(data[0].Round)})
	if err != nil {
		p.SendCommand(config.CodeRPCError, 0, conf.CmdWaiting(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s rpc error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))

		resErr = err
		return
	}

	resData := []config.BeginBettingResData{
		{
			Round: rpcRes.Round,
		},
	}
	b, err := json.Marshal(resData)
	if err != nil {
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json marshal res data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	sendData := base64.StdEncoding.EncodeToString(b)
	p.SendCommand(config.CodeSuccess, 0, conf.CmdBeginBetting(), sendData, user, []string{user.ConnID()})

	logger.LogFile(nxLog.LevelInfo, fmt.Sprintf("%s complete roomID=%d, resData=%+v", logPrefix, user.RoomID(), resData))

	return
}
