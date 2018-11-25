package rpc

import (
	"time"

	domain "github.com/kamilsk/guard/pkg/service/types"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/kamilsk/guard/pkg/transport/grpc/protobuf"
	"github.com/pkg/errors"
)

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

func convertFromDomainContract(from domain.Contract) *protobuf.Contract {
	to := &protobuf.Contract{Requests: from.Requests, Workplaces: from.Workplaces}
	to.Since, to.Until = Timestamp(from.Since), Timestamp(from.Until)
	if !from.Rate.IsEmpty() {
		value, unit := from.Rate.Value()
		to.Rate = &protobuf.Rate{Value: value, Unit: units.convert(domain.RateUnit(unit))}
	}
	return to
}

func convertToDomainContract(from *protobuf.Contract) (to domain.Contract) {
	if from == nil {
		return
	}
	to.Requests, to.Workplaces = from.Requests, from.Workplaces
	to.Since, to.Until = Time(from.Since), Time(from.Until)
	if from.Rate != nil {
		to.Rate = domain.PackRate(domain.RateValue(from.Rate.Value), units.invert(from.Rate.Unit))
	}
	return
}

func ptrToID(id string) *domain.ID {
	if id == "" {
		return nil
	}
	ptr := new(domain.ID)
	*ptr = domain.ID(id)
	return ptr
}

func ptrToToken(token string) *domain.Token {
	if token == "" {
		return nil
	}
	ptr := new(domain.Token)
	*ptr = domain.Token(token)
	return ptr
}

type unitMap map[domain.RateUnit]protobuf.Rate_Unit

func (m unitMap) convert(from domain.RateUnit) protobuf.Rate_Unit {
	to, found := m[from]
	if !found {
		panic(errors.Errorf("unexpected domain rate unit %v", from))
	}
	return to
}

func (m unitMap) invert(from protobuf.Rate_Unit) domain.RateUnit {
	for to, v := range m {
		if v == from {
			return to
		}
	}
	panic(errors.Errorf("unexpected protobuf rate unit %v", from))
}

var units = unitMap{
	domain.RPS: protobuf.Rate_rps,
	domain.RPM: protobuf.Rate_rpm,
	domain.RPH: protobuf.Rate_rph,
	domain.RPD: protobuf.Rate_rpd,
	domain.RPW: protobuf.Rate_rpw,
}
