package builtinEvent

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/cruisechang/dealerConnector/config"
	"github.com/cruisechang/dealerConnector/rpc"
	"github.com/cruisechang/nex"
	"github.com/cruisechang/nex/entity"
	nxLog "github.com/cruisechang/nex/log"
)

type BasicProcessor interface {
	SendCommand(errorCode, step int, cmdName, dataStr string, sender entity.User, receiveConnID []string)
	DisconnectUser(userID int)
	RemoveUser(userID int)
	GetLogger() nxLog.Logger
	GetConfigurer() config.Configurer
	GetHallManager() nex.HallManager
	GetRoomManager() nex.RoomManager
	GetRPCManager() rpc.ClientManager

	EmptySendDataStr() string
}

//BasicProcessor is parent struct for process.
type basicProcessor struct {
	nex              nex.Nex
	conf             config.Configurer
	rpcManager       rpc.ClientManager
	emptySendDataStr string
}

func NewBasicProcessor(nex nex.Nex, conf config.Configurer, rpc rpc.ClientManager) BasicProcessor {
	p := &basicProcessor{
		nex:        nex,
		conf:       conf,
		rpcManager: rpc,
	}

	//response data
	resData := []config.EmptyResData{config.EmptyResData{}}
	b, err := json.Marshal(resData)
	if err != nil {
		return p
	}

	p.emptySendDataStr = base64.StdEncoding.EncodeToString(b)

	return p
}

func (b *basicProcessor) GetLogger() nxLog.Logger {
	return b.nex.GetLogger()
}

func (b *basicProcessor) GetHallManager() nex.HallManager {
	return b.nex.GetHallManager()
}

func (b *basicProcessor) GetRoomManager() nex.RoomManager {
	return b.nex.GetRoomManager()
}

func (b *basicProcessor) Print(msg string) {
	fmt.Printf(msg)
}

func (b *basicProcessor) SendCommand(errorCode, step int, cmdName, dataStr string, sender entity.User, receiveConnID []string) {

	resCmd, _ := b.nex.CreateCommand(errorCode, step, cmdName, dataStr)

	b.nex.SendCommand(resCmd, sender, receiveConnID, true)
}

func (b *basicProcessor) DisconnectUser(userID int) {
	b.nex.DisconnectUser(userID)
}

func (b *basicProcessor) RemoveUser(userID int) {
	b.nex.RemoveUser(userID)
}

func (b *basicProcessor) EmptySendDataStr() string {
	return b.emptySendDataStr
}
func (b *basicProcessor) GetConfigurer() config.Configurer {
	return b.conf
}

func (b *basicProcessor) GetRPCManager() rpc.ClientManager {
	return b.rpcManager
}
