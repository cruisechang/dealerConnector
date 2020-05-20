package rpc

import (
	"testing"

	pb "github.com/cruisechang/liveServer/protobuf"

	"bytes"
	"os"
	"path/filepath"

	nexSpace "github.com/cruisechang/nex"
)

func Test_clientManager_GetRoomInfo(t *testing.T) {

	nx, err := nexSpace.NewNex(getConfigFilePosition("nexConfig.json"))
	if err != nil {
		t.Fatalf("newNex error=%s", err.Error())
	}

	rpc := NewRPCManager(nx)

	rpcClients := nx.GetConfig().GetRPCClients()
	for _, v := range rpcClients {
		_, err := rpc.AddClient(v.Address, v.Port, v.Name)
		if err != nil {
			t.Fatalf("range rpcClients error=%s", err.Error())
		}
	}

	//rpcRes, err := rpc.GetRoomInfo("c0", &pb.GetRoomInfoData{RoomID: int64(user.RoomID()),})

	type args struct {
		clientName string
		data       *pb.GetRoomInfoData
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.GetRoomInfoRes
		wantErr bool
	}{
		{
			name: "0",
			args: args{
				"c0",
				&pb.GetRoomInfoData{RoomID: 1},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "1",
			args: args{
				"c0",
				&pb.GetRoomInfoData{RoomID: 2},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "3",
			args: args{
				"c0",
				&pb.GetRoomInfoData{RoomID: 3},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := rpc.GetRoomInfo(tt.args.clientName, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("clientManager.GetRoomInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("clientManager.GetRoomInfo() = %v, want %v", got, tt.want)
			//}

			t.Logf("getRoomInfo=%+v,name=%s", got, tt.name)
		})
	}
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
