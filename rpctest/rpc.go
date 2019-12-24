package rpctest

type Scheduler struct {
	before []func()
	after  []func()

	runServer  func()
	stopServer func()
	clientCall func()

	doneC chan struct{}
}

func New(runServer, stopServer, clientCall func(), setOpts ...SetOpt) *Scheduler {
	opt := &option{}
	for _, setOpt := range setOpts {
		setOpt(opt)
	}

	return &Scheduler{
		doneC: make(chan struct{}),

		runServer:  runServer,
		stopServer: stopServer,
		clientCall: clientCall,

		before: opt.before,
		after:  opt.after,
	}
}

func (s *Scheduler) Run() {
	defer s.done()

	for _, call := range s.before {
		call()
	}

	defer func() {
		for _, call := range s.after {
			call()
		}
	}()

	go s.runServer()

	go func() {
		<-s.doneC
		s.stopServer()
	}()

	s.clientCall()
}

func (s *Scheduler) done() {
	close(s.doneC)
}
