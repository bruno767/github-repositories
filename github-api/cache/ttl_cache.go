package cache

import (
	"github.com/hashicorp/golang-lru"
	"log"
	"time"
)

func NewTTLCache(ttl time.Duration, size int) *TtlCache {
	v, err := lru.NewARC(size)
	if err != nil {
		log.Fatal(err)
	}
	t := TtlCache{
		ttl:   ttl,
		cache: v,
	}
	return &t
}

type TtlCache struct {
	ttl   time.Duration
	cache *lru.ARCCache
}

func (r *TtlCache) Put(key string, val interface{}) error {
	r.cache.Add(key, &entry{val: val, expires: time.Now().Add(r.ttl)})
	return nil
}

func (r *TtlCache) Get(key string) (interface{}, error) {
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
