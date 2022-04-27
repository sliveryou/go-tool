package aes

import (
	"reflect"
	"sync"
	"testing"
)

const (
	commonKey128 = "uy4ymckgirj3nvas"
	commonKey192 = "uy4ymckgirj3nvasirj3nvas"
	commonKey256 = "uy4ymckgirj3nverpa67vsqp7gbf1yg2"
	commonIV     = "ewfrq37gka4w7pf1"
	commonSrc    = "asdf"
	commonSrc2   = "asdf123"
)

var commonEncrypted128 = []byte{
	172, 217, 105, 207, 247, 197,
	214, 227, 12, 17, 251, 226, 136, 138, 160, 230,
}

var commonEncrypted192 = []byte{
	107, 71, 166, 213, 114, 80,
	117, 24, 190, 108, 119, 132, 232, 175, 252, 2,
}

var commonEncrypted256 = []byte{
	27, 178, 56, 176, 21,
	138, 251, 209, 191, 219, 78, 171, 73, 107, 151, 108,
}

var commonEncrypted256_2 = []byte{
	211, 81, 185, 216, 139, 42,
	234, 201, 232, 51, 75, 16, 188, 228, 40, 21,
}

var (
	commonEncrypted128Hex = "acd969cff7c5d6e30c11fbe2888aa0e6"
	commonEncrypted192Hex = "6b47a6d572507518be6c7784e8affc02"
	commonEncrypted256Hex = "1bb238b0158afbd1bfdb4eab496b976c"
)

var (
	commonEncrypted128Base64 = "rNlpz/fF1uMMEfviiIqg5g=="
	commonEncrypted192Base64 = "a0em1XJQdRi+bHeE6K/8Ag=="
	commonEncrypted256Base64 = "G7I4sBWK+9G/206rSWuXbA=="
)

var (
	aescbc128, _ = NewAesCbc([]byte(commonKey128), []byte(commonIV))
	aescbc192, _ = NewAesCbc([]byte(commonKey192), []byte(commonIV))
	aescbc256, _ = NewAesCbc([]byte(commonKey256), []byte(commonIV))
)

