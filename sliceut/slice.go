package sliceut

import (
	"reflect"
)

func Contains(src interface{}, target interface{}) bool {
	infSlice := toInfSlice(src)
	if len(infSlice) == 0 {
		return false
	}

	return contains(infSlice, target)
}

func contains(src []interface{}, target interface{}) bool {
	for _, e := range src {
		if e == target {
			return true
		}
	}
	return false
}

func Remove(src interface{}, target interface{}) []interface{} {
	infSlice := toInfSlice(src)
	for i, e := range infSlice {
		if e == target {
			return append(infSlice[:i], infSlice[i+1:]...)
		}
	}
	return infSlice
}

func toInfSlice(src interface{}) []interface{} {
	v := reflect.ValueOf(src)
	if v.Kind() != reflect.Slice {
		return nil
	}
	l := v.Len()
	infSlice := make([]interface{}, l)
	for i := 0; i < l; i++ {
		infSlice[i] = v.Index(i).Interface()
	}
	return infSlice
}
