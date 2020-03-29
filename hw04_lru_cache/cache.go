package hw04_lru_cache //nolint:golint,stylecheck
import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool // Добавить значение в кэш по ключу
	Get(key Key) (interface{}, bool)     // Получить значение из кэша по ключу
	Clear()                              // Очистить кэш
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*listItem
	mut      sync.Mutex
}

type cacheItem struct {
	key Key
	val interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mut.Lock()
	defer c.mut.Unlock()
	val, ok := c.items[key]
	if ok {
		c.queue.MoveToFront(val)
		val.Value = cacheItem{key: key, val: value}
		return true
	}
	it := cacheItem{key: key, val: value}
	c.queue.PushFront(it)
	c.items[key] = c.queue.Front()

	//check if capacity is overflow
	if len(c.items) > c.capacity {
		tmp := c.queue.Back()
		c.queue.Remove(tmp)
		delete(c.items, tmp.Value.(cacheItem).key)
	}

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mut.Lock()
	defer c.mut.Unlock()
	val, ok := c.items[key]
	if !ok {
		return nil, false
	}
	c.queue.MoveToFront(val)
	return val.Value.(cacheItem).val, true
}

func (c *lruCache) Clear() {
	c.mut.Lock()
	defer c.mut.Unlock()
	backItem := c.queue.Back()
	for {
		delete(c.items, backItem.Value.(cacheItem).key)
		backItem = backItem.Next
		if backItem == nil {
			break
		}
	}
	c.queue = NewList()
}

func NewCache(capacity int) Cache {
	cache := &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*listItem),
	}
	return cache
}
