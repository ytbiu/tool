package timeut

import (
	"testing"
	"time"
)

func TestTickMil(t *testing.T) {
	y,m,d:=time.Now().Date()
	tt := time.Date(y,m,d,12,0,0,0,time.Local)

	type args struct {
		t []time.Time
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{},
		},

		{
			args: args{
				t: []time.Time{tt},
			},
		},

	}
	for _, tt := range tests {
		if ttt := tt.args.t;len(ttt)==0{
			if TickMil(ttt...) != time.Now().Unix()*1e3{
				t.Errorf("TickMil():%d - now:%d",TickMil(tt.args.t...),time.Now().Unix()*1e3)
			}
			t.Logf("TickMil():%d - now:%d",TickMil(tt.args.t...),time.Now().Unix()*1e3)
		}else {
			if TickMil(ttt...) != ttt[0].Unix()*1e3{
				t.Errorf("TickMil():%d - the time tick:%d",TickMil(tt.args.t...),ttt[0].Unix()*1e3)
			}
		}
	}
}

func TestTickSec(t *testing.T) {
	y,m,d:=time.Now().Date()
	tt := time.Date(y,m,d,12,0,0,0,time.Local)

	type args struct {
		t []time.Time
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{},
		},

		{
			args: args{
				t: []time.Time{tt},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ttt := tt.args.t;len(ttt)==0{
				if TickSec(ttt...) != time.Now().Unix(){
					t.Errorf("TickMil():%d - now:%d",TickSec(tt.args.t...),time.Now().Unix())
				}
				t.Logf("TickMil():%d - now:%d",TickSec(tt.args.t...),time.Now().Unix())
			}else {
				if TickSec(ttt...) != ttt[0].Unix(){
					t.Errorf("TickMil():%d - the time tick:%d",TickSec(tt.args.t...),ttt[0].Unix())
				}
			}
		})
	}
}

func TestZero(t *testing.T) {
	tests := []struct {
		name string
		want time.Time
	}{
	{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%v",Zero())
		})
	}
}

func TestZeroTickMil(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
	{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%v", ZeroTickMil())
		})
	}
}

func TestZeroTickSec(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
	{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				t.Logf("%v", ZeroTickSec())
			})
		})
	}
}
