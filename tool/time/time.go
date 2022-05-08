package time

import "time"

type Time struct {
	startTime time.Time
}

func (t *Time) GetTimeStr() string {
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	durationStr := time.Since(t.startTime).String()
	return nowStr + " 当前:" + durationStr + "\n"
}

var GTime = Time{time.Now()}
