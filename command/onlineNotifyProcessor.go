package command

import (
	"github.com/cruisechang/nex"
	nxLog "github.com/cruisechang/nex/log"
	"fmt"
	"encoding/json"
	"encoding/base64"
	"github.com/cruisechang/dealerConnector/config"
	pb "github.com/cruisechang/liveServer/protobuf"
	"errors"
)

//loginProcessor implements command.Processor
type onlineNotifyProcessor struct {
	BasicProcessor
}

func NewOnlineNotifyProcessor(processor BasicProcessor) (*onlineNotifyProcessor,error) {
	c := &onlineNotifyProcessor{
		BasicProcessor: processor,

	}
	return c,nil
}

//return error means some thing strange happened
//回傳error是表示程式發生不預期錯誤，
//如果是login 失敗，不回傳錯誤
//failed must logout
//failed must  disconnect client
func (p *onlineNotifyProcessor) Run(obj *nex.CommandObject) (resErr error) {

	logger := p.GetLogger()
	conf := p.GetConfigurer()
	user:=obj.User
	logPrefix:="2013 onlineNotify"

	defer func() {
		if r := recover(); r != nil {
			logger.LogFile(nxLog.LevelPanic, fmt.Sprintf("%s panic:%v", logPrefix,r))
			resErr = errors.New(fmt.Sprintf("%s panic %v",  logPrefix,r))
		}
	}()

	//parsing cmd data
	de, err := base64.StdEncoding.DecodeString(obj.Cmd.Data)

	if err != nil {
		p.DisconnectUser(user.UserID())
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s base64 decode cmd data error,user:%s,error:%s",  logPrefix,user.Name(), err.Error()))
		resErr=err
		return
	}

	data := []config.OnlineNotifyCmdData{}

	if err := json.Unmarshal(de, &data); err != nil {
		p.DisconnectUser(user.UserID())
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json Unmarshal cmd data error,user:%s,error:%s",  logPrefix,user.Name(), err.Error()))
		resErr=err
		return
	}

	//rpc game server
	rpc := p.GetRPCManager()

	_, err = rpc.OnlineNotify("c0", &pb.OnlineNotifyData{RoomID: int64(user.UserID())})
	if err != nil {
		p.DisconnectUser(user.UserID())
		//p.SendCommand(config.CodeRPCError, 0, conf.CmdGetRoomInfo(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s rpc error,user:%s,error:%s",  logPrefix,user.Name(), err.Error()))

		resErr = err
		return
	}


	p.SendCommand(config.CodeSuccess, 0, conf.CmdOnlineNotify(), p.EmptySendDataStr(), user, []string{user.ConnID()})


	logger.LogFile(nxLog.LevelInfo, fmt.Sprintf("%s complete roomID=%d",  logPrefix,user.RoomID()))


	return
}

