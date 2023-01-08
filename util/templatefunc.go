package util

import "time"

func UnixToTime(timi int) string {
	t := time.Unix(int64(timi), 0)
	return t.Format("2006-01-02 15:14:13")
}
