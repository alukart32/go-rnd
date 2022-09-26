package prototype

import (
	"errors"
)

const (
	White = 1
	Black = 2
	Blue  = 3
)

var cache *ShirtsCache

func NewShirtsCache() *ShirtsCache {
	if cache == nil {
		cache = new(ShirtsCache)
		cache.elems = make(map[int]Prototype, 3)
		cache.elems[White] = whitePrototype
		cache.elems[Black] = blackPrototype
		cache.elems[Blue] = bluePrototype
	}
	return cache
}

type ShirtsCache struct {
	elems map[int]Prototype
}

func (sc *ShirtsCache) GetClone(m int) (Prototype, error) {
	switch m {
	case White, Black, Blue:
		return sc.elems[m].Clone(), nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}
