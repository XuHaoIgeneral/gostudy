package main

import (
	"sync"
	"fmt"
)

type SafeMap struct {
	sync.RWMutex
	dict map[int]int
}

func newSafeMap() *SafeMap {
	sm := new(SafeMap)
	sm.dict = make(map[int]int)
	return sm
}

func main() {
	dict := newSafeMap()

	for i := 0; i < 100; i++ {

		go dict.writeMap(i, i)

		go dict.readMap(i)

	}
}

func (sm *SafeMap) readMap(key int) int {
	sm.RLock()
	defer sm.RUnlock()
	value := sm.dict[key]
	fmt.Println("读取", value)
	return value
}

func (sm *SafeMap) writeMap(key, value int) {
	//fmt.Println("writeMap")
	sm.Lock()
	defer sm.Unlock()
	sm.dict[key] = value
	fmt.Println("写入", key)
}
