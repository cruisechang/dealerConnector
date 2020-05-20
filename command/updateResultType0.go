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

type updateResultType0Processor struct {
	BasicProcessor
}

func NewUpdateResultType0Processor(processor BasicProcessor) (*updateResultType0Processor, error) {
	p := &updateResultType0Processor{
		BasicProcessor: processor,
	}

	return p, nil

}

func (p *updateResultType0Processor) Run(obj *nex.CommandObject) (resErr error) {
	logger := p.GetLogger()
	conf := p.GetConfigurer()
	user := obj.User
	logPrefix := "2604 updateResultType0"

	defer func() {
		if r := recover(); r != nil {
			logger.LogFile(nxLog.LevelPanic, fmt.Sprintf("%s panic:%v", logPrefix, r))
			resErr = fmt.Errorf("%s panic %v", logPrefix, r)
		}
	}()

	//parsing command data
	deStr, err := base64.StdEncoding.DecodeString(obj.Cmd.Data)

	if err != nil {
		p.SendCommand(config.CodeBase64DecodeFailed, 0, conf.CmdUpdateResultType0(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s base64 decode cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	//取 data
	data := []config.UpdateResultType0CmdData{}

	if err := json.Unmarshal(deStr, &data); err != nil {
		p.SendCommand(config.CodeJsonUnmarshalJsonFailed, 0, conf.CmdUpdateResultType0(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json unmarshal cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	//rpc to server
	rpc := p.GetRPCManager()

	resData := &pb.UpdateResultType0Data{
		RoomID:      int32(user.RoomID()),
		Round:       data[0].Round,
		Result:      data[0].Result,
		BankerPair:  data[0].BankerPair,
		PlayerPair:  data[0].PlayerPair,
		BigSmall:    data[0].BigSmall,
		AnyPair:     data[0].AnyPair,
		PerfectPair: data[0].PerfectPair,
		SuperSix:    data[0].Super6,
		BankerPoint: data[0].BankerPoint,
		PlayerPoint: data[0].PlayerPoint,
		BankerPoker: data[0].BankerPoker,
		PlayerPoker: data[0].PlayerPoker,
	}

	_, err = rpc.UpdateResultType0("c0", resData)
	if err != nil {
		p.SendCommand(config.CodeRPCError, 0, conf.CmdUpdateResultType0(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s rpc error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))

		resErr = err
		return
	}

	p.SendCommand(config.CodeSuccess, 0, conf.CmdUpdateResultType0(), p.EmptySendDataStr(), user, []string{user.ConnID()})

	logger.LogFile(nxLog.LevelInfo, fmt.Sprintf("%s complete roomID=%d, resData=%+v", logPrefix, user.RoomID(), resData))
	return
}
