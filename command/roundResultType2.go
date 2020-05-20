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

type roundResultType2Processor struct {
	BasicProcessor
}

func NewRoundResultType2Processor(processor BasicProcessor) (*roundResultType2Processor, error) {
	p := &roundResultType2Processor{
		BasicProcessor: processor,
	}

	return p, nil

}

func (p *roundResultType2Processor) Run(obj *nex.CommandObject) (resErr error) {
	logger := p.GetLogger()
	conf := p.GetConfigurer()
	user := obj.User
	logPrefix := "2104 roundResultType2"

	defer func() {
		if r := recover(); r != nil {
			logger.LogFile(nxLog.LevelPanic, fmt.Sprintf("%s panic:%v", logPrefix, r))
			resErr = fmt.Errorf("%s panic %v", logPrefix, r)
		}
	}()

	//parsing command data
	deStr, err := base64.StdEncoding.DecodeString(obj.Cmd.Data)

	if err != nil {
		p.SendCommand(config.CodeBase64DecodeFailed, 0, conf.CmdRoundResultType2(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s base64 decode cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	//取 data
	data := []config.RoundResultType2CmdData{}

	if err := json.Unmarshal(deStr, &data); err != nil {
		p.SendCommand(config.CodeJsonUnmarshalJsonFailed, 0, conf.CmdRoundResultType2(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json unmarshal cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	//rpc to server
	rpc := p.GetRPCManager()

	resData := pb.RoundResultType2Data{
		RoomID: int32(user.RoomID()),
		Round:  data[0].Round,
		Head:   data[0].Head,
		Owner0: &pb.RoundResultType2Owner{Result: data[0].Owner0.Result, Poker: data[0].Owner0.Poker, Pattern: data[0].Owner0.Pattern},
		Owner1: &pb.RoundResultType2Owner{Result: data[0].Owner1.Result, Poker: data[0].Owner1.Poker, Pattern: data[0].Owner1.Pattern},
		Owner2: &pb.RoundResultType2Owner{Result: data[0].Owner2.Result, Poker: data[0].Owner2.Poker, Pattern: data[0].Owner2.Pattern},
		Owner3: &pb.RoundResultType2Owner{Result: data[0].Owner3.Result, Poker: data[0].Owner3.Poker, Pattern: data[0].Owner3.Pattern},
	}

	_, err = rpc.RoundResultType2("c0", &resData)
	if err != nil {
		p.SendCommand(config.CodeRPCError, 0, conf.CmdRoundResultType2(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s rpc error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))

		resErr = err
		return
	}

	p.SendCommand(config.CodeSuccess, 0, conf.CmdRoundResultType2(), p.EmptySendDataStr(), user, []string{user.ConnID()})

	logger.LogFile(nxLog.LevelInfo, fmt.Sprintf("%s complete roomID=%d, resData=%+v", logPrefix, user.RoomID(), resData))
	return
}
