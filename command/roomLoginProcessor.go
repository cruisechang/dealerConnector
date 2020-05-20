package command

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/cruisechang/dealerConnector/config"
	pb "github.com/cruisechang/liveServer/protobuf"
	"github.com/cruisechang/nex"
	"github.com/cruisechang/nex/entity"
	nxLog "github.com/cruisechang/nex/log"
)

//loginProcessor implements command.Processor
type roomLoginProcessor struct {
	BasicProcessor
}

func NewRoomLoginProcessor(processor BasicProcessor) (*roomLoginProcessor, error) {
	p := &roomLoginProcessor{
		BasicProcessor: processor,
	}

	return p, nil
}

//return error means some thing strange happend
//回傳error是表示程式發生不預期錯誤，
//如果是login 失敗，不回傳錯誤
//failed must  disconnect client, 之後會 logout
func (p *roomLoginProcessor) Run(obj *nex.CommandObject) (resErr error) {

	logPrefix := "2011 roomLogin"
	conf := p.GetConfigurer()
	logger := p.GetLogger()
	rm := p.GetRoomManager()
	user := obj.User

	defer func() {
		if r := recover(); r != nil {
			logger.LogFile(nxLog.LevelPanic, fmt.Sprintf("%s panic:%v", logPrefix, r))
			resErr = errors.New(fmt.Sprintf("%s panic %v", logPrefix, r))
		}
	}()

	//parsing cmd data
	de, err := base64.StdEncoding.DecodeString(obj.Cmd.Data)

	if err != nil {
		p.DisconnectUser(obj.User.UserID())
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s base64 decode cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}

	data := []config.RoomLoginCmdData{}

	if err := json.Unmarshal(de, &data); err != nil {
		p.DisconnectUser(user.UserID())
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s json Unmarshal cmd data error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))
		resErr = err
		return
	}
	//create room, add user(client)

	roomID := data[0].RoomID
	var room entity.Room
	var ok bool

	// room already exist
	room, ok = rm.GetRoom(roomID)
	if !ok {
		room, _ = rm.CreateRoom(roomID, p.getRoomType(roomID), "room"+strconv.Itoa(roomID))
	}

	err = room.AddUser(user)
	if err != nil {
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s room.AddUser error roomID=%d,error=%s", logPrefix, roomID,err.Error()))

		return err
	}
	user.SetName("room" + strconv.Itoa(roomID))

	//rpc game server
	rpc := p.GetRPCManager()

	resData := &pb.RoomLoginData{
		RoomID: int64(data[0].RoomID),
	}

	_, err = rpc.RoomLlogin("c0", resData)
	if err != nil {

		p.DisconnectUser(user.UserID())
		//p.SendCommand(config.CodeRPCError, 0, conf.CmdGetRoomInfo(), p.EmptySendDataStr(), user, []string{user.ConnID()})
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("%s rpc error,user:%s,error:%s", logPrefix, user.Name(), err.Error()))

		resErr = err
		return
	}

	p.SendCommand(config.CodeSuccess, 0, conf.CmdRoomLogin(), p.EmptySendDataStr(), user, []string{user.ConnID()})

	logger.LogFile(nxLog.LevelInfo, fmt.Sprintf("%s complete  roomID=%d resData=%+v", logPrefix, data[0].RoomID, resData))

	return
}

func (p *roomLoginProcessor) getRoomType(roomID int) int {
	if roomID >= 100 && roomID < 200 {
		return 0
	} else if roomID >= 200 && roomID < 300 {
		return 1
	} else if roomID >= 300 && roomID < 400 {
		return 2
	} else if roomID >= 700 && roomID < 800 {
		return 6
	} else if roomID >= 800 && roomID < 900 {
		return 7
	} else {
		return -1
	}
}
