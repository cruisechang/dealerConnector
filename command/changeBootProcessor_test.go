package command

import (
	"testing"

	"encoding/base64"
	"encoding/json"

	"github.com/cruisechang/dealerConnector/config"
	"github.com/cruisechang/dealerConnector/rpc"
	"github.com/cruisechang/nex"
	"github.com/cruisechang/nex/entity"
)

func Test_changeBootProcessor_Run(t *testing.T) {

	nx, _ := nex.NewNex(getConfigFilePosition("nexConfig.json"))
	conf, _ := config.NewConfigurer("config.json")
	rpc := rpc.NewRPCManager(nx)

	rpcClients := nx.GetConfig().GetRPCClients()
	for _, v := range rpcClients {
		_, err := rpc.AddClient(v.Address, v.Port, v.Name)
		if err != nil {
			t.Fatalf("Test_changeBootProcessor_Run rpc addClent error=%s", err.Error())
		}
	}

	lp, _ := NewChangeBootProcessor(NewBasicProcessor(nx, conf,  rpc))

	user := entity.NewUser(0, "conn")
	if room, ok := nx.GetRoomManager().GetRoom(200); ok {
		room.AddUser(user)
	}

	obj := &[]config.ChangeBootCmdData{
		config.ChangeBootCmdData{},
	}
	c, err := json.Marshal(obj)
	if err != nil {
		t.Logf("%s", err.Error())
	}
	cs := base64.StdEncoding.EncodeToString(c)

	cmd := &nex.CommandObject{
		Cmd: &nex.Command{
			0,
			conf.CmdChangeBoot(),
			0,
			cs,
		},
		User: user,
	}

	errCmd := &nex.CommandObject{
		Cmd: &nex.Command{
			0,
			conf.CmdChangeBoot(),
			0,
			"",
		},
		User: user,
	}

	errCmd2 := &nex.CommandObject{
		Cmd: &nex.Command{
			0,
			conf.CmdChangeBoot(),
			0,
			cs,
		},
		User: nil,
	}

	type args struct {
		obj *nex.CommandObject
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "0",
			args:    args{cmd},
			wantErr: false,
		},
		{
			name:    "1",
			args:    args{errCmd},
			wantErr: true,
		},
		{
			name:    "2",
			args:    args{errCmd2},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := lp.Run(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("changeBootProcessor.Run() error = %v, wantErr %v, name=%s", err, tt.wantErr, tt.name)
			}
		})
	}
}
