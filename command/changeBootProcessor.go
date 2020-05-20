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

type changeBootProcessor struct {
	BasicProcessor
}

func NewChangeBootProcessor(processor BasicProcessor) (*changeBootProcessor, error) {
	p := &changeBootProcessor{
		BasicProcessor: processor,
	}

	return p, nil

}

func (p *changeBootProcessor) Run(obj *nex.CommandObject) (resErr error) {
	logger := p.GetLogger()
	conf := p.GetConfigurer()
	user := obj.User
	logPrefix := "2007 changeBoot"

	defer func() {
		if r := recover(); r != nil {
			logger.LogFile(nxLog.LevelPanic, fmt.Sprintf("%s panic:%v", logPrefix, r))
			resErr = errors.New(fmt.Sprintf("%s panic %v", logPrefix, r))
		}
	}()

	//parsing command data
	deStr, err := base64.StdEncoding.DecodeString(obj.Cmd.Data)
	if err != nil {
		p.SendCommand(config.CodeBase64DecodeFailed, 0, conf.CmdChangeBoot(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s base64 decode cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	//取 data
	data := []config.ChangeBootCmdData{}
	if err := json.Unmarshal(deStr, &data); err != nil {
		p.SendCommand(config.CodeJsonUnmarshalJsonFailed, 0, conf.CmdChangeBoot(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json unmarshal cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	//rpc server
	rpc := p.GetRPCManager()

	rpcRes, err := rpc.ChangeBoot("c0", &pb.ChangeBootData{RoomID: int64(user.RoomID())})
	if err != nil {
		p.SendCommand(config.CodeRPCError, 0, conf.CmdWaiting(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s rpc error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))

		resErr = err
		return
	}

	//res data
	resData := []config.ChangeBootResData{
		config.ChangeBootResData{
			Boot: int(rpcRes.Boot),
		},
	}

	b, err := json.Marshal(resData)
	if err != nil {
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json marshal res data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	sendData := base64.StdEncoding.EncodeToString(b)
	p.SendCommand(config.CodeSuccess, 0, conf.CmdChangeBoot(), sendData, user, []string{user.ConnID()})

	logger.LogFile(nxLog.LevelInfo, fmt.Sprintf("%s complete roomID=%d, resData=%+v", logPrefix, user.RoomID(), resData))

	return
}
