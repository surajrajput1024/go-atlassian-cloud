package util

import "strconv"

func Int64String(i int64) string {
	return strconv.FormatInt(i, 10)
}

func IntString(i int) string {
	return strconv.Itoa(i)
}
