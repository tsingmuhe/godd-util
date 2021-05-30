package gtime

import "time"

const (
	ISO8601Date      = "2006-01-02"
	ISO8601DateTime  = "2006-01-02T15:04:05"
	ReadableDateTime = "2006-01-02 15:04:05"
)

var UTCE8 = time.FixedZone("UTC+8", 8*60*60)

func WithUnix(sec int64) *time.Time {
	tm := time.Unix(sec, 0)
	return &tm
}

func WithUnixMilli(milli int64) *time.Time {
	sec := milli / 1e3
	nsec := (milli % 1e3) * time.Millisecond
	tm := time.Unix(sec, nsec)
	return &tm
}

func LocalDate(year int, month time.Month, day, hour, min, sec int) *time.Time {
	return Date(year, month, day, hour, min, sec, time.Local)
}

func Date(year int, month time.Month, day, hour, min, sec int, loc *time.Location) *time.Time {
	tm := time.Date(year, month, day, hour, min, sec, 0, loc)
	return &tm
}

func ParseLocalDate(layout, value string) (time.Time, error) {
	return ParseDate(layout, value, time.Local)
}

func ParseDate(layout, value string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(layout, value, loc)
}

func DateString(t *time.Time) string {
	return t.Format(ISO8601Date)
}

func DatetimeString(t *time.Time) string {
	return t.Format(ReadableDateTime)
}

func UnixSecond(t *time.Time) int64 {
	return t.UnixNano() / 1e9
}

func UnixMilli(t *time.Time) int64 {
	return t.UnixNano() / 1e6
}

func UnixMicro(t *time.Time) int64 {
	return t.UnixNano() / 1e3
}

func UnixNano(t *time.Time) int64 {
	return t.UnixNano()
}

func PlusHours(t *time.Time, n int64) *time.Time {
	tm := t.Add(time.Hour * time.Duration(n))
	return &tm
}

func PlusMinutes(t *time.Time, n int64) *time.Time {
	tm := t.Add(time.Minute * time.Duration(n))
	return &tm
}

func PlusSeconds(t *time.Time, n int64) *time.Time {
	tm := t.Add(time.Second * time.Duration(n))
	return &tm
}

func PlusMillis(t *time.Time, n int64) *time.Time {
	tm := t.Add(time.Millisecond * time.Duration(n))
	return &tm
}

func MinusHours(t *time.Time, n int64) *time.Time {
	tm := t.Add(time.Hour * time.Duration(-n))
	return &tm
}

func MinusMinutes(t *time.Time, n int64) *time.Time {
	tm := t.Add(time.Minute * time.Duration(-n))
	return &tm
}

func MinusSeconds(t *time.Time, n int64) *time.Time {
	tm := t.Add(time.Second * time.Duration(-n))
	return &tm
}

func MinusMillis(t *time.Time, n int64) *time.Time {
	tm := t.Add(time.Millisecond * time.Duration(-n))
	return &tm
}

func StartOfHour(t *time.Time) *time.Time {
	y, m, d := t.Date()
	tm := time.Date(y, m, d, 0, 0, 0, 0, t.Location())
	return &tm
}

func StartOfDay(t *time.Time) *time.Time {
	y, m, d := t.Date()
	tm := time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())
	return &tm
}

func StartOfMonth(t *time.Time) *time.Time {
	y, m, _ := t.Date()
	tm := time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
	return &tm
}

func StartOfYear(t *time.Time) *time.Time {
	y, _, _ := t.Date()
	tm := time.Date(y, time.January, 1, 0, 0, 0, 0, t.Location())
	return &tm
}

func EndOfHour(t *time.Time) *time.Time {
	y, m, d := t.Date()
	tm := time.Date(y, m, d, t.Hour(), 59, 59, int(time.Second-time.Nanosecond), t.Location())
	return &tm
}

func EndOfDay(t *time.Time) *time.Time {
	y, m, d := t.Date()
	tm := time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
	return &tm
}

func EndOfMonth(t *time.Time) *time.Time {
	y, m, _ := t.Date()
	tm := time.Date(y, m+1, 1, 0, 0, 0, 0, t.Location()).Add(-time.Nanosecond)
	return &tm
}

func EndOfYear(t *time.Time) *time.Time {
	y, _, _ := t.Date()
	tm := time.Date(y+1, time.January, 1, 0, 0, 0, 0, t.Location()).Add(-time.Nanosecond)
	return &tm
}
