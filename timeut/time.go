package timeut

import "time"

const timeFmt = "2006-01-02 15:04:05"

func TickMil(t ...time.Time) int64 {
	if len(t) == 0 {
		return time.Now().Unix() * 1e3
	}
	return t[0].Unix()* 1e3
}

func TickSec(t ...time.Time) int64 {
	if len(t) == 0 {
		return time.Now().Unix()
	}
	return t[0].Unix()
}

func Zero() time.Time {
	year, month, day := time.Now().Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func ZeroTickMil() int64 {
	return TickMil(Zero())
}

func ZeroTickSec() int64 {
	return TickSec(Zero())
}
