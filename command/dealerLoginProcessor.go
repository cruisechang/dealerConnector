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

type dealerLoginProcessor struct {
	BasicProcessor
}

func NewDealerLoginProcessor(processor BasicProcessor) (*dealerLoginProcessor, error) {
	p := &dealerLoginProcessor{
		BasicProcessor: processor,
	}

	return p, nil
}

func (p *dealerLoginProcessor) Run(obj *nex.CommandObject) (resErr error) {
	logger := p.GetLogger()
	conf := p.GetConfigurer()
	rm := p.GetRoomManager()
	user := obj.User
	logPrefix := "2015 dealerLogin"

	defer func() {
		if r := recover(); r != nil {
			logger.LogFile(nxLog.LevelPanic, fmt.Sprintf("%s panic:%v", logPrefix, r))
			resErr = errors.New(fmt.Sprintf("%s panic %v", logPrefix, r))
		}
	}()

	//parsing command data
	deStr, err := base64.StdEncoding.DecodeString(obj.Cmd.Data)

	if err != nil {
		p.SendCommand(config.CodeBase64DecodeFailed, 0, conf.CmdDealerLogin(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s base64 decode cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	//取 data
	data := []config.DealerLoginCmdData{}

	if err := json.Unmarshal(deStr, &data); err != nil {
		p.SendCommand(config.CodeJsonUnmarshalJsonFailed, 0, conf.CmdDealerLogin(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json unmarshal cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}
	receiveData := data[0]

	//not yet create room, create room is in roomLogin
	_, ok := rm.GetRoom(receiveData.RoomID)
	if !ok {
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s DEALER LOGIN START ===================", logPrefix))
	}

	logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s receive data=%+v", logPrefix, receiveData))

	//rpc server
	rpc := p.GetRPCManager()

	rpcRes, err := rpc.DealerLogin("c0", &pb.DealerLoginData{RoomID: int64(receiveData.RoomID), Dealer: int64(receiveData.Dealer), Password: receiveData.Password})
	if err != nil {
		p.SendCommand(config.CodeRPCError, 0, conf.CmdDealerLogin(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s rpc error,roomID=%d,error:=%s", logPrefix, receiveData.RoomID, err.Error()))
		resErr = err
		return
	}

	//res data
	resData := []config.DealerLoginResData{{
		Success: int(rpcRes.Success),
	},
	}
	b, err := json.Marshal(resData)
	if err != nil {
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json marshal res data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}
	sendData := base64.StdEncoding.EncodeToString(b)

	p.SendCommand(config.CodeSuccess, 0, conf.CmdDealerLogin(), sendData, user, []string{user.ConnID()})

	logger.LogFile(nxLog.LevelInfo, fmt.Sprintf("%s complete roomID=%d, resData=%+v", logPrefix, receiveData.RoomID, resData))
	return
}
