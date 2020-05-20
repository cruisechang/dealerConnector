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

type historyResultType0Processor struct {
	BasicProcessor
}

func NewHistoryResultType0Processor(processor BasicProcessor) (*historyResultType0Processor, error) {
	p := &historyResultType0Processor{
		BasicProcessor: processor,
	}

	return p, nil

}

func (p *historyResultType0Processor) Run(obj *nex.CommandObject) (resErr error) {
	logger := p.GetLogger()
	conf := p.GetConfigurer()
	user := obj.User
	logPrefix := "2609 historyResultType0"

	defer func() {
		if r := recover(); r != nil {
			logger.LogFile(nxLog.LevelPanic, fmt.Sprintf("%s panic:%v", logPrefix, r))
			resErr = errors.New(fmt.Sprintf("%s panic %v", logPrefix, r))
		}
	}()

	//parsing command data
	deStr, err := base64.StdEncoding.DecodeString(obj.Cmd.Data)

	if err != nil {
		p.SendCommand(config.CodeBase64DecodeFailed, 0, conf.CmdHistoryResultType0(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s base64 decode cmd data error,user:%s,error:%s", logPrefix, obj.User.Name(), err.Error()))
		resErr = err
		return
	}

	//取 data
	data := []config.HistoryResultType0CmdData{}

	if err := json.Unmarshal(deStr, &data); err != nil {
		p.SendCommand(config.CodeJsonUnmarshalJsonFailed, 0, conf.CmdHistoryResultType0(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json unmarshal cmd data error,user:%s,error:%s", logPrefix, obj.User.Name(), err.Error()))
		resErr = err
		return
	}

	//data[0].Boot
	//data[0].Round

	//rpc
	rpc := p.GetRPCManager()

	rpcRes, err := rpc.HistoryResultType0("c0", &pb.HistoryResultType0Data{RoomID: int32(user.RoomID())})
	if err != nil {
		p.SendCommand(config.CodeRPCError, 0, conf.CmdHistoryResultType0(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s rpc error,roomID=%d, room:%s,error:%s", logPrefix, user.RoomID(), user.Name(), err.Error()))

		resErr = err
		return
	}

	//res data
	resInt := [][]int32{}

	for _, v := range rpcRes.Result {
		resInt = append(resInt, v.Result)
	}
	resData := []config.HistoryResultType0ResData{{
		Result: resInt,
	},
	}

	b, err := json.Marshal(resData)
	if err != nil {
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json marshal res data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	sendData := base64.StdEncoding.EncodeToString(b)
	p.SendCommand(config.CodeSuccess, 0, conf.CmdHistoryResultType0(), sendData, user, []string{user.ConnID()})

	logger.LogFile(nxLog.LevelInfo, fmt.Sprintf("%s complete roomID=%d, resData=%+v", logPrefix, user.RoomID(), resData))
	return
}
