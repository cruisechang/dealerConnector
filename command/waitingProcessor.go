package command

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/cruisechang/dealerConnector/config"
	pb "github.com/cruisechang/liveServer/protobuf"
	"github.com/cruisechang/nex"
	nxLog "github.com/cruisechang/nex/log"
)

type waitingProcessor struct {
	BasicProcessor
}

func NewWaitingProcessor(processor BasicProcessor) (*waitingProcessor, error) {
	p := &waitingProcessor{
		BasicProcessor: processor,
	}

	return p, nil

}

func (p *waitingProcessor) Run(obj *nex.CommandObject) (resErr error) {
	logger := p.GetLogger()
	conf := p.GetConfigurer()
	user := obj.User
	logPrefix := "2002 waiting"

	defer func() {
		if r := recover(); r != nil {
			logger.LogFile(nxLog.LevelPanic, fmt.Sprintf("%s panic:%v", logPrefix, r))
			resErr = fmt.Errorf("%s panic %v", logPrefix, r)
		}
	}()

	//parsing command data
	deStr, err := base64.StdEncoding.DecodeString(obj.Cmd.Data)

	if err != nil {
		p.SendCommand(config.CodeBase64DecodeFailed, 0, conf.CmdWaiting(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s base64 decode cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	//取 data
	data := []config.WaitingCmdData{}

	if err := json.Unmarshal(deStr, &data); err != nil {
		p.SendCommand(config.CodeJsonUnmarshalJsonFailed, 0, conf.CmdWaiting(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json unmarshal cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	//rpc
	rpc := p.GetRPCManager()

	resData := &pb.WaitingData{
		RoomID: int64(user.RoomID()),
		Boot:   int64(data[0].Boot),
		Round:  int64(data[0].Round),
	}

	_, err = rpc.Waiting("c0", resData)
	if err != nil {
		p.SendCommand(config.CodeRPCError, 0, conf.CmdWaiting(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s rpc error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))

		resErr = err
		return
	}

	p.SendCommand(config.CodeSuccess, 0, conf.CmdWaiting(), p.EmptySendDataStr(), user, []string{user.ConnID()})

	logger.LogFile(nxLog.LevelInfo, fmt.Sprintf("%s complete roomID:%d, resData=%+v", logPrefix, user.RoomID(), resData))

	return resErr
}
