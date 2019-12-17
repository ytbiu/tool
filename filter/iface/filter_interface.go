package iface

// 过滤器
type Filter interface {
	Name() string
	Check(interface{}) bool
	Node
}

type Node interface {
	Next() Filter
	SetNext(Filter)
}
