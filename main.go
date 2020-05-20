package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cruisechang/dealerConnector/builtinEvent"
	"github.com/cruisechang/dealerConnector/command"
	"github.com/cruisechang/dealerConnector/config"
	"github.com/cruisechang/dealerConnector/rpc"
	nexSpace "github.com/cruisechang/nex"
	nexBuiltinEvent "github.com/cruisechang/nex/builtinEvent"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
		}
	}()

	nx, err := nexSpace.NewNex(getConfigFilePosition("nexConfig.json"))
	if err != nil {
		exit(1, errors.New(fmt.Sprintf("NewNex error:%s", err.Error())))

	}

	//configure 必須是同一個
	conf, err := config.NewConfigurer("config.json")
	if err != nil {
		exit(2, errors.New(fmt.Sprintf("loadConfig error:%s", err.Error())))
	}

	rpc := rpc.NewRPCManager(nx)

	rpcClients := nx.GetConfig().GetRPCClients()
	for _, v := range rpcClients {
		_, err := rpc.AddClient(v.Address, v.Port, v.Name)
		if err != nil {
			exit(3, errors.New(fmt.Sprintf("rcp addClient error:%s", err.Error())))
		}
	}

	dl, err := command.NewDealerLoginProcessor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(1, err)
	}

	rl, err := command.NewRoomLoginProcessor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(2, err)
	}

	onnp, err := command.NewOnlineNotifyProcessor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(3, err)
	}
	grip, err := command.NewGetRoomInfoProcessor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(4, err)
	}
	wp, err := command.NewWaitingProcessor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(5, err)
	}
	bb, err := command.NewBeginBettingProcessor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(6, err)
	}

	cb, err := command.NewChangeBootProcessor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(8, err)
	}

	eb, err := command.NewEndBettingProcessor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(9, err)
	}

	hbp, err := command.NewHeartbeatProcessor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(10, err)
	}

	rect, err := command.NewReconnectProcessor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(100, err)
	}

	canr, err := command.NewCancelRoundProcessor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(100, err)
	}

	//type0
	rpp0, err := command.NewRoundProcess0(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(00, err)
	}
	rrbadtp0, err := command.NewRoundResultType0Processor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(01, err)
	}

	urbap0, err := command.NewUpdateResultType0Processor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(02, err)
	}
	grhbadt0, err := command.NewHistoryResultType0Processor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(03, err)
	}

	//type1
	rpp1, err := command.NewRoundProcess1(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(11, err)
	}
	rrbadtp1, err := command.NewRoundResultType1Processor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(12, err)
	}

	urbap1, err := command.NewUpdateResultType1Processor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(13, err)
	}
	grhbadt1, err := command.NewHistoryResultType1Processor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(7, err)
	}

	//type2
	rpp2, err := command.NewRoundProcess2(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(00, err)
	}
	rrbadtp2, err := command.NewRoundResultType2Processor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(20, err)
	}

	urbap2, err := command.NewUpdateResultType2Processor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(21, err)
	}
	grhbadt2, err := command.NewHistoryResultType2Processor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(22, err)
	}

	//type6
	rrbadtp6, err := command.NewRoundResultType6Processor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(60, err)
	}

	urbap6, err := command.NewUpdateResultType6Processor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(61, err)
	}
	grhbadt6, err := command.NewHistoryResultType6Processor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(62, err)
	}
	rerodice, err := command.NewRerollDiceProcessor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(63, err)
	}

	//type˙
	rrbadtp7, err := command.NewRoundResultType7Processor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(60, err)
	}

	urbap7, err := command.NewUpdateResultType7Processor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(61, err)
	}
	grhbadt7, err := command.NewHistoryResultType7Processor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(62, err)
	}
	rethrow, err := command.NewRethrowProcessor(command.NewBasicProcessor(nx, conf, rpc))
	if err != nil {
		exit(63, err)
	}

	//command
	nx.RegisterCommandProcessor(conf.CmdDealerLogin(), dl)
	nx.RegisterCommandProcessor(conf.CmdRoomLogin(), rl)
	nx.RegisterCommandProcessor(conf.CmdOnlineNotify(), onnp)
	nx.RegisterCommandProcessor(conf.CmdGetRoomInfo(), grip)
	nx.RegisterCommandProcessor(conf.CmdWaiting(), wp)
	nx.RegisterCommandProcessor(conf.CmdBeginBetting(), bb)
	nx.RegisterCommandProcessor(conf.CmdChangeBoot(), cb)
	nx.RegisterCommandProcessor(conf.CmdEndBetting(), eb)
	nx.RegisterCommandProcessor(conf.CmdHeartbeat(), hbp)
	nx.RegisterCommandProcessor(conf.CmdReconnect(), rect)
	nx.RegisterCommandProcessor(conf.CmdCancelRound(), canr)

	//type0
	nx.RegisterCommandProcessor(conf.CmdRoundProcess0(), rpp0)
	nx.RegisterCommandProcessor(conf.CmdRoundResultType0(), rrbadtp0)
	nx.RegisterCommandProcessor(conf.CmdUpdateResultType0(), urbap0)
	nx.RegisterCommandProcessor(conf.CmdHistoryResultType0(), grhbadt0)

	//type1
	nx.RegisterCommandProcessor(conf.CmdRoundProcess1(), rpp1)
	nx.RegisterCommandProcessor(conf.CmdRoundResultType1(), rrbadtp1)
	nx.RegisterCommandProcessor(conf.CmdUpdateResultType1(), urbap1)
	nx.RegisterCommandProcessor(conf.CmdHistoryResultType1(), grhbadt1)

	//type2
	nx.RegisterCommandProcessor(conf.CmdRoundProcess2(), rpp2)
	nx.RegisterCommandProcessor(conf.CmdRoundResultType2(), rrbadtp2)
	nx.RegisterCommandProcessor(conf.CmdUpdateResultType2(), urbap2)
	nx.RegisterCommandProcessor(conf.CmdHistoryResultType2(), grhbadt2)

	//type6
	nx.RegisterCommandProcessor(conf.CmdRoundResultType6(), rrbadtp6)
	nx.RegisterCommandProcessor(conf.CmdUpdateResultType6(), urbap6)
	nx.RegisterCommandProcessor(conf.CmdHistoryResultType6(), grhbadt6)
	nx.RegisterCommandProcessor(conf.CmdRerollDice(), rerodice)

	//type7
	nx.RegisterCommandProcessor(conf.CmdRoundResultType7(), rrbadtp7)
	nx.RegisterCommandProcessor(conf.CmdUpdateResultType7(), urbap7)
	nx.RegisterCommandProcessor(conf.CmdHistoryResultType7(), grhbadt7)
	nx.RegisterCommandProcessor(conf.CmdRethrow(), rethrow)

	//builtin event
	ule, _ := builtinEvent.NewUserLostEventProcessor(command.NewBasicProcessor(nx, conf, rpc))
	nx.RegisterBuiltinEventProcessor(nexBuiltinEvent.EventUserLost, ule)

	nx.Start()

}

func exit(id int, err error) {
	fmt.Println(err)
	log.Println(err)
	os.Exit(id)
}
func profileStatus() {
	//http://localhost:6060/debug/pprof/  to see data
	log.Println(http.ListenAndServe("localhost:6060", nil))
}

func getConfigFilePosition(fileName string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
	}

	var buf bytes.Buffer
	buf.WriteString(dir)
	buf.WriteString("/")
	buf.WriteString(fileName)

	return buf.String()
}
