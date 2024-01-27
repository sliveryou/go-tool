package snowflake

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sliveryou/go-tool/id-generator/encoding/base58"
	"github.com/sliveryou/go-tool/id-generator/encoding/base62"
)

func TestNewSnowflake(t *testing.T) {
	_, err := NewSnowflake(&Config{
		MaxTolerateMillis: -1,
	})
	require.EqualError(t, err, "invalid max tolerate millis")

	_, err = NewSnowflake(&Config{
		NodeId: func() (int64, error) {
			return -1, nil
		},
	})
	require.EqualError(t, err, "invalid node id")

	_, err = NewSnowflake(&Config{
		StartTime: time.Now().Add(time.Minute),
	})
	require.EqualError(t, err, "invalid start time")

	_, err = NewSnowflake(&Config{
		LastGenerateTime: func() (time.Time, error) {
			return time.Now().Add(time.Minute), nil
		},
	})
	require.EqualError(t, err, "invalid last generate time")

	s, err := NewSnowflake(&Config{
		NodeId:            NodeId(1),
		MaxTolerateMillis: 10,
	})
	require.NoError(t, err)
	assert.NotNil(t, s)
}

func TestSnowflake_NextId(t *testing.T) {
	snowflake, err := NewSnowflake(&Config{NodeId: NodeId(1), MaxTolerateMillis: 10})
	require.NoError(t, err)
	assert.NotNil(t, snowflake)

	c := make(chan int64)
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			id, err := snowflake.NextId()
			require.NoError(t, err)
			c <- id
		}()
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	r := make([]int64, 0, 100)
	for s := range c {
		r = append(r, s)
	}

	uniqueMap := make(map[int64]struct{}, len(r))
	for _, s := range r {
		uniqueMap[s] = struct{}{}
	}
	isUnique := len(uniqueMap) == len(r)
	assert.True(t, isUnique)
}

func TestSnowflake_NextId_MultiNodes(t *testing.T) {
	snowflake1, err := NewSnowflake(&Config{NodeId: NodeId(1), MaxTolerateMillis: 10})
	require.NoError(t, err)
	assert.NotNil(t, snowflake1)

	snowflake2, err := NewSnowflake(&Config{NodeId: NodeId(2), MaxTolerateMillis: 10})
	require.NoError(t, err)
	assert.NotNil(t, snowflake2)

	snowflake3, err := NewSnowflake(&Config{NodeId: NodeId(3), MaxTolerateMillis: 10})
	require.NoError(t, err)
	assert.NotNil(t, snowflake3)

	c := make(chan int64)
	wg := sync.WaitGroup{}
	wg.Add(150)

	nextIdFunc := func(s *Snowflake) {
		defer wg.Done()
		id, err := s.NextId()
		require.NoError(t, err)
		c <- id
	}

	for i := 0; i < 50; i++ {
		go nextIdFunc(snowflake1)
		go nextIdFunc(snowflake2)
		go nextIdFunc(snowflake3)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	r := make([]int64, 0, 150)
	for s := range c {
		r = append(r, s)
	}

	uniqueMap := make(map[int64]struct{}, len(r))
	for _, s := range r {
		uniqueMap[s] = struct{}{}
	}
	isUnique := len(uniqueMap) == len(r)
	assert.True(t, isUnique)
}

func TestParse(t *testing.T) {
	snowflake, err := NewSnowflake(&Config{NodeId: NodeId(1), MaxTolerateMillis: 10})
	require.NoError(t, err)
	assert.NotNil(t, snowflake)

	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			id, err := snowflake.NextId()
			require.NoError(t, err)
			t.Logf("id:%v parse:%v base58:%v base62:%v",
				id, Parse(id),
				base58.StdEncoding.Encode(id),
				base62.StdEncoding.Encode(id),
			)
		}()
	}

	wg.Wait()
}

func BenchmarkSnowflake_NextId(b *testing.B) {
	snowflake, _ := NewSnowflake(&Config{NodeId: NodeId(1), MaxTolerateMillis: 10})

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = snowflake.NextId()
	}
}
