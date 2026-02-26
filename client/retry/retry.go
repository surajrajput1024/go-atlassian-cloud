package retry

import "time"

func Backoff(attempt int, min, max time.Duration) time.Duration {
	if min <= 0 {
		min = 500 * time.Millisecond
	}
	if max <= 0 {
		max = 5 * time.Second
	}
	d := min * time.Duration(1<<uint(attempt))
	if d > max {
		return max
	}
	return d
}

func IsRetryableStatusCode(code int) bool {
	return code >= 500 || code == 429
}
