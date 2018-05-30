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
