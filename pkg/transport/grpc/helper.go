package grpc

import (
	"time"

	domain "github.com/kamilsk/guard/pkg/service/types"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/pkg/errors"
)

// Timestamp converts a time.Time to a google.protobuf.Timestamp proto.
// It panics if the resulting Timestamp is invalid.
func Timestamp(tp *time.Time) *timestamp.Timestamp {
	if tp == nil {
		return nil
	}
	ts, err := ptypes.TimestampProto(*tp)
	if err != nil {
		panic(errors.Wrapf(err, "converting %#v into google.protobuf.Timestamp", *tp))
	}
	return ts
}

// Time converts a google.protobuf.Timestamp proto to a time.Time.
// It panics if the passed Timestamp is invalid.
func Time(ts *timestamp.Timestamp) *time.Time {
	if ts == nil {
		return nil
	}
	tp, err := ptypes.Timestamp(ts)
	if err != nil {
		panic(errors.Wrapf(err, "converting %#v into time.Time", *ts))
	}
	return &tp
}

func ptrToID(id string) *domain.ID {
	if id == "" {
		return nil
	}
	ptr := new(domain.ID)
	*ptr = domain.ID(id)
	return ptr
}

func convertToDomainContract(from *Contract) (to domain.Contract) {
	if from == nil {
		return
	}
	to.Request, to.Workplaces = from.Requests, from.Workplaces
	to.Since, to.Until = Time(from.Since), Time(from.Until)
	if from.Rate != nil {
		to.Rate = domain.PackRate(domain.RateValue(from.Rate.Value), convertToDomainRateUnit(from.Rate.Unit))
	}
	return
}

func convertToDomainRateUnit(enum Rate_Unit) domain.RateUnit {
	switch enum {
	case Rate_rps:
		return domain.RPS
	case Rate_rpm:
		return domain.RPM
	case Rate_rph:
		return domain.RPH
	case Rate_rpd:
		return domain.RPD
	case Rate_rpw:
		return domain.RPW
	default:
		panic(errors.Errorf("unknown protobuf unit %v", enum))
	}
}
