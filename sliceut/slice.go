package sliceut

import (
	"reflect"
	"sync"
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

func isSame(src, dst interface{}) bool {

	srcSlice := toInfSlice(src)
	dstSlice := toInfSlice(dst)

	if len(srcSlice) != len(dstSlice) {
		return false
	}

	dataC := make(chan interface{}, 1)
	resC := make(chan bool, 1)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, e := range srcSlice {
			dataC <- e
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, e := range dstSlice {
			if <-dataC != e {
				resC <- false
			}
		}
	}()

	go func() {
		wg.Wait()
		close(resC)
	}()

	for res := range resC {
		return res
	}

	return true
}
