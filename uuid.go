package uuid

import (
	"time"
)

const Num100nsIntervalsSinceUUIDEpoch = 0x01b21dd213814000

type UUID struct {
	MostSigBits  int64
	LeastSigBits int64
}

func (p *UUID) Timestamp() uint64 {
	return (uint64(p.MostSigBits)&0x0FFF)<<48 | ((uint64(p.MostSigBits)>>16)&0x0FFFF)<<32 | uint64(p.MostSigBits)>>32
}

func (p *UUID) UnixTimestamp() int64 {
	return int64((p.Timestamp() - Num100nsIntervalsSinceUUIDEpoch) / 10000)
}

func (p *UUID) Time() time.Time {
	return time.Unix(int64(p.UnixTimestamp()/1000), 0)
}
