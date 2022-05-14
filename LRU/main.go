package main

import "fmt"

type LRUCache struct {
	capacity int
	index    map[int]*node
	head     *node
}

type node struct {
	key  int
	val  int
	pre  *node
	next *node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{capacity: capacity}
	cache.index = make(map[int]*node)

	return cache
}

func (lru *LRUCache) Get(key int) int {
	if ptr, ok := lru.index[key]; ok {
		if ptr == lru.head {
			return ptr.val
		}

		lru.PutToFirst(ptr)

		return ptr.val
	} else {
		return -1
	}
}

func (lru *LRUCache) Put(key, val int) {
	// 存在则更新，放到第一位去
	if ptr, ok := lru.index[key]; ok {
		ptr.val = val

		if ptr != lru.head {
			lru.PutToFirst(ptr)
		}

		return
	}

	// 不存在，新建一个放第一位
	ptr := &node{key: key, val: val}
	if lru.head == nil {
		lru.head = ptr
		ptr.pre = ptr
		ptr.next = ptr
	} else {
		ptr.next = lru.head
		ptr.pre = lru.head.pre
		lru.head.pre.next = ptr
		lru.head.pre = ptr
		lru.head = ptr
	}

	// 容量不够，删掉最后一个
	if len(lru.index) >= lru.capacity {
		end := lru.head.pre
		delete(lru.index, end.key)
		end.pre.next = lru.head
		lru.head.pre = end.pre
	}

	lru.index[key] = ptr
	return
}

func (lru *LRUCache) PutToFirst(ptr *node) {
	ptr.pre.next = ptr.next
	ptr.next.pre = ptr.pre

	ptr.next = lru.head
	ptr.pre = lru.head.pre

	lru.head.pre.next = ptr
	lru.head.pre = ptr

	lru.head = ptr
}

func main() {
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3)
	cache.Put(4, 1)
	val := cache.Get(1)
	fmt.Println(val)
	val = cache.Get(2)
	fmt.Println(val)
}
