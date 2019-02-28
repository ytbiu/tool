package sliceut

import (
	"reflect"
	"testing"
)

func Test_Contains(t *testing.T) {
	type args struct {
		src    interface{}
		target interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				src:    []int{1, 2, 3},
				target: 1,
			},
			want: true,
		},
		{
			args: args{
				src:    []int{},
				target: 1,
			},
			want: false,
		},

		{
			args: args{
				src:    []string{"a", "b", "c"},
				target: "d",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if find := Contains(tt.args.src, tt.args.target); find != tt.want {
				t.Errorf("Contains() = %v, want %v", find, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	type args struct {
		src    interface{}
		target interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			args: args{
				src:    []int{1, 2, 3},
				target: 1,
			},
			want: []interface{}{2, 3},
		},
		{
			args: args{
				src:    []int{1, 2, 3},
				target: 4,
			},
			want: []interface{}{1, 2, 3},
		},
		{
			args: args{
				src:    []string{"a", "b", "c"},
				target: "b",
			},
			want: []interface{}{"a", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Remove(tt.args.src, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			} else {
				t.Logf("%v", got)
			}
		})
	}
}
