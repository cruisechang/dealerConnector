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

type historyResultType2Processor struct {
	BasicProcessor
}

func NewHistoryResultType2Processor(processor BasicProcessor) (*historyResultType2Processor, error) {
	p := &historyResultType2Processor{
		BasicProcessor: processor,
	}

	return p, nil

}

func (p *historyResultType2Processor) Run(obj *nex.CommandObject) (resErr error) {
	logger := p.GetLogger()
	conf := p.GetConfigurer()
	user := obj.User
	logPrefix := "2109 historyResultType2"

	defer func() {
		if r := recover(); r != nil {
			logger.LogFile(nxLog.LevelPanic, fmt.Sprintf("%s panic:%v", logPrefix, r))
			resErr = errors.New(fmt.Sprintf("%s panic %v", logPrefix, r))
		}
	}()

	//parsing command data
	deStr, err := base64.StdEncoding.DecodeString(obj.Cmd.Data)

	if err != nil {
		p.SendCommand(config.CodeBase64DecodeFailed, 0, conf.CmdHistoryResultType2(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s base64 decode cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	//取 data
	data := []config.HistoryResultType2CmdData{}

	if err := json.Unmarshal(deStr, &data); err != nil {
		p.SendCommand(config.CodeJsonUnmarshalJsonFailed, 0, conf.CmdHistoryResultType2(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json unmarshal cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	//data[0].Boot
	//data[0].Round

	//rpc
	rpc := p.GetRPCManager()

	rpcRes, err := rpc.HistoryResultType2("c0", &pb.HistoryResultType2Data{RoomID: int32(user.RoomID())})
	if err != nil {
		p.SendCommand(config.CodeRPCError, 0, conf.CmdHistoryResultType2(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s rpc error,roomID=%d, room:%s,error:%s", logPrefix, user.RoomID(), user.Name(), err.Error()))

		resErr = err
		return
	}

	//res data
	resInt := [][]int32{}

	for _, v := range rpcRes.Result {
		resInt = append(resInt, v.Result)
	}
	resData := []config.HistoryResultType2ResData{
		{
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
	p.SendCommand(config.CodeSuccess, 0, conf.CmdHistoryResultType2(), sendData, user, []string{user.ConnID()})

	logger.LogFile(nxLog.LevelInfo, fmt.Sprintf("%s complete roomID=%d, resData=%+v", logPrefix, user.RoomID(), resData))
	return
}
