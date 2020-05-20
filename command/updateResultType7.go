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

type updateResultType7Processor struct {
	BasicProcessor
}

func NewUpdateResultType7Processor(processor BasicProcessor) (*updateResultType7Processor, error) {
	p := &updateResultType7Processor{
		BasicProcessor: processor,
	}

	return p, nil

}

func (p *updateResultType7Processor) Run(obj *nex.CommandObject) (resErr error) {
	logger := p.GetLogger()
	conf := p.GetConfigurer()
	user := obj.User
	logPrefix := "2808 updateResultType7"

	defer func() {
		if r := recover(); r != nil {
			logger.LogFile(nxLog.LevelPanic, fmt.Sprintf("%s panic:%v", logPrefix, r))
			resErr = fmt.Errorf("%s panic %v", logPrefix, r)
		}
	}()

	//parsing command data
	deStr, err := base64.StdEncoding.DecodeString(obj.Cmd.Data)

	if err != nil {
		p.SendCommand(config.CodeBase64DecodeFailed, 0, conf.CmdUpdateResultType7(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s base64 decode cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	//取 data
	var data []config.UpdateResultType7CmdData

	if err := json.Unmarshal(deStr, &data); err != nil {
		p.SendCommand(config.CodeJsonUnmarshalJsonFailed, 0, conf.CmdUpdateResultType7(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json unmarshal cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	//rpc to server
	rpc := p.GetRPCManager()

	rpcProto := &pb.UpdateResultType7Data{
		RoomID:   int32(user.RoomID()),
		Round:    data[0].Round,
		Result:   data[0].Result,
		BigSmall: data[0].BigSmall,
		OddEven:  data[0].OddEven,
		RedBlack: data[0].RedBlack,
		Dozen:    data[0].Dozen,
		Column:   data[0].Column,
	}

	_, err = rpc.UpdateResultType7("c0", rpcProto)
	if err != nil {
		p.SendCommand(config.CodeRPCError, 0, conf.CmdUpdateResultType7(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s rpc error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))

		resErr = err
		return
	}

	resData := p.EmptySendDataStr()
	p.SendCommand(config.CodeSuccess, 0, conf.CmdUpdateResultType7(), resData, user, []string{user.ConnID()})

	logger.LogFile(nxLog.LevelInfo, fmt.Sprintf("%s complete roomID=%d, rpcProto=%+v, resData=%+v", logPrefix, user.RoomID(), rpcProto, resData))
	return
}
