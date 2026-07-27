package helpers

import "time"

func ParseDate(s string) time.Time {
	t, _ := time.Parse("2006-01-02", s)
	return t
}
