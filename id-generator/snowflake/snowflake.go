package snowflake

import (
	"errors"
	"sync"
	"time"
)

// +----------------+------------------------------------------------+-------------------+-----------------------+
// | retain (1 bit) | timestamp (41 bits)                            | node id (10 bits) | sequence id (12 bits) |
// +----------------+------------------------------------------------+-------------------+-----------------------+
// | 0              | 0 0000000000 0000000000 0000000000 0000000000  | 0000000000        | 00 0000000000         |
// +----------------+------------------------------------------------+-------------------+-----------------------+

// snowflake algorithm, used to generate the globally unique id under the distributed system（64 bit）
// 1 bit sign
// 41 bits timestamp (millisecond), subtract the start timestamp from the current timestamp, which can be used 69.7 years after the start timestamp
// 10 bits node id
// 12 bits sequence id

const (
	sequenceBits     uint8 = 12                      // number of bits of sequence id, 12 bits indicates that each machine can generate up to 2^12=4096 unique ids in 1 millisecond
	nodeBits         uint8 = 10                      // number of bits of node id, 10 bits indicates that there can be at most 2^10=1024 nodes
	timestampBits    uint8 = 41                      // number of bits of timestamp, 41 bits indicates that there can be at most 2^41/(24*3600*365*1000)=69.7 years
	sequenceMax      int64 = 1<<sequenceBits - 1     // max value of sequence id, default is 4095
	nodeMax          int64 = 1<<nodeBits - 1         // max value of node id, default is 1023
	timestampMax     int64 = 1<<timestampBits - 1    // max value of timestamp, default is 2199023255551
	nodeShift        uint8 = sequenceBits            // number of left shifts of node id, default is 12
	timestampShift   uint8 = nodeBits + sequenceBits // number of left shifts of timestamp, default is 22
	defaultStartTime int64 = 1590940800000           // default start timestamp (default is 2020-06-01 00:00:00 UTC/GMT +8.00 millisecond timestamp, it cannot be modified after formal use)
)

var (
	// ErrOverTimestampLimit over the timestamp limit error.
	ErrOverTimestampLimit = errors.New("over the timestamp limit")
	// ErrClockBackward clock backward error.
	ErrClockBackward = errors.New("the clock backward")
	// ErrInvalidStartTime invalid start time error.
	ErrInvalidStartTime = errors.New("invalid start time")
	// ErrInvalidLastGenerateTime invalid last generate time error.
	ErrInvalidLastGenerateTime = errors.New("invalid last generate time")
	// ErrInvalidMaxTolerateMillis invalid max tolerate millis error.
	ErrInvalidMaxTolerateMillis = errors.New("invalid max tolerate millis")
	// ErrInvalidNodeId invalid node id error.
	ErrInvalidNodeId = errors.New("invalid node id")
)

// Config snowflake generator config.
type Config struct {
	StartTime         time.Time                 // start time, default is 2020-06-01 00:00:00 UTC/GMT +8.00
	MaxTolerateMillis int64                     // max tolerated clock fallback milliseconds
	LastGenerateTime  func() (time.Time, error) // last id generation time, default is the current time
	NodeId            func() (int64, error)     // node id
}

// Snowflake distributed unique id generator based on snowflake algorithm.
type Snowflake struct {
	mutex          *sync.Mutex // mutex, used to ensure concurrency security
	tolerateMillis int64       // max tolerated clock fallback milliseconds
	startTime      int64       // start timestamp (millisecond)
	elapsedTime    int64       // elapsed timestamp (millisecond)
	nodeId         int64       // node id
	sequenceId     int64       // sequence id
}

// NewSnowflake new a snowflake generator.
func NewSnowflake(c *Config) (*Snowflake, error) {
	s := new(Snowflake)
	s.mutex = new(sync.Mutex)
	s.sequenceId = sequenceMax

	if c.StartTime.After(time.Now()) {
		return nil, ErrInvalidStartTime
	}

	if c.StartTime.IsZero() {
		s.startTime = defaultStartTime
	} else {
		s.startTime = unixMilli(c.StartTime)
	}

	if c.MaxTolerateMillis < 0 {
		return nil, ErrInvalidMaxTolerateMillis
	}
	s.tolerateMillis = c.MaxTolerateMillis

	if c.LastGenerateTime != nil {
		lastGenerateTime, err := c.LastGenerateTime()
		if err != nil {
			return nil, err
		}
		if !lastGenerateTime.IsZero() {
			elapsedTime := unixMilli(lastGenerateTime) - s.startTime
			if elapsedTime > currentElapsedTime(s.startTime) {
				return nil, ErrInvalidLastGenerateTime
			}
			s.elapsedTime = elapsedTime
		}
	}

	if c.NodeId != nil {
		nodeId, err := c.NodeId()
		if err != nil {
			return nil, err
		}
		if nodeId < 0 || nodeId > nodeMax {
			return nil, ErrInvalidNodeId
		}
		s.nodeId = nodeId
	}

	return s, nil
}

// NextId generates snowflake id.
func (s *Snowflake) NextId() (int64, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	current := currentElapsedTime(s.startTime)
	if s.elapsedTime < current {
		s.sequenceId = 0
		s.elapsedTime = current
	} else {
		s.sequenceId = (s.sequenceId + 1) & sequenceMax
		if s.sequenceId == 0 {
			if s.elapsedTime-current <= s.tolerateMillis {
				s.elapsedTime = s.waitNextElapsedTime(current)
			} else {
				return 0, ErrClockBackward
			}
		}
	}

	return s.calculateId()
}

// waitNextElapsedTime waits and returns the elapsed timestamp after growth.
func (s *Snowflake) waitNextElapsedTime(current int64) int64 {
	for current <= s.elapsedTime {
		current = currentElapsedTime(s.startTime)
	}

	return current
}

// calculateId calculates snowflake id.
func (s *Snowflake) calculateId() (int64, error) {
	if s.elapsedTime > timestampMax {
		return 0, ErrOverTimestampLimit
	}

	return s.elapsedTime<<timestampShift | s.nodeId<<nodeShift | s.sequenceId, nil
}

// currentElapsedTime gets the current elapsed timestamp (subtract the start timestamp from the current timestamp).
func currentElapsedTime(startTime int64) int64 {
	return unixMilli(time.Now()) - startTime
}

// unixMilli gets the millisecond timestamp of the specified time.
func unixMilli(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// NodeId returns the node id constructor.
func NodeId(nodeId int64) func() (int64, error) {
	return func() (int64, error) { return nodeId, nil }
}

// Parse parses snowflake id.
func Parse(id int64, startTime ...time.Time) map[string]int64 {
	st := defaultStartTime
	if len(startTime) != 0 {
		st = unixMilli(startTime[0])
	}

	elapsedTime := id >> timestampShift
	nodeId := id & (nodeMax << sequenceBits) >> sequenceBits
	sequenceId := id & sequenceMax

	return map[string]int64{
		"id":           id,
		"startTime":    st,
		"elapsedTime":  elapsedTime,
		"generateTime": st + elapsedTime,
		"nodeId":       nodeId,
		"sequenceId":   sequenceId,
	}
}
