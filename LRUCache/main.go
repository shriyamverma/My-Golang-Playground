package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type CacheEntry struct {
	Key       string
	Value     interface{}
	ExpiredAt time.Time
}

type LRUCache struct {
	storage   map[string]*list.Element // Store pointers to the linked list nodes
	evictList *list.List               // Tracks the LRU order (Front = newest, Back = oldest)
	capacity  int64                    // Maximum allowed items
	mu        sync.Mutex               // Use a standard Mutex to handle lookups + inline deletions safely
}

func NewLRUCache(capacity int64) *LRUCache {
	return &LRUCache{
		storage:   make(map[string]*list.Element),
		evictList: list.New(),
		capacity:  capacity,
	}
}

func (c *LRUCache) Put(key string, value interface{}, ttlSeconds int64) {
	c.mu.Lock()
	defer c.mu.Unlock()

	expiredAt := time.Now().Add(time.Duration(ttlSeconds) * time.Second)

	// Case 1: Key already exists. Update value/TTL and mark it as most recently used.
	if elem, exists := c.storage[key]; exists {
		c.evictList.MoveToFront(elem)
		entry := elem.Value.(*CacheEntry)
		entry.Value = value
		entry.ExpiredAt = expiredAt
		return
	}

	// Case 2: New key. Check if we hit the capacity limit.
	if int64(c.evictList.Len()) >= c.capacity {
		// Evict the oldest item from the back of the list
		oldest := c.evictList.Back()
		if oldest != nil {
			c.evictList.Remove(oldest)
			oldestEntry := oldest.Value.(*CacheEntry)
			delete(c.storage, oldestEntry.Key)
		}
	}

	// Insert the new item at the front of the list and save it to map
	entry := &CacheEntry{
		Key:       key,
		Value:     value,
		ExpiredAt: expiredAt,
	}
	elem := c.evictList.PushFront(entry)
	c.storage[key] = elem
}

func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.mu.Lock() // Using a full Lock because a read might trigger an eviction/delete
	defer c.mu.Unlock()

	elem, exists := c.storage[key]
	if !exists {
		return nil, false
	}

	entry := elem.Value.(*CacheEntry)

	// Check if the item has expired
	if time.Now().After(entry.ExpiredAt) {
		// Lazy eviction: remove it right now
		c.evictList.Remove(elem)
		delete(c.storage, key)
		return nil, false
	}

	// Item is valid: refresh its positioning as Most Recently Used
	c.evictList.MoveToFront(elem)
	return entry.Value, true
}

func main() {
	// Initialize cache with a capacity of 2 items
	cache := NewLRUCache(2)

	// Add two items
	cache.Put("A", "Apple", 5)
	cache.Put("B", "Banana", 5)

	// Access "A" to make it the most recently used item
	cache.Get("A")

	// Add a 3rd item. This will evict "B" because "A" was just refreshed!
	cache.Put("C", "Cherry", 5)

	valA, foundA := cache.Get("A")
	valB, foundB := cache.Get("B")
	valC, foundC := cache.Get("C")

	fmt.Printf("A: %v (Found: %v)\n", valA, foundA) // A: Apple (Found: true)
	fmt.Printf("B: %v (Found: %v)\n", valB, foundB) // B: <nil> (Found: false) -> Evicted by LRU
	fmt.Printf("C: %v (Found: %v)\n", valC, foundC) // C: Cherry (Found: true)

	// Wait for TTL to expire
	fmt.Println("\nWaiting 6 seconds for TTL expiration...")
	time.Sleep(6 * time.Second)

	valAAfter, foundAAfter := cache.Get("A")
	fmt.Printf("A after TTL: %v (Found: %v)\n", valAAfter, foundAAfter) // A after TTL: <nil> (Found: false)
}
