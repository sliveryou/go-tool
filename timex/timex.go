package timex

import (
	"context"
	"errors"
	"log"
	"time"
)

// time package variables.
var (
	local    = time.Local                // Local *time.Location
	utc      = time.UTC                  // UTC *time.Location
	shanghai = Location("Asia/Shanghai") // Shanghai *time.Location
)

// Location returns *time.Location by location name.
func Location(name string) *time.Location {
	loc, err := time.LoadLocation(name)
	if err != nil {
		loc = time.Local
	}
	return loc
}

// Shanghai returns Shanghai *time.Location.
func Shanghai() *time.Location {
	return shanghai
}

// Local returns Local *time.Location.
func Local() *time.Location {
	return local
}

// UTC returns UTC *time.Location.
func UTC() *time.Location {
	return utc
}

// NowAdd returns now added time.Time.
func NowAdd(d time.Duration, location ...*time.Location) time.Time {
	return Now(location...).Add(d)
}

// Now returns now time.Time.
func Now(location ...*time.Location) time.Time {
	loc := getLocation(location...)
	return time.Now().In(loc)
}

// NowUnixSecond returns now unix second timestamp.
func NowUnixSecond(location ...*time.Location) int64 {
	return UnixSecond(time.Now(), location...)
}

// NowUnixMillisecond returns now unix millisecond timestamp.
func NowUnixMillisecond(location ...*time.Location) int64 {
	return UnixMillisecond(time.Now(), location...)
}

// NowUnixMicrosecond returns now unix microsecond timestamp.
func NowUnixMicrosecond(location ...*time.Location) int64 {
	return UnixMicrosecond(time.Now(), location...)
}

// NowUnixNanosecond returns now unix nanosecond timestamp.
func NowUnixNanosecond(location ...*time.Location) int64 {
	return UnixNanosecond(time.Now(), location...)
}

// NowAddUnixSecond returns now added unix second timestamp.
func NowAddUnixSecond(d time.Duration, location ...*time.Location) int64 {
	return UnixSecond(time.Now().Add(d), location...)
}

// NowAddUnixMillisecond returns now added unix millisecond timestamp.
func NowAddUnixMillisecond(d time.Duration, location ...*time.Location) int64 {
	return UnixMillisecond(time.Now().Add(d), location...)
}

// NowAddUnixMicrosecond returns now added unix microsecond timestamp.
func NowAddUnixMicrosecond(d time.Duration, location ...*time.Location) int64 {
	return UnixMicrosecond(time.Now().Add(d), location...)
}

// NowAddUnixNanosecond returns now added unix nanosecond timestamp.
func NowAddUnixNanosecond(d time.Duration, location ...*time.Location) int64 {
	return UnixNanosecond(time.Now().Add(d), location...)
}

// UnixSecond returns t as unix second timestamp.
func UnixSecond(t time.Time, location ...*time.Location) int64 {
	loc := getLocation(location...)
	return t.In(loc).Unix()
}

// UnixMillisecond returns t as unix millisecond timestamp.
func UnixMillisecond(t time.Time, location ...*time.Location) int64 {
	return UnixNanosecond(t, location...) / int64(time.Millisecond)
}

// UnixMicrosecond returns t as unix microsecond timestamp.
func UnixMicrosecond(t time.Time, location ...*time.Location) int64 {
	return UnixNanosecond(t, location...) / int64(time.Microsecond)
}

// UnixNanosecond returns t as unix nanosecond timestamp.
func UnixNanosecond(t time.Time, location ...*time.Location) int64 {
	loc := getLocation(location...)
	return t.In(loc).UnixNano()
}

// NowDate returns a date representation of now time value.
func NowDate(location ...*time.Location) string {
	return Date(time.Now(), location...)
}

// NowDateTime returns a datetime representation of now time value.
func NowDateTime(location ...*time.Location) string {
	return DateTime(time.Now(), location...)
}

// NowFormat returns a textual representation of now time value formatted according to layout.
func NowFormat(layout string, location ...*time.Location) string {
	return Format(time.Now(), layout, location...)
}

// NowAddDate returns a date representation of now added time value.
func NowAddDate(d time.Duration, location ...*time.Location) string {
	return Date(time.Now().Add(d), location...)
}

// NowAddDateTime returns a datetime representation of now added time value.
func NowAddDateTime(d time.Duration, location ...*time.Location) string {
	return DateTime(time.Now().Add(d), location...)
}

// NowAddFormat returns a textual representation of now added time value formatted according to layout.
func NowAddFormat(d time.Duration, layout string, location ...*time.Location) string {
	return Format(time.Now().Add(d), layout, location...)
}

// Date returns a date representation of the time value t.
func Date(t time.Time, location ...*time.Location) string {
	return Format(t, "2006-01-02", location...)
}

// DateTime returns a datetime representation of the time value t.
func DateTime(t time.Time, location ...*time.Location) string {
	return Format(t, "2006-01-02 15:04:05", location...)
}

// Format returns a textual representation of the time value t formatted according to layout.
func Format(t time.Time, layout string, location ...*time.Location) string {
	loc := getLocation(location...)
	return t.In(loc).Format(layout)
}

