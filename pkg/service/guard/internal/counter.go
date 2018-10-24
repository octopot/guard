package internal

import (
	"sort"
	"sync"
	"sync/atomic"
	"time"

	domain "github.com/kamilsk/guard/pkg/service/types"
)

// TODO issue#draft {

var (
	// LicenseRequests is a blob of memory to store request counters.
	LicenseRequests = NewLicenseRequestCounter(maxLicenses)
	// LicenseWorkplaces is a blob of memory to store workplace counters.
	LicenseWorkplaces = NewLicenseWorkplaceCounter(maxLicenses, maxWorkplaces)
)

// NewLicenseRequestCounter returns a blob of memory
// to store license request counters.
func NewLicenseRequestCounter(capacity int) interface {
	Increment(license domain.ID) uint32
	Rollback(license domain.ID)
} {
	return &lrCounter{
		pool: make([]uint32, 0, capacity),
		idx:  make(map[domain.ID]int, capacity),
	}
}

type lrCounter struct {
	mu   sync.RWMutex
	pool []uint32
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

	return atomic.AddUint32(&c.pool[i], 1)
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

	atomic.AddUint32(&c.pool[i], ^uint32(0))
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

	i = len(c.pool)
	c.idx[license] = i
	c.pool = append(c.pool, 0)
	return i
}

// NewLicenseWorkplaceCounter returns a blob of memory
// to store license workplace counters.
func NewLicenseWorkplaceCounter(capacity, wpCapacity int) interface {
	Acquire(license, workplace domain.ID, capacity int) bool
} {
	return &lwCounter{
		pool: make([]slot, 0, capacity),
		idx:  make(map[domain.ID]int, capacity),
		wpc:  wpCapacity,
	}
}

type lwCounter struct {
	mu   sync.RWMutex
	pool []slot
	idx  map[domain.ID]int
	wpc  int
}

// Acquire tries to hold available workplace slot.
func (c *lwCounter) Acquire(license, workplace domain.ID, capacity int) bool {
	if capacity == 0 {
		return true
	}

	c.mu.RLock()
	i, found := c.idx[license]
	c.mu.RUnlock()

	if !found {
		i = c.init(license)
	}

	return c.pool[i].acquire(workplace, capacity)
}

func (c *lwCounter) init(license domain.ID) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	i, found := c.idx[license]
	if found {
		return i
	}

	if len(c.pool)+1 > cap(c.pool) {
		panic("segmentation fault: increase license workplace counter capacity")
	}

	i = len(c.pool)
	c.idx[license] = i
	c.pool = append(c.pool, slot{
		pool: make([]record, 0, c.wpc),
		idx:  make(map[domain.ID]int, c.wpc),
		wpc:  c.wpc,
	})
	return i
}

type slot struct {
	mu   sync.Mutex
	pool []record
	idx  map[domain.ID]int
	wpc  int // to check assertion fast without lock to prevent data race
}

func (s *slot) acquire(workplace domain.ID, capacity int) bool {
	// assert(capacity <= cap(s.pool))
	if capacity > s.wpc {
		panic("segmentation fault: increase license workplace slot capacity")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.pool) > capacity {
		s.shrink(capacity)
	}

	i, found := s.idx[workplace]
	if found {
		s.pool[i].touch()
		return true
	}

	if i = len(s.pool); i < capacity {
		s.idx[workplace] = i
		s.pool = append(s.pool, record{workplace: workplace, lastActive: time.Now()})
		return true
	}

	// try to displace
	now := time.Now()
	for i = range s.pool {
		if now.Sub(s.pool[i].lastActive) > workplaceTTL {
			delete(s.idx, s.pool[i].workplace)
			s.idx[workplace] = i
			s.pool[i] = record{workplace: workplace, lastActive: now}
			return true
		}
	}
	return false
}

func (s *slot) shrink(size int) {
	s.idx = make(map[domain.ID]int, size)
	sort.Sort(sort.Reverse(recordsByActivity(s.pool)))
	s.pool = s.pool[:size]
	for i := range s.pool {
		s.idx[s.pool[i].workplace] = i
	}
}

type record struct {
	workplace  domain.ID
	lastActive time.Time
}

func (r *record) touch() {
	r.lastActive = time.Now()
}

type recordsByActivity []record

// Len is the number of elements in the collection.
func (rr recordsByActivity) Len() int {
	return len(rr)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (rr recordsByActivity) Less(i, j int) bool {
	return rr[i].lastActive.Before(rr[j].lastActive)
}

// Swap swaps the elements with indexes i and j.
func (rr recordsByActivity) Swap(i, j int) {
	rr[i], rr[j] = rr[j], rr[i]
}

// issue#draft }
