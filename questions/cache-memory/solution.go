package main

import (
	"errors"
	"fmt"
)

type CacheMemory struct {
	capacity int
	store    map[string]string
	keys     []string
}

func NewCacheMemory(capacity int) *CacheMemory {
	return &CacheMemory{
		capacity: capacity,
		store:    make(map[string]string),
		keys:     []string{},
	}
}

func (c *CacheMemory) Set(key, value string) {
	if _, ok := c.store[key]; ok {
		c.store[key] = value
		idx := -1
		for i, k := range c.keys {
			if k == key {
				idx = i
				break
			}
		}
		if idx != -1 {
			c.keys = append(c.keys[:idx], c.keys[idx+1:]...)
			c.keys = append(c.keys, key)
		}
		return
	}
	if len(c.store) < c.capacity {
		c.store[key] = value
		c.keys = append(c.keys, key)
	} else {
		delete(c.store, c.keys[0])
		c.store[key] = value
		c.keys = append(c.keys[1:], key)
	}
}

func (c *CacheMemory) Get(key string) (string, error) {
	if value, ok := c.store[key]; ok {
		return value, nil
	}
	return "", errors.New("key not found")
}

func main() {
	cache := NewCacheMemory(4)
	cache.Set("key-1", "value-1")
	cache.Set("key-2", "value-2")
	cache.Set("key-3", "value-3")
	cache.Set("key-4", "value-4")
	cache.Set("key-5", "value-5")
	cache.Set("key-6", "value-6")
	cache.Set("key-7", "value-7")
	cache.Set("key-8", "value-8")
	cache.Set("key-9", "value-9")
	cache.Set("key-10", "value-10")
	cache.Set("key-11", "value-11")
	fmt.Println(cache.Get("key-1"))
	fmt.Println(cache.Get("key-2"))
	fmt.Println(cache.Get("key-3"))
	fmt.Println(cache.Get("key-4"))
	fmt.Println(cache.Get("key-5"))
	fmt.Println(cache.Get("key-6"))
	fmt.Println(cache.Get("key-7"))
	fmt.Println(cache.Get("key-8"))
	fmt.Println(cache.Get("key-9"))
	fmt.Println(cache.Get("key-10"))
	fmt.Println(cache.Get("key-11"))
}
