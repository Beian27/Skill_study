package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	// W
	m.Store("zhy", 15)
	m.Store("mmy", 16)

	// R
	load, ok := m.Load("zhy")
	if ok {
		fmt.Println(load.(int))
	}

	// range
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key.(string), value.(int))
		return true
	})

	m.Delete("zhy")
	m.LoadOrStore("zhy", 1000)
	load1, ok := m.Load("zhy")
	if ok {
		fmt.Println(load1.(int))
	}
}
