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

type roundResultType6Processor struct {
	BasicProcessor
}

func NewRoundResultType6Processor(processor BasicProcessor) (*roundResultType6Processor, error) {
	p := &roundResultType6Processor{
		BasicProcessor: processor,
	}

	return p, nil

}

func (p *roundResultType6Processor) Run(obj *nex.CommandObject) (resErr error) {
	logger := p.GetLogger()
	conf := p.GetConfigurer()
	user := obj.User
	logPrefix := "2504 roundResultType6"

	defer func() {
		if r := recover(); r != nil {
			logger.LogFile(nxLog.LevelPanic, fmt.Sprintf("%s panic:%v", logPrefix, r))
			resErr = fmt.Errorf("%s panic %v", logPrefix, r)
		}
	}()

	//parsing command data
	deStr, err := base64.StdEncoding.DecodeString(obj.Cmd.Data)

	if err != nil {
		p.SendCommand(config.CodeBase64DecodeFailed, 0, conf.CmdRoundResultType6(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s base64 decode cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	//取 data
	data := []config.RoundResultType6CmdData{}

	if err := json.Unmarshal(deStr, &data); err != nil {
		p.SendCommand(config.CodeJsonUnmarshalJsonFailed, 0, conf.CmdRoundResultType6(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		sprintf := fmt.Sprintf("%s json unmarshal cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error())
		logger.LogFile(nxLog.LevelError, sprintf)
		resErr = err
		return
	}

	//rpc to server
	rpc := p.GetRPCManager()

	//paigou [][]32
	var pbAry []*pb.RoundResultType6DataInnerType
	for _, v := range data[0].Paigow {
		pbAry = append(pbAry, &pb.RoundResultType6DataInnerType{Result: v})

	}

	rpcProto := pb.RoundResultType6Data{
		RoomID:   int32(user.RoomID()),
		Round:    data[0].Round,
		Dice:     data[0].Dice,
		Sum:      data[0].Sum,
		BigSmall: data[0].BigSmall,
		OddEven:  data[0].OddEven,
		Triple:   data[0].Triple,
		Pair:     data[0].Pair,
		Paigow:   pbAry,
	}

	_, err = rpc.RoundResultType6("c0", &rpcProto)
	if err != nil {
		p.SendCommand(config.CodeRPCError, 0, conf.CmdRoundResultType6(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s rpc error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))

		resErr = err
		return
	}

	//res
	resData := p.EmptySendDataStr()
	p.SendCommand(config.CodeSuccess, 0, conf.CmdRoundResultType6(), resData, user, []string{user.ConnID()})

	logger.LogFile(nxLog.LevelInfo, fmt.Sprintf("%s complete roomID=%d, rpcProto=%+v, resData=%+v", logPrefix, user.RoomID(), rpcProto, resData))
	return
}
