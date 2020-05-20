package builtinEvent

import (
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/cruisechang/dealerConnector/config"
	"github.com/cruisechang/dealerConnector/rpc"
	nexSpace "github.com/cruisechang/nex"
	"github.com/cruisechang/nex/entity"
)

var (
	nexForBasicProcessorTester nexSpace.Nex
)

type BasicProcessorTester interface {
	Run(t *testing.T) error
}

type basicProcessorTester struct {
	BasicProcessor
}

func NewBasicProcessorTester(processor BasicProcessor) (BasicProcessorTester, error) {
	p := &basicProcessorTester{
		BasicProcessor: processor,
	}

	return p, nil

}

func TestInheritBasicProcessor(t *testing.T) {

	conf, _ := config.NewConfigurer("config.json")
	rpc := rpc.NewRPCManager(nexForBasicProcessorTester)

	bpt, err := NewBasicProcessorTester(NewBasicProcessor(nexForBasicProcessorTester, conf, rpc))
	if err != nil {

	}

	bpt.Run(t)

}

func (p *basicProcessorTester) Run(t *testing.T) error {

	conf := p.GetConfigurer()

	//sendCommand
	resData := []config.RoomLoginCmdData{config.RoomLoginCmdData{}}

	b, _ := json.Marshal(resData)

	//[]byte encode to base64 string
	sendDataStr := base64.StdEncoding.EncodeToString(b)

	t.Logf("SendCommand %s\n", sendDataStr)
	p.SendCommand(config.CodeJsonUnmarshalJsonFailed, 0, conf.CmdRoomLogin(), sendDataStr, entity.NewUser(0, "connid"), []string{"uuid"})

	return nil
}

func TestPrepare(t *testing.T) {

	n, err := nexSpace.NewNex("nexConfig.json")
	if err != nil {
		t.Errorf("TestLogin error %s \n", err.Error())
		return
	}
	nexForBasicProcessorTester = n

	t.Logf("TestPrepare %#v\n", nexForBasicProcessorTester)
}
