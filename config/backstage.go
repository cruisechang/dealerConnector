package config

/*
const (
	PathLogin       = "gameLogin"
	PathLogout      = "gameLogout"
	PathGetGameInfo = "gameGetGameInfo"
	PathGetUserInfo = "gameGetUserInfo"
	PathUserBetLog  = "gameUserBetLog"
	PathHeartbeat   = "gameHeartbeat"

	PathGetPoolSetting = "gameGetPoolSetting"
	PathGetPool        = "gameGetPool"
	PathJackpotWin     = "gameJackpotWin"
)

//backstageResContainer is a res wrapper
type BackstageResContainer struct {
	Code   int    `json:"code"`
	Reason string `json:"reason"`
	Data   string `json:"data"`
}

type BackstageLoginParams struct {
	GameID    string `json:"gameId"`
	SessionID string `json:"sessionId"`
	Lang      string `json:"lang"`
}

type BackstageLoginResData struct {
	AccessID  string `json:"accessId"`
	PartnerID int    `json:"partnerId"`
	Name      string `json:"name"`
}

type BackstageLogoutParams struct {
	AccessID string `json:"accessId"`
}
type BackstageLogoutResData struct {
}

type BackstageGetGameInfoParams struct {
	GameID string `json:"gameId"`
}

type BackstageGetGameInfoResData struct {
	Active     int    `json:"active"`
	Name       string `json:"name"`
	CategoryID int    `json:"categoryId"`
}

type BackstageGetUserInfoParams struct {
	PartnerID int    `json:"partnerId"`
	Name      string `json:"name"`
}

type BackstageGetUserInfoResData struct {
	Name          string  `json:"name"`
	Credit        float32 `json:"credit"`
	PartnerID     int     `json:"partnerId"`
	PresentCredit float32 `json:"presentCredit"`
	IsTest        int     `json:"isTest"`
	LastLogin     string  `json:"lastLogin"`
}

type BackstageUserBetLogParams struct {
	AccessID         string  `json:"accessId"`
	Mode             int     `json:"mode"`
	BetCredit        float32 `json:"betCredit"`
	ResultCredit     float32 `json:"resultCredit"`
	BalanceCredit    float32 `json:"balanceCredit"`
	ActiveBetCredit  float32 `json:"activeBetCredit"`
	CommissionCredit float32 `json:"commissionCredit"`
	ResultData       string  `json:"resultData"`
	Time             string  `json:"time"`
	JackpotWinCredit float32 `json:"jackpotWinCredit"`
	JackpotWinPool   int     `json:"jackpotWinPool"`
}
type BackstageUserBetLogResData struct {
	LogID string `json:"logId"`
}

type BackstageHeartbeatParams struct {
	AccessID string `json:"accessId"`
}
type BackstageHeartbeatResData struct {
	Expire string `json:"expire"`
}

//get pool setting
type BackstageGetPoolSettingParam struct {
}
type BackstageGetPoolSettingResData struct {
	Jackpot         float32 `json:"jackpot"`
	JackpotReserved float32 `json:"jackpotReserved"`
	Return          float32 `json:"return"`
	Kill            []BackstageGetPoolSettingResDataKill `json:'kill'`
}

type BackstageGetPoolSettingResDataKill struct {
	PartnerID int     `json:"partnerId"`
	KillRatio      float32 `json:"rate"`
}

//get pool

type BackstageGetPoolParam struct {
}
type BackstageGetPoolResData struct {
	Jackpot         []BackstageGetPoolResDataJackpot `json:"jackpot"`
	JackpotReserved float32                          `json:"jackpotReserved"`
	Return          float32                          `json:"return"`
}
type BackstageGetPoolResDataJackpot struct {
	PoolID int     `json:"poolId"`
	Credit float32 `json:"credit"`
	Time   string  `json:"time"`
}

//jackpot win
type BackstageJackpotWinParam struct {
	AccessID string  `json:"accessId"`
	PoolID   int     `json:"poolId"`
	Rate     float32 `json:"rate"`
}

//result 0=false, 1=success
type BackstageJackpotWinResData struct {
	Result int     `json:"result"`
	Credit float32 `json:"credit"`
}
*/