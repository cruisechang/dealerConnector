package builtinEvent

import (
	//"testing"
	//
	//"github.com/cruisechang/dealerConnector/config"
	//hp "github.com/cruisechang/dealerConnector/http"
	//"github.com/cruisechang/dealerConnector/rpc"
	//nexSpace "github.com/cruisechang/nex"
	//nxBuiltinEvent "github.com/cruisechang/nex/builtinEvent"
	//"github.com/cruisechang/nex/entity"
)

//func TestNewUserLostEventProcessor(t *testing.T) {
//
//	nx, err := nexSpace.NewNex("nexConfig.json")
//	if err != nil {
//		t.Errorf("TestLogin error %s \n", err.Error())
//		return
//	}
//	conf, _ := config.NewConfigurer("config.json")
//	req := hp.NewRequester(nx)
//	rpc := rpc.NewRPCManager(nx)
//
//	ule, _ := NewUserLostEventProcessor(NewBasicProcessor(nx, conf, req, rpc))
//
//	user := entity.NewUser(0, "connID")
//	user.SetAccessID("accessID")
//
//	obj := nxBuiltinEvent.EventObject{
//		Code: 1,
//		User: user,
//	}
//
//	t.Logf("%#v", obj)
//
//	ule.Run(&obj)
//
//}
