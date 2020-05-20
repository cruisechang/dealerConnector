package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	//value
	clientReady    = 1
	clientNotReady = 0

	unAssigned = -1

	hallID0  = 0
	hallID1  = 1
	hallID10 = 10

	//command
	cmdRoomLogin    = "2011"
	cmdOnlineNotify = "2013"
	cmdGetRoomInfo  = "2001"
	cmdWaiting      = "2002"
	cmdBeginBetting = "2005"
	cmdChangeBoot   = "2007"
	cmdDealerLogin  = "2015"
	cmdEndBetting   = "2006"
	cmdHeartbeat    = "2012"
	cmdReconnect    = "2016"
	cmdCancelRound  = "2017"

	//baccarat
	cmdHistoryResultType0 = "2609"
	cmdRoundResultType0   = "2604"
	cmdUpdateResultType0  = "2608"
	cmdRoundProcess0      = "2614"

	// dragonTiger
	cmdHistoryResultType1 = "2709"
	cmdRoundResultType1   = "2704"
	cmdUpdateResultType1  = "2708"
	cmdRoundProcess1      = "2714"

	//sicbo
	cmdHistoryResultType6 = "2509"
	cmdRoundResultType6   = "2504"
	cmdUpdateResultType6  = "2508"
	cmdRerollDice         = "2510"

	//rolette
	cmdHistoryResultType7 = "2809"
	cmdRoundResultType7   = "2804"
	cmdUpdateResultType7  = "2808"
	cmdRethrow            = "2810"

	//niouniu
	cmdHistoryResultType2 = "2109"
	cmdRoundResultType2   = "2104"
	cmdUpdateResultType2  = "2108"
	cmdRoundProcess2      = "2114"

	roomTypeBaccarat    = 0
	roomTypeDragonTiger = 1
	roomTypeNiuniu      = 2
	roomTypeFantan      = 3
	roomTypeSangong     = 4
	roomTypeThisBar     = 5
	roomTypeSicbo       = 6

	//variable
	hallVarRoomIDs = "roomIDs"

	roomVarType         = "rvtpy"
	roomVarHLSURL       = "rvuls"
	roomVarBetMin       = "rvbi"
	roomVarBetMax       = "rvbm"
	roomVarBoot         = "rvbo"
	roomVarRound        = "rvru"
	roomVarDealer       = "rvde"
	roomVarBetCountDown = "rvbcd"

	userVarUnitID = "pid"
	userVarHallID = "hid"
	userVarRoomID = "rid"
)

type Configurer interface {
	Version() string

	//未指定的id
	UnAssigned() int

	//hall / room
	HallID0() int
	HallID1() int
	HallID10() int

	//common command
	CmdRoomLogin() string
	CmdOnlineNotify() string
	CmdGetRoomInfo() string
	CmdWaiting() string
	CmdChangeBoot() string
	CmdBeginBetting() string
	CmdDealerLogin() string
	CmdEndBetting() string
	CmdHeartbeat() string
	CmdReconnect() string
	CmdCancelRound() string

	// baccarat
	CmdHistoryResultType0() string
	CmdRoundResultType0() string
	CmdUpdateResultType0() string
	CmdRoundProcess0() string

	//dragon tiger
	CmdHistoryResultType1() string
	CmdRoundResultType1() string
	CmdUpdateResultType1() string
	CmdRoundProcess1() string

	//sicbo
	CmdHistoryResultType6() string
	CmdRoundResultType6() string
	CmdUpdateResultType6() string
	CmdRerollDice() string

	//rolette
	CmdHistoryResultType7() string
	CmdRoundResultType7() string
	CmdUpdateResultType7() string
	CmdRethrow() string

	//niuniu
	CmdHistoryResultType2() string
	CmdRoundResultType2() string
	CmdUpdateResultType2() string
	CmdRoundProcess2() string

	RoomTypeBaccarat() int
	RoomTypeDragonTiger() int
	RoomTypeNiuniu() int
	RoomTypeSangong() int
	RoomTypeFantan() int
	RoomTypeSicbo() int
	RoomTypeThisBar() int

	//hall variable
	HallVarRoomIDs() string

	RoomVarType() string
	RoomVarHLSURL() string
	RoomVarBetMin() string
	RoomVarBetMax() string
	RoomVarBoot() string
	RoomVarRound() string
	RoomVarDealer() string
	RoomVarBetCountDown() string

	UserVarUnitID() string
	UserVarHallID() string
	UserVarRoomID() string
}

//Config config main struct
type configurer struct {
	data *configData
}

type configData struct {
	Version string
}

//NewConfig make a new config struct
func NewConfigurer(configFileName string) (Configurer, error) {

	defer func() {
		if r := recover(); r != nil {
			//p.GetLogger().LogFile(nxLog.LevelPanic, fmt.Sprintf("CountLinkNum panic:%v\n", r))
			log.Printf("NewConfigurer panic=%v", r)
		}
	}()

	cf := &configurer{
		data: &configData{},
	}

	//config
	path := cf.getConfigFilePosition(configFileName)

	_, err := cf.loadConfig(path, cf.data)

	if err != nil {
		return nil, err
	}

	return cf, err
}

func (c *configurer) getConfigFilePosition(fileName string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
	}

	var buf bytes.Buffer
	buf.WriteString(dir)
	buf.WriteString("/")
	buf.WriteString(fileName)

	return buf.String()
}

