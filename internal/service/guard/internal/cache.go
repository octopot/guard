package internal

import (
	"context"
	"time"

	"github.com/pkg/errors"

	domain "go.octolab.org/ecosystem/guard/internal/service/types"
	"go.octolab.org/ecosystem/guard/internal/storage/types"
	repository "go.octolab.org/ecosystem/guard/internal/storage/types"
)

// TODO issue#draft {

type licenseCache interface {
	LicenseByID(context.Context, domain.ID) (repository.License, error)
	LicenseByEmployee(context.Context, domain.ID) (repository.License, error)
}

// NewLicenseCache returns a blob of memory
// to store licenses.
func NewLicenseCache(in licenseCache) (out licenseCache) {
	c := &cache{
		origin: in,
		idx:    make(map[domain.ID]response, maxLicenses),

		byID:       make(chan request, 1),
		byEmployee: make(chan request, 1),
		memorize:   make(chan response, maxLicenses/100),
	}
	go c.listen()
	return c
}

type cache struct {
	origin licenseCache
	idx    map[domain.ID]response

	memorize   chan response
	byID       chan request
	byEmployee chan request
}

// LicenseByID TODO issue#docs
func (c *cache) LicenseByID(ctx context.Context, id domain.ID) (types.License, error) {
	req := request{ctx, id, make(chan response, 1)}
	c.byID <- req
	resp := <-req.result
	return resp.license, resp.err
}

// LicenseByEmployee TODO issue#docs
func (c *cache) LicenseByEmployee(ctx context.Context, id domain.ID) (types.License, error) {
	req := request{ctx, id, make(chan response, 1)}
	c.byEmployee <- req
	resp := <-req.result
	return resp.license, resp.err
}

func (c *cache) listen() {
	defer func() {
		if r := recover(); r != nil {
			// TODO issue#critical
			// TODO issue#6
			go c.listen()
		}
	}()
	for {
		select {
		case res := <-c.memorize:
			c.idx[res.license.ID] = response{res.license, res.err, time.Now().Add(licenseTTL)}
		case req := <-c.byID:
			res, found := c.idx[req.id]
			if found && res.ttl.After(time.Now()) {
				req.result <- res
				continue
			}
			go func() {
				defer func() {
					if r := recover(); r != nil {
						// TODO issue#critical
						// TODO issue#6
						req.result <- response{err: errors.New("unexpected panic handled")}
					}
				}()
				res.license, res.err = c.origin.LicenseByID(req.ctx, req.id)
				req.result <- res
				c.memorize <- res
			}()
		case req := <-c.byEmployee:
			res, found := c.idx[req.id]
			if found && res.ttl.After(time.Now()) {
				req.result <- res
				continue
			}
			go func() {
				defer func() {
					if r := recover(); r != nil {
						// TODO issue#critical
						// TODO issue#6
						req.result <- response{err: errors.New("unexpected panic handled")}
					}
				}()
				res.license, res.err = c.origin.LicenseByEmployee(req.ctx, req.id)
				req.result <- res
				c.memorize <- res
			}()
		}
	}
}

type request struct {
	ctx    context.Context
	id     domain.ID
	result chan response
}

type response struct {
	license repository.License
	err     error
	ttl     time.Time
}

// issue#draft }
