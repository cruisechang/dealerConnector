package command

import (
	"github.com/cruisechang/nex"
	nxLog "github.com/cruisechang/nex/log"
	"github.com/cruisechang/dealerConnector/config"
	pb "github.com/cruisechang/liveServer/protobuf"
	"fmt"
	"encoding/base64"
	"encoding/json"
	"errors"
)

type rerollDiceProcessor struct {
	BasicProcessor
}

func NewRerollDiceProcessor(processor BasicProcessor) (*rerollDiceProcessor, error) {
	p := &rerollDiceProcessor{
		BasicProcessor: processor,
	}

	return p, nil

}

func (p *rerollDiceProcessor) Run(obj *nex.CommandObject) (resErr error) {
	logger := p.GetLogger()
	conf := p.GetConfigurer()
	user := obj.User
	logPrefix:="2510 rerollDice"

	defer func() {
		if r := recover(); r != nil {
			logger.LogFile(nxLog.LevelPanic, fmt.Sprintf("%s panic:%v ", logPrefix,r))
			resErr = errors.New(fmt.Sprintf("%s panic %v", logPrefix,r))
		}
	}()

	//parsing command data
	deStr, err := base64.StdEncoding.DecodeString(obj.Cmd.Data)

	if err != nil {
		p.SendCommand(config.CodeBase64DecodeFailed, 0, conf.CmdRerollDice(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s base64 decode cmd data error,user:%s,error:%s", logPrefix,user.Name(), err.Error()))
		resErr = err
		return
	}

	//取 data
	var data []config.RerollDiceCmdData

	if err := json.Unmarshal(deStr, &data); err != nil {
		p.SendCommand(config.CodeJsonUnmarshalJsonFailed, 0, conf.CmdRerollDice(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json unmarshal cmd data error,user:%s,error:%s", logPrefix,user.Name(), err.Error()))
		resErr = err
		return
	}

	//data[0].Boot
	//data[0].Round

	//rpc
	rpc := p.GetRPCManager()
	rpcProto:=&pb.RerollDiceData{RoomID: int32(user.RoomID()),Round:data[0].Round}

	_, err = rpc.RerollDice("c0", rpcProto)
	if err != nil {
		p.SendCommand(config.CodeRPCError, 0, conf.CmdRerollDice(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s rpc error,roomID=%d, room:%s,error:%s",logPrefix, user.RoomID(), user.Name(), err.Error()))

		resErr = err
		return
	}

	//res
	resData := p.EmptySendDataStr()
	p.SendCommand(config.CodeSuccess, 0, conf.CmdRerollDice(),  resData, user, []string{user.ConnID()})

	logger.LogFile(nxLog.LevelInfo, fmt.Sprintf("%s complete roomID=%d, rpcProto=%+v, resData=%+v ", logPrefix,user.RoomID(),rpcProto,resData))
	return
}