// UnixToTime returns time.Time by unix timestamp.
func UnixToTime(timestamp int64, location ...*time.Location) time.Time {
	var t time.Time
	loc := getLocation(location...)
	switch {
	case timestamp < 1e10:
		t = time.Unix(timestamp, 0)
	case timestamp < 1e13:
		t = time.Unix(0, timestamp*int64(time.Millisecond))
	case timestamp < 1e16:
		t = time.Unix(0, timestamp*int64(time.Microsecond))
	default:
		t = time.Unix(0, timestamp)
	}

	return t.In(loc)
}

// UnixAddDate returns time.Time after unix timestamp has been added date.
func UnixAddDate(timestamp int64, years, months, days int, location ...*time.Location) time.Time {
	t := UnixToTime(timestamp, location...)
	return t.AddDate(years, months, days)
}

// UnixAddYears returns time.Time after unix timestamp has been added years.
func UnixAddYears(timestamp int64, years int, location ...*time.Location) time.Time {
	return UnixAddDate(timestamp, years, 0, 0, location...)
}

// UnixAddMonths returns time.Time after unix timestamp has been added months.
func UnixAddMonths(timestamp int64, months int, location ...*time.Location) time.Time {
	return UnixAddDate(timestamp, 0, months, 0, location...)
}

// UnixAddDays returns time.Time after unix timestamp has been added days.
func UnixAddDays(timestamp int64, days int, location ...*time.Location) time.Time {
	return UnixAddDate(timestamp, 0, 0, days, location...)
}

// UnixEqual reports whether timestamp1 is equal timestamp2.
func UnixEqual(timestamp1, timestamp2 int64) bool {
	t1, t2 := UnixToTime(timestamp1), UnixToTime(timestamp2)
	return t1.Equal(t2)
}

// UnixBefore reports whether timestamp1 is before timestamp2.
func UnixBefore(timestamp1, timestamp2 int64) bool {
	t1, t2 := UnixToTime(timestamp1), UnixToTime(timestamp2)
	return t1.Before(t2)
}

// UnixAfter reports whether timestamp1 is after timestamp2.
func UnixAfter(timestamp1, timestamp2 int64) bool {
	t1, t2 := UnixToTime(timestamp1), UnixToTime(timestamp2)
	return t1.After(t2)
}

// UnixDifferDays returns the number of days between two timestamp.
func UnixDifferDays(timestamp1, timestamp2 int64) int {
	t1, t2 := UnixToTime(timestamp1), UnixToTime(timestamp2)
	return int(t1.Sub(t2).Hours() / 24)
}

// UnixDifferHours returns the number of hours between two timestamp.
func UnixDifferHours(timestamp1, timestamp2 int64) float64 {
	t1, t2 := UnixToTime(timestamp1), UnixToTime(timestamp2)
	return t1.Sub(t2).Hours()
}

// SleepSecond pauses the current goroutine for at least n second.
func SleepSecond(n int64) {
	time.Sleep(time.Duration(n) * time.Second)
}

// SleepMillisecond pauses the current goroutine for at least n millisecond.
func SleepMillisecond(n int64) {
	time.Sleep(time.Duration(n) * time.Millisecond)
}

// SleepMicrosecond pauses the current goroutine for at least n microsecond.
func SleepMicrosecond(n int64) {
	time.Sleep(time.Duration(n) * time.Microsecond)
}

// StringToTime returns time.Time representation of str value parsed according to layout.
// layout example:
//
//	20060102150405
//	2006-01-02 15:04:05
//	2006/01/02 15/04/05
func StringToTime(str, layout string, location ...*time.Location) (time.Time, error) {
	loc := getLocation(location...)
	if len(str) != len(layout) {
		return time.Now(), errors.New("timex: str does not match layout")
	}

	return time.ParseInLocation(layout, str, loc)
}

// StringToUnix returns unix second timestamp representation of str value parsed according to layout.
// If str parsed err, it returns now unix second timestamp.
// layout example:
//
//	20060102150405
//	2006-01-02 15:04:05
//	2006/01/02 15/04/05
func StringToUnix(str, layout string, location ...*time.Location) int64 {
	t, err := StringToTime(str, layout, location...)
	if err != nil {
		return NowUnixSecond()
	}

	return t.Unix()
}

// UnixTodayRange returns today start unix second timestamp and today end unix second timestamp.
func UnixTodayRange(location ...*time.Location) (int64, int64) {
	loc := getLocation(location...)
	year, month, day := time.Now().In(loc).Date()
	start := time.Date(year, month, day, 0, 0, 0, 0, loc).Unix()
	end := time.Date(year, month, day, 23, 59, 59, 0, loc).Unix()

	return start, end
}

// DoCycleTask processing the cycle task with a period specified by the duration.
func DoCycleTask(ctx context.Context, f func(), d time.Duration) {
	if d == 0 {
		return
	}

	t := time.NewTicker(d)
	go func() {
		defer func() {
			t.Stop()

			if r := recover(); r != nil {
				log.Printf("timex: DoCycleTask panic: %v", r)
			}
		}()

		f()
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				f()
			}
		}
	}()
}

func getLocation(location ...*time.Location) *time.Location {
	loc := Shanghai()
	if len(location) != 0 {
		loc = location[0]
	}

	return loc
}
