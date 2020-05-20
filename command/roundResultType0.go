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

type roundResultType0Processor struct {
	BasicProcessor
}

func NewRoundResultType0Processor(processor BasicProcessor) (*roundResultType0Processor, error) {
	p := &roundResultType0Processor{
		BasicProcessor: processor,
	}

	return p, nil

}

func (p *roundResultType0Processor) Run(obj *nex.CommandObject) (resErr error) {
	logger := p.GetLogger()
	conf := p.GetConfigurer()
	user := obj.User
	logPrefix := "2604 roundResultType0"

	defer func() {
		if r := recover(); r != nil {
			logger.LogFile(nxLog.LevelPanic, fmt.Sprintf("%s panic:%v", logPrefix, r))
			resErr = fmt.Errorf("%s panic %v", logPrefix, r)
		}
	}()

	//parsing command data
	deStr, err := base64.StdEncoding.DecodeString(obj.Cmd.Data)

	if err != nil {
		p.SendCommand(config.CodeBase64DecodeFailed, 0, conf.CmdRoundResultType0(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s base64 decode cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	//取 data
	data := []config.RoundResultType0CmdData{}

	if err := json.Unmarshal(deStr, &data); err != nil {
		p.SendCommand(config.CodeJsonUnmarshalJsonFailed, 0, conf.CmdRoundResultType0(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json unmarshal cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	//rpc to server
	rpc := p.GetRPCManager()

	resData := pb.RoundResultType0Data{
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
		PlayerPoker: data[0].PlayerPoker,
		BankerPoker: data[0].BankerPoker,
	}

	_, err = rpc.RoundResultType0("c0", &resData)
	if err != nil {
		p.SendCommand(config.CodeRPCError, 0, conf.CmdRoundResultType0(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s rpc error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))

		resErr = err
		return
	}

	p.SendCommand(config.CodeSuccess, 0, conf.CmdRoundResultType0(), p.EmptySendDataStr(), user, []string{user.ConnID()})

	logger.LogFile(nxLog.LevelInfo, fmt.Sprintf("%s complete roomID=%d, resData=%+v", logPrefix, user.RoomID(), resData))
	return
}
