package cryptout

import (
	"reflect"
	"testing"
)

func TestXORDecrypt(t *testing.T) {
	xorKey := "1234"
	data := []byte("13579abc")
	oriData := make([]byte, len(data))
	copy(oriData, data)
	encData := XOREncrypt(data, []byte(xorKey))

	type args struct {
		value string
		key   []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			args: args{
				value: encData,
				key:   []byte(xorKey),
			},
			want:    oriData,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := XORDecrypt(tt.args.value, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("XORDecrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XORDecrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
