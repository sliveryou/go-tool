package timex

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLocation(t *testing.T) {
	assertion := assert.New(t)

	assertion.Panics(func() {
		Location("ErrLocationName")
	})
	assertion.NotPanics(func() {
		Location("Asia/Chongqing")
	})
	assertion.NotPanics(func() {
		Shanghai()
	})
	assertion.NotPanics(func() {
		HongKong()
	})
	assertion.NotPanics(func() {
		Local()
	})
	assertion.NotPanics(func() {
		UTC()
	})
}

func TestNow(t *testing.T) {
	t.Log(Now())
	t.Log(NowAdd(time.Hour))
}

func TestNowUnix(t *testing.T) {
	t.Log(NowUnixSecond())
	t.Log(NowUnixMillisecond())
	t.Log(NowUnixMicrosecond())
	t.Log(NowUnixNanosecond())
}

func TestNowAddUnix(t *testing.T) {
	t.Log(NowAddUnixSecond(time.Hour))
	t.Log(NowAddUnixMillisecond(time.Hour))
	t.Log(NowAddUnixMicrosecond(time.Hour))
	t.Log(NowAddUnixNanosecond(time.Hour))
}

func TestUnix(t *testing.T) {
	now := time.Now()
	t.Log(UnixSecond(now))
	t.Log(UnixMillisecond(now))
	t.Log(UnixMicrosecond(now))
	t.Log(UnixNanosecond(now))
}

func TestNowFormat(t *testing.T) {
	t.Log(NowDate())
	t.Log(NowDateTime())
	t.Log(NowFormat("2006/01/02 15:04:05"))
}

func TestNowAddFormat(t *testing.T) {
	t.Log(NowAddDate(time.Hour))
	t.Log(NowAddDateTime(time.Hour))
	t.Log(NowAddFormat(time.Hour, "2006/01/02 15:04:05"))
}

func TestFormat(t *testing.T) {
	now := time.Now()
	t.Log(Date(now))
	t.Log(DateTime(now))
	t.Log(Format(now, "2006/01/02 15:04:05"))
}

func TestUnixToTime(t *testing.T) {
	nowUnixSecond := NowUnixSecond()
	nowUnixMillisecond := NowUnixMillisecond()
	nowUnixMicrosecond := NowUnixMicrosecond()
	nowUnixNanosecond := NowUnixNanosecond()

	t.Log(nowUnixSecond, UnixToTime(nowUnixSecond))
	t.Log(nowUnixMillisecond, UnixToTime(nowUnixMillisecond))
	t.Log(nowUnixMicrosecond, UnixToTime(nowUnixMicrosecond))
	t.Log(nowUnixNanosecond, UnixToTime(nowUnixNanosecond))
}

func TestUnixAddDate(t *testing.T) {
	now := NowUnixSecond()
	t.Log(UnixAddDate(now, 1, 1, 1))
	t.Log(UnixAddYears(now, 1))
	t.Log(UnixAddMonths(now, 1))
	t.Log(UnixAddDays(now, 1))
}

func TestUnixCompare(t *testing.T) {
	assertion := assert.New(t)

	now := NowUnixSecond()
	nowAdd := UnixNanosecond(UnixAddDays(now, 1))
	nowSub := UnixMillisecond(UnixAddDays(now, -1))

	assertion.Less(now, nowAdd)
	assertion.Less(now, nowSub)
	assertion.Less(nowSub, nowAdd)

	assertion.True(UnixBefore(now, nowAdd))
	assertion.True(UnixBefore(nowSub, now))

	assertion.True(UnixAfter(nowAdd, now))
	assertion.True(UnixAfter(now, nowSub))

	assertion.False(UnixEqual(now, nowAdd))
	assertion.False(UnixEqual(nowSub, now))
	assertion.True(UnixEqual(now, now))
}

func TestUnixDiffer(t *testing.T) {
	now := NowUnixNanosecond()
	nowAdd := UnixSecond(UnixAddDate(now, 1, 1, 1))

	assert.Greater(t, now, nowAdd)
	assert.Greater(t, UnixDifferDays(nowAdd, now), 365)
	assert.Greater(t, UnixDifferHours(nowAdd, now), float64(365*24))
}

func TestSleep(t *testing.T) {
	SleepSecond(1)
	t.Log("sleep 1 second")
	SleepMillisecond(1)
	t.Log("sleep 1 millisecond")
	SleepMicrosecond(1)
	t.Log("sleep 1 microsecond")
}

func TestStringToTime(t *testing.T) {
	cases := []struct {
		str     string
		layout  string
		expect  string
		wantErr bool
	}{
		{
			str: "2020-06-01 00:00:00", layout: "2006-01-02 15:04:05",
			expect: "2020-06-01 00:00:00 +0800 CST", wantErr: false,
		},
		{
			str: "2020/06/18 18:00:00", layout: "2006/01/02 15:04:05",
			expect: "2020-06-18 18:00:00 +0800 CST", wantErr: false,
		},
		{
			str: "20200618180000", layout: "20060102150405",
			expect: "2020-06-18 18:00:00 +0800 CST", wantErr: false,
		},
		{
			str: "20200618180000", layout: "20060102",
			expect: "2020-06-18 18:00:00 +0800 CST", wantErr: true,
		},
	}

	for _, c := range cases {
		get, err := StringToTime(c.str, c.layout)
		if c.wantErr {
			assert.EqualError(t, err, "timex: str does not match layout")
		} else if assert.NoError(t, err) {
			assert.Equal(t, c.expect, get.String())
		}
	}
}

func TestStringToUnix(t *testing.T) {
	cases := []struct {
		str    string
		layout string
		expect int64
	}{
		{str: "20091225091010", layout: "20060102150405", expect: 1261703410},
		{str: "2020/06/18 18:00:00", layout: "2006/01/02 15:04:05", expect: 1592474400},
		{str: "2020-06-18 18:00:00", layout: "2006-01-02 15:04:05", expect: 1592474400},
	}

	for _, c := range cases {
		get := StringToUnix(c.str, c.layout)
		assert.Equal(t, c.expect, get)
	}
}

func TestUnixTodayRange(t *testing.T) {
	start, end := UnixTodayRange(Shanghai())
	assert.Less(t, start, end)
	t.Log(start, end)
	t.Log(UnixToTime(start), UnixToTime(end))
}

func TestDoCycleTask(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	DoCycleTask(ctx, func() {
		t.Log("Hello, world!")
	}, time.Second)

	SleepSecond(5)
	cancel()
}
