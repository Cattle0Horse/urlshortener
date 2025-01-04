package tools

import (
	"errors"
	"sync"
	"time"
)

const (
	epoch          = int64(1672531200000) // 2023-01-01 00:00:00 UTC
	workerIDBits   = uint(10)
	sequenceBits   = uint(12)
	workerIDShift  = sequenceBits
	timestampShift = sequenceBits + workerIDBits
	sequenceMask   = int64(-1) ^ (int64(-1) << sequenceBits)
	maxWorkerID    = int64(-1) ^ (int64(-1) << workerIDBits)
)

type Snowflake struct {
	mu        sync.Mutex
	lastStamp int64
	workerID  int64
	sequence  int64
}

func NewSnowflake(workerID int64) (*Snowflake, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, errors.New("worker ID out of range")
	}
	return &Snowflake{
		lastStamp: -1,
		workerID:  workerID,
		sequence:  0,
	}, nil
}

func (s *Snowflake) NextID() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	timestamp := time.Now().UnixMilli()

	if timestamp == s.lastStamp {
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			for timestamp <= s.lastStamp {
				timestamp = time.Now().UnixMilli()
			}
		}
	} else {
		s.sequence = 0
	}

	s.lastStamp = timestamp

	return ((timestamp - epoch) << timestampShift) |
		(s.workerID << workerIDShift) |
		s.sequence
}