func TestAesCbc_Decrypt(t *testing.T) {
	type fields struct {
		a *AesCbc
	}
	type args struct {
		in []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:   "TestAesCbc_Decrypt-128",
			fields: fields{a: aescbc128},
			args: args{
				in: commonEncrypted128,
			},
			want:    []byte(commonSrc),
			wantErr: false,
		},
		{
			name:   "TestAesCbc_Decrypt-192",
			fields: fields{a: aescbc192},
			args: args{
				in: commonEncrypted192,
			},
			want:    []byte(commonSrc),
			wantErr: false,
		},
		{
			name:   "TestAesCbc_Decrypt-256",
			fields: fields{a: aescbc256},
			args: args{
				in: commonEncrypted256,
			},
			want:    []byte(commonSrc),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.a.Decrypt(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesCbc_Encrypt(t *testing.T) {
	type fields struct {
		a *AesCbc
	}
	type args struct {
		in []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:   "TestAesCbc_Encrypt-128",
			fields: fields{a: aescbc128},
			args: args{
				in: []byte(commonSrc),
			},
			want:    commonEncrypted128,
			wantErr: false,
		},
		{
			name:   "TestAesCbc_Encrypt-192",
			fields: fields{a: aescbc192},
			args: args{
				in: []byte(commonSrc),
			},
			want:    commonEncrypted192,
			wantErr: false,
		},
		{
			name:   "TestAesCbc_Encrypt-256",
			fields: fields{a: aescbc256},
			args: args{
				in: []byte(commonSrc),
			},
			want:    commonEncrypted256,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.a.Encrypt(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesCbc_decrypt(t *testing.T) {
	type fields struct {
		a *AesCbc
	}
	type args struct {
		in []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:   "TestAesCbc_decrypt-256",
			fields: fields{a: aescbc256},
			args: args{
				in: commonEncrypted256,
			},
			want:    []byte(commonSrc),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.a.decrypt(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesCbc_encrypt(t *testing.T) {
	type fields struct {
		a *AesCbc
	}
	type args struct {
		in []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:   "TestAesCbc_encrypt-256",
			fields: fields{a: aescbc256},
			args: args{
				in: []byte(commonSrc),
			},
			want:    commonEncrypted256,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.a.encrypt(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAesCbc(t *testing.T) {
	type args struct {
		key []byte
		iv  []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *AesCbc
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				key: []byte(commonKey256),
				iv:  []byte(commonIV),
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "short key",
			args: args{
				key: []byte("123"),
				iv:  []byte(commonIV),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "short iv",
			args: args{
				key: []byte(commonKey256),
				iv:  []byte("123"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewAesCbc(tt.args.key, tt.args.iv)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAesCbc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAesCbcDecrypt(t *testing.T) {
	type args struct {
		key []byte
		iv  []byte
		src []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "TestAesCbcDecrypt-128",
			args: args{
				key: []byte(commonKey128),
				iv:  []byte(commonIV),
				src: commonEncrypted128,
			},
			want:    []byte(commonSrc),
			wantErr: false,
		},
		{
			name: "TestAesCbcDecrypt-192",
			args: args{
				key: []byte(commonKey192),
				iv:  []byte(commonIV),
				src: commonEncrypted192,
			},
			want:    []byte(commonSrc),
			wantErr: false,
		},
		{
			name: "TestAesCbcDecrypt-256",
			args: args{
				key: []byte(commonKey256),
				iv:  []byte(commonIV),
				src: commonEncrypted256,
			},
			want:    []byte(commonSrc),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AesCbcDecrypt(tt.args.key, tt.args.iv, tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesCbcDecrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AesCbcDecrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesCbcDecryptBase64(t *testing.T) {
	type args struct {
		key []byte
		iv  []byte
		msg string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "TestAesCbcDecryptBase64-128",
			args: args{
				key: []byte(commonKey128),
				iv:  []byte(commonIV),
				msg: commonEncrypted128Base64,
			},
			want:    []byte(commonSrc),
			wantErr: false,
		},
		{
			name: "TestAesCbcDecryptBase64-192",
			args: args{
				key: []byte(commonKey192),
				iv:  []byte(commonIV),
				msg: commonEncrypted192Base64,
			},
			want:    []byte(commonSrc),
			wantErr: false,
		},
		{
			name: "TestAesCbcDecryptBase64-256",
			args: args{
				key: []byte(commonKey256),
				iv:  []byte(commonIV),
				msg: commonEncrypted256Base64,
			},
			want:    []byte(commonSrc),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AesCbcDecryptBase64(tt.args.key, tt.args.iv, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesCbcDecryptBase64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AesCbcDecryptBase64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesCbcDecryptHex(t *testing.T) {
	type args struct {
		key []byte
		iv  []byte
		msg string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "TestAesCbcDecryptHex-128",
			args: args{
				key: []byte(commonKey128),
				iv:  []byte(commonIV),
				msg: commonEncrypted128Hex,
			},
			want:    []byte(commonSrc),
			wantErr: false,
		},
		{
			name: "TestAesCbcDecryptHex-192",
			args: args{
				key: []byte(commonKey192),
				iv:  []byte(commonIV),
				msg: commonEncrypted192Hex,
			},
			want:    []byte(commonSrc),
			wantErr: false,
		},
		{
			name: "TestAesCbcDecryptHex-256",
			args: args{
				key: []byte(commonKey256),
				iv:  []byte(commonIV),
				msg: commonEncrypted256Hex,
			},
			want:    []byte(commonSrc),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AesCbcDecryptHex(tt.args.key, tt.args.iv, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesCbcDecryptHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AesCbcDecryptHex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesCbcEncrypt(t *testing.T) {
	type args struct {
		key []byte
		iv  []byte
		src []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "TestAesCbcEncrypt-128",
			args: args{
				key: []byte(commonKey128),
				iv:  []byte(commonIV),
				src: []byte(commonSrc),
			},
			want:    commonEncrypted128,
			wantErr: false,
		},
		{
			name: "TestAesCbcEncrypt-192",
			args: args{
				key: []byte(commonKey192),
				iv:  []byte(commonIV),
				src: []byte(commonSrc),
			},
			want:    commonEncrypted192,
			wantErr: false,
		},
		{
			name: "TestAesCbcEncrypt-256",
			args: args{
				key: []byte(commonKey256),
				iv:  []byte(commonIV),
				src: []byte(commonSrc),
			},
			want:    commonEncrypted256,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AesCbcEncrypt(tt.args.key, tt.args.iv, tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesCbcEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AesCbcEncrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesCbcEncryptBase64(t *testing.T) {
	type args struct {
		key []byte
		iv  []byte
		src []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "TestAesCbcEncryptBase64-128",
			args: args{
				key: []byte(commonKey128),
				iv:  []byte(commonIV),
				src: []byte(commonSrc),
			},
			want:    commonEncrypted128Base64,
			wantErr: false,
		},
		{
			name: "TestAesCbcEncryptBase64-192",
			args: args{
				key: []byte(commonKey192),
				iv:  []byte(commonIV),
				src: []byte(commonSrc),
			},
			want:    commonEncrypted192Base64,
			wantErr: false,
		},
		{
			name: "TestAesCbcEncryptBase64-256",
			args: args{
				key: []byte(commonKey256),
				iv:  []byte(commonIV),
				src: []byte(commonSrc),
			},
			want:    commonEncrypted256Base64,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AesCbcEncryptBase64(tt.args.key, tt.args.iv, tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesCbcEncryptBase64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AesCbcEncryptBase64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesCbcEncryptHex(t *testing.T) {
	type args struct {
		key []byte
		iv  []byte
		src []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "TestAesCbcEncryptHex-128",
			args: args{
				key: []byte(commonKey128),
				iv:  []byte(commonIV),
				src: []byte(commonSrc),
			},
			want:    commonEncrypted128Hex,
			wantErr: false,
		},
		{
			name: "TestAesCbcEncryptHex-192",
			args: args{
				key: []byte(commonKey192),
				iv:  []byte(commonIV),
				src: []byte(commonSrc),
			},
			want:    commonEncrypted192Hex,
			wantErr: false,
		},
		{
			name: "TestAesCbcEncryptHex-256",
			args: args{
				key: []byte(commonKey256),
				iv:  []byte(commonIV),
				src: []byte(commonSrc),
			},
			want:    commonEncrypted256Hex,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AesCbcEncryptHex(tt.args.key, tt.args.iv, tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesCbcEncryptHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AesCbcEncryptHex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcurrentAccess(t *testing.T) {
	w := sync.WaitGroup{}
	amount := 10000
	// first Encrypt
	go func() {
		for i := 0; i < amount; i++ {
			en1, err := aescbc256.Encrypt([]byte(commonSrc))
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(en1, commonEncrypted256) {
				t.Errorf("AesCbcDecrypt() got = %v, want %v", en1, commonEncrypted256)
			}
		}
		w.Done()
	}()

	// second Encrypt  test same encrypt together
	go func() {
		for i := 0; i < amount; i++ {
			en1, err := aescbc256.Encrypt([]byte(commonSrc))
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(en1, commonEncrypted256) {
				t.Errorf("AesCbcDecrypt() got = %v, want %v", en1, commonEncrypted256)
			}
		}
		w.Done()
	}()

	// third Encrypt test diff encrypt together
	go func() {
		for i := 0; i < amount; i++ {
			en1, err := aescbc256.Encrypt([]byte(commonSrc2))
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(en1, commonEncrypted256_2) {
				t.Errorf("AesCbcDecrypt() got = %v, want %v", en1, commonEncrypted256_2)
			}
		}
		w.Done()
	}()

	// first Decrypt test encrypt decrypt together
	go func() {
		for i := 0; i < amount; i++ {
			en1, err := aescbc256.Decrypt(commonEncrypted256)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(en1, []byte(commonSrc)) {
				t.Errorf("AesCbcDecrypt() got = %v, want %v", en1, commonSrc)
			}
		}
		w.Done()
	}()

	// second Decrypt test decrypt same together
	go func() {
		for i := 0; i < amount; i++ {
			en1, err := aescbc256.Decrypt(commonEncrypted256)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(en1, []byte(commonSrc)) {
				t.Errorf("AesCbcDecrypt() got = %v, want %v", en1, commonSrc)
			}
		}
		w.Done()
	}()

	// third Decrypt test decrypt diff together
	go func() {
		for i := 0; i < amount; i++ {
			en1, err := aescbc256.Decrypt(commonEncrypted256_2)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(en1, []byte(commonSrc2)) {
				t.Errorf("AesCbcDecrypt() got = %v, want %v", en1, commonSrc2)
			}
		}
		w.Done()
	}()

	w.Add(6)
	w.Wait()
}
