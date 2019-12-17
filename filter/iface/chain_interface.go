package iface

// 过滤器责任链
type Chain interface {
	Append(...Filter)
	Remove(Filter)
	Check(interface{}) bool
}
