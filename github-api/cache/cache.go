package cache

type Cache interface {
	Put(key string, val interface{}) error
	Get(key string) (interface{}, error)
}
