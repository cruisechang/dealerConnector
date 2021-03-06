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

type updateResultType1Processor struct {
	BasicProcessor
}

func NewUpdateResultType1Processor(processor BasicProcessor) (*updateResultType1Processor, error) {
	p := &updateResultType1Processor{
		BasicProcessor: processor,
	}

	return p, nil

}

func (p *updateResultType1Processor) Run(obj *nex.CommandObject) (resErr error) {
	logger := p.GetLogger()
	conf := p.GetConfigurer()
	user := obj.User
	logPrefix := "2708 updateResultType1"

	defer func() {
		if r := recover(); r != nil {
			logger.LogFile(nxLog.LevelPanic, fmt.Sprintf("%s panic:%v", logPrefix, r))
			resErr = fmt.Errorf("%s panic %v", logPrefix, r)
		}
	}()

	//parsing command data
	deStr, err := base64.StdEncoding.DecodeString(obj.Cmd.Data)

	if err != nil {
		p.SendCommand(config.CodeBase64DecodeFailed, 0, conf.CmdUpdateResultType1(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s base64 decode cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	//取 data
	data := []config.UpdateResultType1CmdData{}

	if err := json.Unmarshal(deStr, &data); err != nil {
		p.SendCommand(config.CodeJsonUnmarshalJsonFailed, 0, conf.CmdUpdateResultType1(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json unmarshal cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	//rpc to server
	rpc := p.GetRPCManager()

	resData := &pb.UpdateResultType1Data{
		RoomID:         int32(user.RoomID()),
		Round:          data[0].Round,
		Result:         data[0].Result,
		DragonPoker:    data[0].DragonPoker,
		TigerPoker:     data[0].TigerPoker,
		DragonOddEven:  data[0].DragonOddEven,
		DragonRedBlack: data[0].DragonRedBlack,
		TigerOddEven:   data[0].TigerOddEven,
		TigerRedBlack:  data[0].TigerRedBlack,
	}

	_, err = rpc.UpdateResultType1("c0", resData)
	if err != nil {
		p.SendCommand(config.CodeRPCError, 0, conf.CmdUpdateResultType1(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s rpc error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))

		resErr = err
		return
	}

	p.SendCommand(config.CodeSuccess, 0, conf.CmdUpdateResultType1(), p.EmptySendDataStr(), user, []string{user.ConnID()})

	logger.LogFile(nxLog.LevelInfo, fmt.Sprintf("%s complete roomID=%d, resData=%+v", logPrefix, user.RoomID(), resData))
	return
}
