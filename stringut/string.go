package stringut

import (
	"strings"
	"unsafe"
)

func StrAppend(sep string, vs ...string) string {
	var builder strings.Builder
	for i, v := range vs {
		builder.WriteString(v)
		if i < len(vs)-1 {
			builder.WriteString(sep)
		}
	}

	return builder.String()
}

func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}

	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func AnyIsBlack(vs ...string) bool {
	for _, v := range vs {
		if strings.TrimSpace(v) == "" {
			return true
		}
	}
	return false
}
