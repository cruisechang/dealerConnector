package config

const (
	//code for client
	CodeSuccess = 0
	CodeFailed  = 1

	//encode
	//解析json錯誤
	CodeJsonUnmarshalJsonFailed = 100
	//編碼成json錯誤
	CodeMarshalJsonFailed = 101

	CodeBase64DecodeFailed = 102
	CodeBase64EncodeFailed = 103

	CodeRPCError = 104

	CodeRoomExist    = 150
	CodeRoomNotFound = 151
)

//command data
type EmptyResData struct {
}
type DealerLoginCmdData struct {
	RoomID   int    `json:"roomID"`
	Dealer   int    `json:"dealer"`
	Password string `json:"password"`
}

type DealerLoginResData struct {
	Success int `json:"success"`
}

type RoomLoginCmdData struct {
	RoomID int `json:"roomID"`
}

type OnlineNotifyCmdData struct {
}

type GetRoomInfoCmdData struct {
}
type GetRoomInfoResData struct {
	Boot                int    `json:"bootnum"`
	Round               int    `json:"inningnum"`
	RoomID              int    `json:"desknum"`
	RoomName            string `json:"deskname"`
	BankerPlayerMax     int    `json:"bankplaymax"`
	BankerPlayerMin     int    `json:"bankplaymin"`
	TieMax              int    `json:"tiemax"`
	TieMin              int    `json:"tiemin"`
	BankerPlayerPairMax int    `json:"bankplaypairmax"`
	BankerPlayerPairMin int    `json:"bankplaypairmin"`
	Online              int    `json:"online"`
	BetCountDown        int    `json:"bettime"`
}

type WaitingCmdData struct {
	Boot  int `json:"boot"`
	Round int `json:"inning"`
}
type WaitingResData struct {
}

type BeginBettingCmdData struct {
	Boot  int `json:"boot"`
	Round int `json:"inning"`
}
type BeginBettingResData struct {
	Round int64 `json:"inning"`
}

type ChangeBootCmdData struct {
}
type ChangeBootResData struct {
	Boot int `json:"boot"`
}

type EndBettingCmdData struct {
}
type WaitingResultCmdData struct {
}

type ReconnectCmdData struct {
	RoomID int `json:"roomID"`
}

type ReconnectResData struct {
	Success int `json:"success"`
}

type HeartbeatCmdData struct {
}
type CancelRoundCmdData struct {
	Boot  int   `json:"boot"`
	Round int64 `json:"inning"`
}

//baccarat
type RoundResultType0CmdData struct {
	Boot        int32   `json:"boot"`
	Round       int64   `json:"inning"`
	Result      int32   `json:"result"`
	BankerPoker []int32 `json:"bankerPoker"`
	PlayerPoker []int32 `json:"playerPoker"`
	BankerPair  int32   `json:"bankerPair"`
	PlayerPair  int32   `json:"playerPair"`
	BigSmall    int32   `json:"bigSmall"`
	AnyPair     int32   `json:"anyPair"`
	PerfectPair int32   `json:"perfectPair"`
	Super6      int32   `json:"super6"`
	BankerPoint int32   `json:"BankerPoint"`
	PlayerPoint int32   `json:"PlayerPoint"`
}
type UpdateResultType0CmdData struct {
	Round       int64   `json:"inning"`
	Result      int32   `json:"result"`
	BankerPoker []int32 `json:"bankerPoker"`
	PlayerPoker []int32 `json:"playerPoker"`
	BankerPair  int32   `json:"bankerPair"`
	PlayerPair  int32   `json:"playerPair"`
	BigSmall    int32   `json:"bigSmall"`
	AnyPair     int32   `json:"anyPair"`
	PerfectPair int32   `json:"perfectPair"`
	Super6      int32   `json:"super6"`
	BankerPoint int32   `json:"BankerPoint"`
	PlayerPoint int32   `json:"PlayerPoint"`
}
type RoundProcess0CmdData struct {
	Boot  int32 `json:"boot"`
	Round int64 `json:"inning"`
	Owner int32 `json:"owner"`
	Index int32 `json:"index"`
	Poker int32 `json:"pork"`
}
type HistoryResultType0CmdData struct {
	Boot  int `json:"boot"`
	Round int `json:"inning"`
}
type HistoryResultType0ResData struct {
	Result [][]int32 `json:"result"`
}

//dragonTiger
type RoundResultType1CmdData struct {
	Boot           int32 `json:"boot"`
	Round          int64 `json:"inning"`
	Result         int32 `json:"result"`
	DragonPoker    int32 `json:"dragonPoker"`
	TigerPoker     int32 `json:"tigerPoker"`
	DragonOddEven  int32 `json:"dragonOddEven"`
	DragonRedBlack int32 `json:"dragonRedBlack"`
	TigerOddEven   int32 `json:"tigerOddEven"`
	TigerRedBlack  int32 `json:"tigerRedBlack"`
}

