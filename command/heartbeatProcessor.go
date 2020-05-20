package command

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cruisechang/dealerConnector/config"
	"github.com/cruisechang/nex"
	nxLog "github.com/cruisechang/nex/log"
)

type heartbeatProcessor struct {
	BasicProcessor
}

func NewHeartbeatProcessor(processor BasicProcessor) (*heartbeatProcessor, error) {
	p := &heartbeatProcessor{
		BasicProcessor: processor,
	}

	return p, nil

}

func (p *heartbeatProcessor) Run(obj *nex.CommandObject) (resErr error) {
	logger := p.GetLogger()
	conf := p.GetConfigurer()
	user := obj.User
	logPrefix := "2012 heartBeat"

	defer func() {
		if r := recover(); r != nil {
			logger.LogFile(nxLog.LevelPanic, fmt.Sprintf("%s panic:%v", logPrefix, r))
			resErr = errors.New(fmt.Sprintf("%s panic %v", logPrefix, r))
		}
	}()

	//parsing command data
	deStr, err := base64.StdEncoding.DecodeString(obj.Cmd.Data)

	if err != nil {
		p.SendCommand(config.CodeBase64DecodeFailed, 0, conf.CmdHeartbeat(), p.EmptySendDataStr(), obj.User, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s base64 decode cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	data := []config.HeartbeatCmdData{}

	if err := json.Unmarshal(deStr, &data); err != nil {
		p.SendCommand(config.CodeJsonUnmarshalJsonFailed, 0, conf.CmdHeartbeat(), p.EmptySendDataStr(), obj.User, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json unmarshal cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	p.SendCommand(config.CodeSuccess, 0, conf.CmdHeartbeat(), p.EmptySendDataStr(), obj.User, []string{user.ConnID()})

	logger.LogFile(nxLog.LevelInfo, fmt.Sprintf("%s complete userID=%d, roomID=%d", logPrefix, user.UserID(), user.RoomID()))
	return
}
