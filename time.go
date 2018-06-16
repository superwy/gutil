package gutil

import (
	"github.com/pkg/errors"
	"strings"
	"time"
)

func TickParse(layout, value string) (int64, error) {
	if trimValue := strings.TrimSpace(value); trimValue != "" {
		tm, err := time.ParseInLocation(layout, trimValue, time.Local)
		if err != nil {
			return 0, err
		}
		if tm.IsZero() {
			return 0, errors.New("zero time")
		}
		return tm.UnixNano() / 1e6, nil
	}
	return 0, errors.New("empty value")
}

func TickNowFirst() (tick int64) {
	layout := "2006-01-02"
	strDate := time.Now().Format(layout)
	tick, _ = TickParse(layout, strDate)
	return
}

func TickNowLast() (tick int64) {
	layout := "2006-01-02"
	strDate := time.Now().Format(layout)
	tm, _ := time.ParseInLocation(layout, strDate, time.Local)
	tm = tm.Add(time.Duration(23) * time.Hour)
	tm = tm.Add(time.Duration(59) * time.Minute)
	tm = tm.Add(time.Duration(59) * time.Second)
	tick = tm.UnixNano() / 1e6
	return
}

func TickNow() int64 {
	return time.Now().UnixNano() / 1e6
}

// 如： 201801, 201712
func YearMonth(tick int64) int {
	tm := time.Unix(tick/1e3, 0)
	return tm.Year()*100 + int(tm.Month())
}

// 和 MonthVal 对应，转为一个时间
func YearMonth2Time(monthVal int, day int) time.Time {
	year := monthVal / 100
	month := monthVal % 100
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Now().Location())
}

func MonthDays(tm time.Time) int {
	newMonth := time.Date(tm.Year(), tm.Month(), 1, 0, 0, 0, 0, tm.Location()).AddDate(0, 1, 0)
	return newMonth.AddDate(0, 0, -1).Day()
}

//tick to time
func Tick2Time(tick int64) time.Time {
	return time.Unix(tick/1e3, 0)
}