type UpdateResultType1CmdData struct {
	Round          int64 `json:"inning"`
	Result         int32 `json:"result"`
	DragonPoker    int32 `json:"dragonPoker"`
	TigerPoker     int32 `json:"tigerPoker"`
	DragonOddEven  int32 `json:"dragonOddEven"`
	DragonRedBlack int32 `json:"dragonRedBlack"`
	TigerOddEven   int32 `json:"tigerOddEven"`
	TigerRedBlack  int32 `json:"tigerRedBlack"`
}
type RoundProcess1CmdData struct {
	Boot  int32   `json:"boot"`
	Round int64 `json:"inning"`
	Owner int32   `json:"owner"`
	Index int32   `json:"index"`
	Poker int32   `json:"pork"`
}
type HistoryResultType1CmdData struct {
	Boot  int32 `json:"boot"`
	Round int64 `json:"inning"`
}
type HistoryResultType1ResData struct {
	Result []int32 `json:"result"`
}

//sicbo
type RoundResultType6CmdData struct {
	Boot     int32     `json:"boot"`
	Round    int64     `json:"inning"`
	Dice     []int32   `json:"dice"`
	Sum      int32     `json:"sum"`
	BigSmall int32     `json:"bigSmall"`
	OddEven  int32     `json:"oddEven"`
	Triple   int32     `json:"triple"`
	Pair     int32     `json:"pair"`
	Paigow   [][]int32 `json:"paigow"`
}
type UpdateResultType6CmdData struct {
	Round    int64     `json:"inning"`
	Dice     []int32   `json:"dice"`
	Sum      int       `json:"sum"`
	BigSmall int       `json:"bigSmall"`
	OddEven  int       `json:"oddEven"`
	Triple   int       `json:"triple"`
	Pair     int       `json:"pair"`
	Paigow   [][]int32 `json:"paigow"`
}
type HistoryResultType6CmdData struct {
	Boot  int   `json:"boot"`
	Round int64 `json:"inning"`
}
type HistoryResultType6ResData struct {
	Result []int32 `json:"result"`
}
type RerollDiceCmdData struct {
	Boot  int   `json:"boot"`
	Round int64 `json:"inning"`
}

//rolette
type RoundResultType7CmdData struct {
	Boot     int32 `json:"boot"`
	Round    int64 `json:"inning"`
	Result   int32 `json:"result"`
	BigSmall int32 `json:"bigSmall"`
	OddEven  int32 `json:"oddEven"`
	RedBlack int32 `json:"redBlack"`
	Dozen    int32 `json:"dozen"`
	Column   int32 `json:"column"`
}
type UpdateResultType7CmdData struct {
	Boot     int32 `json:"boot"`
	Round    int64 `json:"inning"`
	Result   int32 `json:"result"`
	BigSmall int32 `json:"bigSmall"`
	OddEven  int32 `json:"oddEven"`
	RedBlack int32 `json:"redBlack"`
	Dozen    int32 `json:"dozen"`
	Column   int32 `json:"column"`
}
type HistoryResultType7CmdData struct {
	Boot  int   `json:"boot"`
	Round int64 `json:"inning"`
}
type HistoryResultType7ResData struct {
	Result []int32 `json:"result"`
}
type RethrowCmdData struct {
	Boot  int   `json:"boot"`
	Round int64 `json:"inning"`
}

//niuniu
type RoundResultType2CmdData struct {
	Boot   int32                 `json:"boot"`
	Round  int64                 `json:"inning"`
	Head   int32                 `json:"head"`
	Owner0 RoundResultType2Owner `json:"owner0"`
	Owner1 RoundResultType2Owner `json:"owner1"`
	Owner2 RoundResultType2Owner `json:"owner2"`
	Owner3 RoundResultType2Owner `json:"owner3"`
}
type RoundResultType2Owner struct {
	Result  int32   `json:"result`
	Poker   []int32 `json:"poker`
	Pattern int32   `json:"pattern`
}

type UpdateResultType2CmdData struct {
	Round  int64                 `json:"inning"`
	Head   int32                 `json:"head"`
	Owner0 RoundResultType2Owner `json:"owner0"`
	Owner1 RoundResultType2Owner `json:"owner1"`
	Owner2 RoundResultType2Owner `json:"owner2"`
	Owner3 RoundResultType2Owner `json:"owner3"`
}

type HistoryResultType2CmdData struct {
	Boot  int32 `json:"boot"`
	Round int64 `json:"inning"`
}
type HistoryResultType2ResData struct {
	Result [][]int32 `json:"result"`
}

type RoundProcess2CmdData struct {
	Boot  int32 `json:"boot"`
	Round int64 `json:"inning"`
	Owner int32 `json:"owner"`
	Index int32 `json:"index"`
	Poker int32 `json:"pork"`
}
