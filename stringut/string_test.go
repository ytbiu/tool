package stringut

import (
	"reflect"
	"testing"
)

func TestStrAppend(t *testing.T) {
	type args struct {
		sep string
		vs  []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				sep: "-",
				vs:  []string{"1","2","3"},
			},
			want: "1-2-3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrAppend(tt.args.sep, tt.args.vs...); got != tt.want {
				t.Errorf("StrAppend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStr2bytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{
				s: "12345",
			},
			want: []byte("12345"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Str2bytes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Str2bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes2str(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				b: []byte("12345"),
			},
			want: "12345",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bytes2str(tt.args.b); got != tt.want {
				t.Errorf("Bytes2str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnyIsBlack(t *testing.T) {
	type args struct {
		vs []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				vs: []string{"1","a",""},
			},
			want: true,
		},

		{
			args: args{
				vs: []string{"1","a","b"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnyIsBlack(tt.args.vs...); got != tt.want {
				t.Errorf("AnyIsBlack() = %v, want %v", got, tt.want)
			}
		})
	}
}
