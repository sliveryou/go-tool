package cipher

import (
	"testing"
)

const (
	aesCbcKey = "uy4ymckgirj3nverpa67vsqp7gbf1yg2"
	aesCbcIv  = "ewfrq37gka4w7pf1"
)

func TestNewAesCbc(t *testing.T) {
	type args struct {
		key string
		iv  string
	}
	tests := []struct {
		name    string
		args    args
		want    Cipher
		wantErr bool
	}{
		{
			name: "suc",
			args: args{
				key: aesCbcKey,
				iv:  aesCbcIv,
			},
			wantErr: false,
		},
		{
			name: "errkey",
			args: args{
				key: "errkey",
				iv:  aesCbcIv,
			},
			wantErr: true,
		},
		{
			name: "erriv",
			args: args{
				key: aesCbcKey,
				iv:  "erriv",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewAesCbc(tt.args.key, tt.args.iv)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesCbc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
