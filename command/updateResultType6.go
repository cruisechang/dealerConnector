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

type updateResultType6Processor struct {
	BasicProcessor
}

func NewUpdateResultType6Processor(processor BasicProcessor) (*updateResultType6Processor, error) {
	p := &updateResultType6Processor{
		BasicProcessor: processor,
	}

	return p, nil

}

func (p *updateResultType6Processor) Run(obj *nex.CommandObject) (resErr error) {
	logger := p.GetLogger()
	conf := p.GetConfigurer()
	user := obj.User
	logPrefix := "2508 updateResultType6"

	defer func() {
		if r := recover(); r != nil {
			logger.LogFile(nxLog.LevelPanic, fmt.Sprintf("%s panic:%v", logPrefix, r))
			resErr = fmt.Errorf("%s panic %v", logPrefix, r)
		}
	}()

	//parsing command data
	deStr, err := base64.StdEncoding.DecodeString(obj.Cmd.Data)

	if err != nil {
		p.SendCommand(config.CodeBase64DecodeFailed, 0, conf.CmdUpdateResultType6(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s base64 decode cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	//取 data
	var data []config.UpdateResultType6CmdData

	if err := json.Unmarshal(deStr, &data); err != nil {
		p.SendCommand(config.CodeJsonUnmarshalJsonFailed, 0, conf.CmdUpdateResultType6(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json unmarshal cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
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

	rpcProto := &pb.UpdateResultType6Data{
		RoomID:   int32(user.RoomID()),
		Round:    data[0].Round,
		Dice:     data[0].Dice,
		Sum:      int32(data[0].Sum),
		BigSmall: int32(data[0].BigSmall),
		OddEven:  int32(data[0].OddEven),
		Triple:   int32(data[0].Triple),
		Pair:     int32(data[0].Pair),
		Paigow:   pbAry,
	}

	_, err = rpc.UpdateResultType6("c0", rpcProto)
	if err != nil {
		p.SendCommand(config.CodeRPCError, 0, conf.CmdUpdateResultType6(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s rpc error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))

		resErr = err
		return
	}

	resData := p.EmptySendDataStr()
	p.SendCommand(config.CodeSuccess, 0, conf.CmdUpdateResultType6(), resData, user, []string{user.ConnID()})

	logger.LogFile(nxLog.LevelInfo, fmt.Sprintf("%s complete roomID=%d, rpcProto=%+v, resData=%+v", logPrefix, user.RoomID(), rpcProto, resData))
	return
}
