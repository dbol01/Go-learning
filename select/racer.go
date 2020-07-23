package racer

import (
	"fmt"
	"net/http"
	"time"
)

// func Racer(url1, url2 string) (winner string) {
// 	aDuration := measureTime(url1)
// 	bDuration := measureTime(url2)

// 	if aDuration > bDuration {
// 		return url2
// 	}

// 	return url1
// }

// func measureTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	Duration := time.Since(start)

// 	return Duration
// }

var tenSecondTimout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
