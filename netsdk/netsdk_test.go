package netsdk

import (
	"testing"

	"github.com/yudai/pp"
)

func TestMain(m *testing.M) {
	Init()
	defer Cleanup()

	m.Run()
}

func TestVersion(t *testing.T) {
	t.Logf(SDKVersion())
}

func TestLogin(t *testing.T) {
	type args struct {
		addr string
		user string
		pass string
	}
	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{{
		name: "login with correct admin, pass",
		args: args{
			addr: "192.168.1.64:8000",
			user: "admin",
			pass: "adm89679005",
		},
		wantErr: false,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Login(tt.args.addr, tt.args.user, tt.args.pass)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%s", string(got.DeviceInfo.ST_sSerialNumber[:]))
			pp.Print(got)
			got.Logout()
		})
	}
}
