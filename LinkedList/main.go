package main

import (
	"LinkedList/cache"
	"fmt"
)

func main() {
	cache := cache.CreateCache(3)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	value, error := cache.Get(3)
	cache.Put(4, 4)
	cache.Put(5, 5)
	value, error = cache.Get(4)
	cache.Put(6, 6)

	if error == nil {
		fmt.Println(value)
	} else {
		fmt.Printf(error.Error())
	}
	cache.PrintCache()

}
