package cache

import (
	"log"
	"time"
	"github.com/hashicorp/golang-lru"
)

func NewTTLCache(ttl time.Duration, size int) *ttlCache {
	v, err := lru.NewARC(size)
	if err != nil {
		log.Fatal(err)
	}
	t := ttlCache{
		ttl:   ttl,
		cache: v,
	}
	return &t
}

type ttlCache struct {
	ttl   time.Duration
	cache *lru.ARCCache
}

func (r *ttlCache) Put(key string, val interface{}) error {
	r.cache.Add(key, &entry{val: val, expires: time.Now().Add(r.ttl)})
	return nil
}

func (r *ttlCache) Get(key string) (interface{}, error) {
	v, found := r.cache.Get(key)
	if found {
		e := v.(*entry)
		if e.expires.After(time.Now()) {
			return e.val, nil
		}
		r.cache.Remove(key)
	}
	return nil, ErrorNotFound
}
