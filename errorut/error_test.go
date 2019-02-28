package errorut

import (
	"errors"
	"testing"
)

func TestIsErrByName(t *testing.T) {
	errName := "something is wrong"
	newErr := New(errName, "newErr")

	type args struct {
		err  error
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				err:  newErr,
				name: errName,
			},
			want: true,
		},

		{
			args: args{
				err:  newErr,
				name: "what ever",
			},
			want: false,
		},

		{
			args: args{
				err:  errors.New("a"),
				name: errName,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsErrByName(tt.args.err, tt.args.name); got != tt.want {
				t.Errorf("IsErrByName() = %v, want %v", got, tt.want)
			}
		})
	}
}
