package cache

import "fmt"

type Cache struct {
	cache map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]interface{}),
	}
}

func (c Cache) Set(key string, value interface{}) {
	_, ok := c.cache[key]
	if !ok {
		c.cache[key] = value
		fmt.Println("К кеш успешно записано значение ", key, " - ", value)
	} else {
		fmt.Println("В кэше уже присутствует ключ - ", key)
	}
}

func (c Cache) Get(key string) interface{} {
	value, ok := c.cache[key]
	if ok {
		return value
	}
	return nil
}

func (c Cache) Delete(key string) {
	_, ok := c.cache[key]
	if ok {
		delete(c.cache, key)
		fmt.Println("Ключ ", key, " успешно удален!")
	} else {
		fmt.Println("Ключ ", key, " не найден!")
	}
}
