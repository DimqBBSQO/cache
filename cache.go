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
}

func New() *сache {
	c := &сache{
		dataBase: sync.Map{},
	}
	go c.ClearDataBase()
	return c
}

func (c *сache) ClearDataBase() {
	for {
		c.dataBase.Range(func(k, v interface{}) bool {
			val, _ := v.(value)
			if val.t.Add(val.dur).Before(time.Now()) {
				c.dataBase.Delete(k)
			}
			return true
		})
		time.Sleep(time.Second / 2)
	}
}

func (c *сache) Set(key string, v interface{}, ttl time.Duration) {
	c.dataBase.Store(key, value{v, time.Now(), ttl})
}

func (c *сache) Get(key string) (interface{}, bool) {
	return c.dataBase.Load(key)
}

func (c *сache) Delete(key string) {
	c.dataBase.Delete(key)
}
