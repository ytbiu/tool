package filter

import "github.com/ytbiu/tool/filter/iface"

// 责任链实现(链表)
type LinkChain struct {
	size  int
	first iface.Filter
}

func NewLinkChain() iface.Chain {
	return &LinkChain{}
}

// 追加过滤器
func (lc *LinkChain) Append(filters ...iface.Filter) {
	for _, filter := range filters {
		lc.add(filter)
	}
}

// 删除过滤器
func (lc *LinkChain) Remove(filter iface.Filter) {
	if lc.first == nil {
		return
	}

	defer func() {
		lc.size--
	}()
	currentFilter := lc.first
	if currentFilter.Name() == filter.Name() {
		lc.first = nil
		return
	}

	for i := 0; i < lc.size-1; i++ {
		if currentFilter.Next().Name() != filter.Name() {
			continue
		}

		currentFilter.SetNext(currentFilter.Next().Next())
	}

}

// 执行过滤方法
func (lc *LinkChain) Check(content interface{}) bool {
	return lc.check(lc.first, content)
}

func (lc *LinkChain) check(f iface.Filter, content interface{}) bool {
	if f == nil {
		return true
	}

	if !f.Check(content) {
		return false
	}

	return lc.check(f.Next(), content)
}

func (lc *LinkChain) add (filter iface.Filter) {
	if filter == nil {
		return
	}

	defer func() {
		lc.size++
	}()

	if lc.first == nil {
		lc.first = filter
		return
	}

	currentFilter := lc.first
	for i := 0; i < lc.size; i++ {
		if currentFilter.Next() != nil {
			currentFilter = currentFilter.Next()
			continue
		}

		currentFilter.SetNext(filter)
	}
}