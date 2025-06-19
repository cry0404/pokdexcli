package pokecache

import (

	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val  			[]byte //用字节切片来缓存我们当前使用的内容
}

type Cache struct {
	data 			map[string]cacheEntry
	mu 				sync.Mutex
	interval 	time.Duration
}
