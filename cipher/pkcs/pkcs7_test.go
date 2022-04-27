package pkcs

import (
	"reflect"
	"testing"
)

var (
	cipherText = []byte("asdf")
	encrypt_16 = []byte{
		97, 115, 100, 102, 12, 12,
		12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	}
)

func TestPKCS7Padding(t *testing.T) {
	type args struct {
		cipherText []byte
		blockSize  int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "",
			args: args{
				cipherText: cipherText,
				blockSize:  16,
			},
			want: encrypt_16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PKCS7Padding(tt.args.cipherText, tt.args.blockSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PKCS5Padding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPKCS7Trimming(t *testing.T) {
	type args struct {
		encrypt []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "",
			args: args{
				encrypt: encrypt_16,
			},
			want:    cipherText,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PKCS7Trimming(tt.args.encrypt)
			if (err != nil) != tt.wantErr {
				t.Errorf("PKCS5Trimming() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PKCS5Trimming() got = %v, want %v", got, tt.want)
			}
		})
	}
}
