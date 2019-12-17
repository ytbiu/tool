package filter

import "github.com/ytbiu/tool/filter/iface"

type LinkFilter struct {
	next iface.Filter

	iface.Filter
}

func NewLinkFilter(filter iface.Filter) iface.Filter {
	return &LinkFilter{
		Filter: filter,
	}
}

func (ln *LinkFilter) Next() iface.Filter {
	return ln.next
}

func (ln *LinkFilter) SetNext(f iface.Filter) {
	ln.next = f
}
