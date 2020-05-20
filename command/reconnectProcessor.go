package command

import (
	"github.com/cruisechang/nex"
	nxLog "github.com/cruisechang/nex/log"
	"fmt"
	"encoding/json"
	"encoding/base64"
	"github.com/cruisechang/dealerConnector/config"
	"strconv"
	"errors"
)

//loginProcessor implements command.Processor
type reconnectProcessor struct {
	BasicProcessor
}

func NewReconnectProcessor(processor BasicProcessor) (*reconnectProcessor, error) {
	p := &reconnectProcessor{
		BasicProcessor: processor,
	}

	return p, nil
}

func (p *reconnectProcessor) Run(obj *nex.CommandObject) (resErr error) {

	conf := p.GetConfigurer()
	logger := p.GetLogger()
	user := obj.User
	logPrefix := "2016 reconnect"

	defer func() {
		if r := recover(); r != nil {
			logger.LogFile(nxLog.LevelPanic, fmt.Sprintf("%s panic:%v", logPrefix, r))
			resErr = errors.New(fmt.Sprintf("%s panic %v", logPrefix, r))
		}
	}()

	//parsing cmd data
	de, err := base64.StdEncoding.DecodeString(obj.Cmd.Data)

	if err != nil {
		p.DisconnectUser(user.UserID())
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s base64 decode cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	data := []config.ReconnectCmdData{}

	if err := json.Unmarshal(de, &data); err != nil {
		p.DisconnectUser(user.UserID())
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json Unmarshal cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	roomID := data[0].RoomID

	// find old room
	room, ok := p.GetRoomManager().GetRoom(roomID)
	if !ok {
		p.SendCommand(config.CodeRoomNotFound, 0, conf.CmdReconnect(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		return
	}

	room.AddUser(user)
	user.SetName(strconv.Itoa(roomID))
	room.AddUser(user)

	//res data
	resData := []config.ReconnectResData{{
			Success: 1,
		},
	}

	b, err := json.Marshal(resData)
	if err != nil {
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json marshal res data error,user:%s,error:%s ", logPrefix, user.Name(), err.Error()))
		p.SendCommand(config.CodeMarshalJsonFailed, 0, conf.CmdReconnect(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		resErr = err
		return
	}

	sendData := base64.StdEncoding.EncodeToString(b)

	p.SendCommand(config.CodeSuccess, 0, conf.CmdReconnect(), sendData, user, []string{user.ConnID()})

	logger.LogFile(nxLog.LevelInfo, fmt.Sprintf("%s complete  roomID=%d ", logPrefix, data[0].RoomID))

	return
}
