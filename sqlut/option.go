package sqlut

type Options struct {
	RawSQL     string
	RawValue   []interface{}
	WhereQuery string
	WhereArgs  []interface{}
}

type SetOpt func(opts *Options)

func RawSQL(sql string, value ...interface{}) SetOpt {
	return func(opts *Options) {
		opts.RawSQL = sql
		opts.RawValue = value
	}
}

func WhereQuery(value string) SetOpt {
	return func(opts *Options) {
		opts.WhereQuery = value
	}
}

func WhereArgs(value interface{}) SetOpt {
	return func(opts *Options) {
		opts.WhereArgs = append(opts.WhereArgs, value)
	}
}

func (o *Options) IsRawQuery() bool {
	return o.RawSQL != ""
}

func (o *Options) IsORMQuery() bool {
	return o.RawSQL == "" && o.WhereQuery != ""
}
