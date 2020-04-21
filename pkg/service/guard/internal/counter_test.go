package internal_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	. "github.com/kamilsk/guard/pkg/service/guard/internal"
	domain "github.com/kamilsk/guard/pkg/service/types"
)

func TestLicenseRequestCounter(t *testing.T) {
	var wg sync.WaitGroup

	threads := runtime.GOMAXPROCS(0)
	starter := make(chan struct{})
	counter := NewLicenseRequestCounter(threads)

	licenses := make([]domain.ID, threads)
	for i := range licenses {
		licenses[i] = domain.ID(fmt.Sprintf("10000000-2000-4000-8000-1600000000%02d", i+1))
	}

	requests := 3 * threads
	wg.Add(requests)
	for i := 0; i < requests; i++ {
		go func() {
			defer wg.Done()

			r := rand.New(rand.NewSource(time.Now().Unix()))
			<-starter

			license := licenses[r.Intn(len(licenses))]
			counter.Increment(license)
			counter.Rollback(license)
		}()
	}
	close(starter)
	wg.Wait()

	for _, license := range licenses {
		if counter.Increment(license) != 1 {
			t.Fail()
		}
	}
}

func BenchmarkLicenseRequestCounter(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	licenses := make([]domain.ID, runtime.GOMAXPROCS(0))
	for i := range licenses {
		licenses[i] = domain.ID(fmt.Sprintf("10000000-2000-4000-8000-1600000000%02d", i+1))
	}
	counter := NewLicenseRequestCounter(len(licenses))

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		counter.Increment(licenses[r.Intn(len(licenses))])
	}
}

func TestLicenseWorkplaceCounter(t *testing.T) {
	var wg sync.WaitGroup

	threads := runtime.GOMAXPROCS(0)
	starter := make(chan struct{})
	counter := NewLicenseWorkplaceCounter(threads, threads)

	ids := make([]domain.ID, threads)
	for i := range ids {
		ids[i] = domain.ID(fmt.Sprintf("10000000-2000-4000-8000-1600000000%02d", i+1))
	}

	passed, requests := uint32(0), 3*threads
	wg.Add(requests)
	for i := 0; i < requests; i++ {
		go func() {
			defer wg.Done()

			r := rand.New(rand.NewSource(time.Now().Unix()))
			<-starter

			license, workplace := ids[r.Intn(len(ids))], ids[r.Intn(len(ids))]
			if counter.Acquire(license, workplace, threads) {
				atomic.AddUint32(&passed, 1)
			}
		}()
	}
	close(starter)
	wg.Wait()

	if passed == 0 {
		t.Fail()
	}
}

func BenchmarkLicenseWorkplaceCounter(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	ids := make([]domain.ID, runtime.GOMAXPROCS(0))
	for i := range ids {
		ids[i] = domain.ID(fmt.Sprintf("10000000-2000-4000-8000-1600000000%02d", i+1))
	}
	counter := NewLicenseWorkplaceCounter(len(ids), len(ids))

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		counter.Acquire(ids[r.Intn(len(ids))], ids[r.Intn(len(ids))], len(ids))
	}
}
