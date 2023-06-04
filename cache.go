package cache

import (
	"sync"
	"time"
)

type value struct {
	v   interface{}
	t   time.Time
	dur time.Duration
}

type сache struct {
	dataBase sync.Map
	ticker   time.Ticker
}

func New() *сache {
	c := &сache{
		dataBase: sync.Map{},
		ticker:   *time.NewTicker(time.Second),
	}
	go c.ClearDataBase()
	return c
}

func (c *сache) ClearDataBase() {
	for {
		<-c.ticker.C
		c.dataBase.Range(func(k, v interface{}) bool {
			val, _ := v.(value)
			if val.t.Add(val.dur).Before(time.Now()) {
				c.dataBase.Delete(k)
			}
			return true
		})
	}
}

func (c *сache) Set(key string, v interface{}, ttl time.Duration) {
	c.dataBase.Store(key, value{v, time.Now(), ttl})
}

func (c *сache) Get(key string) interface{} {
	v, _ := c.dataBase.Load(key)
	val := v.(value)
	return val.v
}

func (c *сache) Delete(key string) {
	c.dataBase.Delete(key)
}
