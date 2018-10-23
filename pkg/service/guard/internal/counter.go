package internal

import (
	"sync"
	"sync/atomic"

	domain "github.com/kamilsk/guard/pkg/service/types"
)

// LicenseRequests is blob of memory to store request counters.
var LicenseRequests = lrCounter{
	mu:   &sync.RWMutex{},
	pool: make([]*uint32, 0, maxLicenses),
	idx:  make(map[domain.ID]int, maxLicenses),
}

type lrCounter struct {
	mu   *sync.RWMutex
	pool []*uint32
	idx  map[domain.ID]int
}

// IncrementFor increments request counter
// of the license and returns its new value.
func (c *lrCounter) IncrementFor(license domain.ID) uint32 {
	c.mu.RLock()
	i, found := c.idx[license]
	c.mu.RUnlock()

	if !found {
		i = c.init(license)
	}

	return atomic.AddUint32(c.pool[i], 1)
}

func (c *lrCounter) init(license domain.ID) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	i, found := c.idx[license]
	if found {
		return i
	}

	if len(c.pool)+1 > maxLicenses {
		panic("segmentation fault: increase maxLicenses const")
	}

	i, counter := len(c.pool), new(uint32)
	c.idx[license] = i
	c.pool = append(c.pool, counter)
	return i
}
