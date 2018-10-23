package internal

import (
	"sync"
	"sync/atomic"

	domain "github.com/kamilsk/guard/pkg/service/types"
)

// LicenseRequests is a blob of memory to store request counters.
var LicenseRequests = NewLicenseRequestCounter(maxLicenses)

// NewLicenseRequestCounter returns a blob of memory
// to store license request counters.
func NewLicenseRequestCounter(capacity int) interface {
	Increment(license domain.ID) uint32
	Rollback(license domain.ID)
} {
	return &lrCounter{
		mu:   &sync.RWMutex{},
		pool: make([]*uint32, 0, capacity),
		idx:  make(map[domain.ID]int, capacity),
	}
}

type lrCounter struct {
	mu   *sync.RWMutex
	pool []*uint32
	idx  map[domain.ID]int
}

// Increment increments request counter of the license
// and returns its new value.
func (c *lrCounter) Increment(license domain.ID) uint32 {
	c.mu.RLock()
	i, found := c.idx[license]
	c.mu.RUnlock()

	if !found {
		i = c.init(license)
	}

	return atomic.AddUint32(c.pool[i], 1)
}

// Rollback decrements request counter of the license by 1.
// It must be called after Increment. For example:
//
//     if limit < counter.Increment(license) {
//         go counter.Rollback(license)
//         return errors.New("limit exceeded")
//     }
//
// Otherwise panic occurs.
func (c *lrCounter) Rollback(license domain.ID) {
	c.mu.RLock()
	i := c.idx[license]
	c.mu.RUnlock()

	atomic.AddUint32(c.pool[i], ^uint32(0))
}

func (c *lrCounter) init(license domain.ID) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	i, found := c.idx[license]
	if found {
		return i
	}

	if len(c.pool)+1 > cap(c.pool) {
		panic("segmentation fault: increase license request counter capacity")
	}

	i, counter := len(c.pool), new(uint32)
	c.idx[license] = i
	c.pool = append(c.pool, counter)
	return i
}
