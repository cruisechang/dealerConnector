package util

import (
	"strconv"
	"time"
)

func GetNewRoundID(roomID, bootID int) (int64, error) {
	tm := time.Now()
	date := tm.Format("20180901")

	if bootID < 10 {

	}
	newID := date + strconv.Itoa(bootID) + strconv.Itoa(roomID)

	return strconv.ParseInt(newID, 10, 64)

}
