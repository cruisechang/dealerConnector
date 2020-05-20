package builtinEvent

import (
	"errors"
	"fmt"

	nxBuiltinEvent "github.com/cruisechang/nex/builtinEvent"
	nxLog "github.com/cruisechang/nex/log"
)

type userLostEventProcessor struct {
	BasicProcessor
}

func NewUserLostEventProcessor(processor BasicProcessor) (*userLostEventProcessor, error) {
	p := &userLostEventProcessor{
		BasicProcessor: processor,
	}

	return p, nil
}

func (p *userLostEventProcessor) Run(obj *nxBuiltinEvent.EventObject) error {
	logger := p.GetLogger()
	user := obj.User

	if user == nil {
		logger.LogFile(nxLog.LevelError, fmt.Sprintf("userLostEventProcessor user==nil "))
		return errors.New("userLostEventProcessor user==nil")
	}

	defer func() {
		if r := recover(); r != nil {
			logger.LogFile(nxLog.LevelPanic, fmt.Sprintf("UserLostEventProcessor panic:%v", r))
		}
	}()

	//remove user from room
	if room, ok := p.GetRoomManager().GetRoom(user.RoomID()); ok {
		room.RemoveUser(user)
	}

	//remove user
	p.RemoveUser(user.UserID())

	logger.LogFile(nxLog.LevelInfo, fmt.Sprintf("userLostEventProcessor complete  user id=%d,user=%s", user.UserID(), user.Name()))

	return nil
}
