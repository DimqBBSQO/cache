package main

import (
	"fmt"
	"golang-ninja/basic/cache"
)

func main() {
	cacheNew := cache.NewCache()
	cacheNew.Set("userId", 42)
	userId := cacheNew.Get("userId")
	fmt.Println(userId)
	cacheNew.Delete("userId")
	userId = cacheNew.Get("userId")
	fmt.Println(userId)

}
