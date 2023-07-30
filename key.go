package hashring

import (
	"encoding/binary"
	"fmt"
)

type Int64PairHashKey struct {
	High int64
	Low  int64
}

func (k *Int64PairHashKey) Less(other HashKey) int64 {
	o := other.(*Int64PairHashKey)
	high := k.High - o.High
	if high != 0 {
		return high
	}
	return k.Low - o.Low
}

func NewInt64PairHashKey(bytes []byte) (HashKey, error) {
	const expected = 16
	if len(bytes) != expected {
		return nil, fmt.Errorf(
			"expected %d bytes, got %d bytes",
			expected, len(bytes),
		)
	}
	return &Int64PairHashKey{
		High: int64(binary.LittleEndian.Uint64(bytes[:8])),
		Low:  int64(binary.LittleEndian.Uint64(bytes[8:])),
	}, nil
}
