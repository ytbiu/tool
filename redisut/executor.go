package redisut

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

type PoolOpt struct {
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
	Wait        bool
	Addr        string
	DBNo        int
	PassWord    string
}

func NewPool(opt *PoolOpt) *redis.Pool {
	p := &redis.Pool{
		MaxIdle:     opt.MaxIdle,
		MaxActive:   opt.MaxActive,
		IdleTimeout: opt.IdleTimeout,
		Wait:        opt.Wait,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", opt.Addr)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("auth", opt.PassWord); err != nil {
				c.Close()
				return nil, err
			}
			if _, err := c.Do("select", opt.DBNo); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Now().After(t.Add(time.Second * 10)) {
				_, err := c.Do("ping")
				return err
			}
			return nil
		},
	}
	return p
}

type Executor interface {
	Do(string, ...interface{}) (interface{}, error)
	Set(k, v interface{}) error
	SetEX(k, v interface{}, ex int) error
	Get(k interface{}) (interface{}, error)
	GetInt(k interface{}) (int, error)
	GetString(k interface{}) (string, error)
	Ttl(k interface{}) (int, error)
}

type executor struct {
	*redis.Pool
}

func (t *executor) Do(cmd string, args ...interface{}) (interface{}, error) {
	c := t.Pool.Get()
	defer t.Close()
	return c.Do(cmd, args...)
}

func (t *executor) Set(k, v interface{}) error {
	_, err := t.Do("SET", k, v)
	return err
}

func (t *executor) SetEX(k, v interface{}, ex int) error {
	_, err := t.Do("SETEX", k, ex, v)
	return err
}

func (t *executor) Get(k interface{}) (interface{}, error) {
	return t.Do("GET", k)
}

func (t *executor) GetInt(k interface{}) (int, error) {
	return redis.Int(t.Get(k))
}

func (t *executor) GetString(k interface{}) (string, error) {
	return redis.String(t.Get(k))
}

func (t *executor) Ttl(k interface{}) (int, error) {
	return redis.Int(t.Do("TTL", k))
}

func (t *executor) Del(k interface{}) error {
	_, err := t.Do("DEL", k)
	return err
}

func NewExecutor(p *redis.Pool) Executor {
	return &executor{Pool: p}
}

type Queue interface {
}

//func Do() func (string, ...interface{}) (interface{}, error) {
//
//}