func (c *configurer) loadConfig(filePath string, container interface{}) (interface{}, error) {

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	//con := &configData{}
	//unmarshal to struct
	if err := json.Unmarshal(b, container); err != nil {
		return nil, err
	}

	return container, nil
}
func (c *configurer) loadFileToStruct(fileName string, container [][][]int) ([][][]int, error) {
	path := c.getConfigFilePosition(fileName)
	_, err := c.loadConfig(path, &container)

	if err != nil {
		return nil, err
	}
	return container, nil
}

func (c *configurer) Version() string {
	return c.data.Version
}

//未指定的id
func (c *configurer) UnAssigned() int { return unAssigned }

//hall / room
func (c *configurer) HallID0() int  { return hallID0 }
func (c *configurer) HallID1() int  { return hallID1 }
func (c *configurer) HallID10() int { return hallID10 }

//command
func (c *configurer) CmdRoomLogin() string    { return cmdRoomLogin }
func (c *configurer) CmdOnlineNotify() string { return cmdOnlineNotify }
func (c *configurer) CmdGetRoomInfo() string  { return cmdGetRoomInfo }

func (c *configurer) CmdWaiting() string      { return cmdWaiting }
func (c *configurer) CmdBeginBetting() string { return cmdBeginBetting }
func (c *configurer) CmdChangeBoot() string   { return cmdChangeBoot }
func (c *configurer) CmdDealerLogin() string  { return cmdDealerLogin }
func (c *configurer) CmdEndBetting() string   { return cmdEndBetting }
func (c *configurer) CmdHeartbeat() string    { return cmdHeartbeat }
func (c *configurer) CmdReconnect() string    { return cmdReconnect }
func (c *configurer) CmdCancelRound() string  { return cmdCancelRound }

//baccarat
func (c *configurer) CmdHistoryResultType0() string { return cmdHistoryResultType0 }
func (c *configurer) CmdRoundResultType0() string   { return cmdRoundResultType0 }
func (c *configurer) CmdUpdateResultType0() string  { return cmdUpdateResultType0 }
func (c *configurer) CmdRoundProcess0() string      { return cmdRoundProcess0 }

//dragon tiger
func (c *configurer) CmdHistoryResultType1() string { return cmdHistoryResultType1 }
func (c *configurer) CmdRoundResultType1() string   { return cmdRoundResultType1 }
func (c *configurer) CmdUpdateResultType1() string  { return cmdUpdateResultType1 }
func (c *configurer) CmdRoundProcess1() string      { return cmdRoundProcess1 }

//sicbo
func (c *configurer) CmdHistoryResultType6() string { return cmdHistoryResultType6 }
func (c *configurer) CmdRoundResultType6() string   { return cmdRoundResultType6 }
func (c *configurer) CmdUpdateResultType6() string  { return cmdUpdateResultType6 }
func (c *configurer) CmdRerollDice() string         { return cmdRerollDice }

//rolette
func (c *configurer) CmdHistoryResultType7() string { return cmdHistoryResultType7 }
func (c *configurer) CmdRoundResultType7() string   { return cmdRoundResultType7 }
func (c *configurer) CmdUpdateResultType7() string  { return cmdUpdateResultType7 }
func (c *configurer) CmdRethrow() string            { return cmdRethrow }

//baccarat
func (c *configurer) CmdHistoryResultType2() string { return cmdHistoryResultType2 }
func (c *configurer) CmdRoundResultType2() string   { return cmdRoundResultType2 }
func (c *configurer) CmdUpdateResultType2() string  { return cmdUpdateResultType2 }
func (c *configurer) CmdRoundProcess2() string      { return cmdRoundProcess2 }

//

func (c *configurer) RoomTypeBaccarat() int    { return roomTypeBaccarat }
func (c *configurer) RoomTypeDragonTiger() int { return roomTypeDragonTiger }
func (c *configurer) RoomTypeNiuniu() int      { return roomTypeNiuniu }
func (c *configurer) RoomTypeSangong() int     { return roomTypeSangong }
func (c *configurer) RoomTypeFantan() int      { return roomTypeFantan }
func (c *configurer) RoomTypeSicbo() int       { return roomTypeSicbo }
func (c *configurer) RoomTypeThisBar() int     { return roomTypeThisBar }

//user variable
func (c *configurer) UserVarUnitID() string { return userVarUnitID }
func (c *configurer) UserVarHallID() string { return userVarHallID }
func (c *configurer) UserVarRoomID() string { return userVarRoomID }

//
func (c *configurer) RoomVarType() string         { return roomVarType }
func (c *configurer) RoomVarHLSURL() string       { return roomVarHLSURL }
func (c *configurer) RoomVarBetMin() string       { return roomVarBetMin }
func (c *configurer) RoomVarBetMax() string       { return roomVarBetMax }
func (c *configurer) RoomVarBoot() string         { return roomVarBoot }
func (c *configurer) RoomVarRound() string        { return roomVarRound }
func (c *configurer) RoomVarDealer() string       { return roomVarDealer }
func (c *configurer) RoomVarBetCountDown() string { return roomVarBetCountDown }

//hall variable
func (c *configurer) HallVarRoomIDs() string { return hallVarRoomIDs }
