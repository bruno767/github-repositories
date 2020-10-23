package cache

import "time"

type entry struct {
	val     interface{}
	expires time.Time
}
