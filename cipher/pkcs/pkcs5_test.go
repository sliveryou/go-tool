package pkcs

import (
	"reflect"
	"testing"
)

var encrypt_8 = []byte{97, 115, 100, 102, 4, 4, 4, 4}

func TestPKCS5Padding(t *testing.T) {
	type args struct {
		cipherText []byte
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
			},
			want: encrypt_8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PKCS5Padding(tt.args.cipherText); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PKCS5Padding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPKCS5Trimming(t *testing.T) {
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
				encrypt: encrypt_8,
			},
			want:    cipherText,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PKCS5Trimming(tt.args.encrypt)
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
