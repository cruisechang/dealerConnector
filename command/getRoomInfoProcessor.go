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

type getRoomInfoProcessor struct {
	BasicProcessor
}

func NewGetRoomInfoProcessor(processor BasicProcessor) (*getRoomInfoProcessor, error) {
	p := &getRoomInfoProcessor{
		BasicProcessor: processor,
	}

	return p, nil

}

func (p *getRoomInfoProcessor) Run(obj *nex.CommandObject) (resErr error) {
	logger := p.GetLogger()
	conf := p.GetConfigurer()
	user := obj.User
	logPrefix := "2001 getRoomInfo"

	defer func() {
		if r := recover(); r != nil {
			logger.LogFile(nxLog.LevelPanic, fmt.Sprintf("%s panic:%v ", logPrefix, r))
			resErr = errors.New(fmt.Sprintf("%s panic %v", logPrefix, r))
		}
	}()

	//parsing command data
	deStr, err := base64.StdEncoding.DecodeString(obj.Cmd.Data)

	if err != nil {
		p.SendCommand(config.CodeBase64DecodeFailed, 0, conf.CmdGetRoomInfo(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s base64 decode cmd data error,user:%s,error:%s ", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	//取 data
	data := []config.GetRoomInfoCmdData{}

	if err := json.Unmarshal(deStr, &data); err != nil {
		p.SendCommand(config.CodeJsonUnmarshalJsonFailed, 0, conf.CmdGetRoomInfo(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json unmarshal cmd data error,user:%s,error:%s ", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	//rpc game server
	rpc := p.GetRPCManager()

	rpcRes, err := rpc.GetRoomInfo("c0", &pb.GetRoomInfoData{RoomID: int64(user.RoomID())})
	if err != nil {
		p.SendCommand(config.CodeRPCError, 0, conf.CmdGetRoomInfo(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s rpc error,user:%s, error:%s ", logPrefix, user.Name(), err.Error()))

		resErr = err
		return
	}

	//res data
	resData := []config.GetRoomInfoResData{
		{
			Boot:                int(rpcRes.Boot),
			Round:               int(rpcRes.Round),
			RoomID:              user.RoomID(),
			RoomName:            rpcRes.RoomName,
			BankerPlayerMax:     int(rpcRes.BankerPlayerMax),
			BankerPlayerMin:     int(rpcRes.BankerPlayerMin),
			TieMax:              int(rpcRes.BankerPlayerMax),
			TieMin:              int(rpcRes.BankerPlayerMin),
			BankerPlayerPairMax: int(rpcRes.BankerPlayerMax),
			BankerPlayerPairMin: int(rpcRes.BankerPlayerMin),
			Online:              int(rpcRes.Online),
			BetCountDown:        int(rpcRes.BetCountDown) + 2,
		},
	}

	b, err := json.Marshal(resData)
	if err != nil {
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json marshal res data error,user:%s,error:%s ", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	sendData := base64.StdEncoding.EncodeToString(b)
	p.SendCommand(config.CodeSuccess, 0, conf.CmdGetRoomInfo(), sendData, user, []string{user.ConnID()})

	logger.LogFile(nxLog.LevelInfo, fmt.Sprintf("%s complete roomID=%d, resData=%+v", logPrefix, user.RoomID(), resData))
	return
}
