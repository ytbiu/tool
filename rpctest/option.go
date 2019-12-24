package rpctest

type option struct {
	before []func()
	after  []func()
}

type SetOpt func(opt *option)

func WithBefore(before ...func()) SetOpt {
	return func(opt *option) {
		opt.before = before
	}
}

func WithAfter(after ...func()) SetOpt {
	return func(opt *option) {
		opt.after = after
	}
}
