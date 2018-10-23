package internal_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"

	domain "github.com/kamilsk/guard/pkg/service/types"

	. "github.com/kamilsk/guard/pkg/service/guard/internal"
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
			r := rand.New(rand.NewSource(time.Now().Unix()))
			<-starter
			license := licenses[r.Intn(len(licenses))]
			counter.Increment(license)
			counter.Rollback(license)
			wg.Done()
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
