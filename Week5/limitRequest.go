package main

import (
	"fmt"
	"net/http"
	"time"
)

var window []int64 = []int64{}

func limit(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Too many requests\n")
}

func getTimeStamp() int64 {
	return time.Now().UnixNano() / 1000 / 1000 // 转换成毫秒
}

func LimitRequest(rw http.ResponseWriter, r *http.Request, handler func(http.ResponseWriter, *http.Request), maxQps int) {

	now := getTimeStamp()
	start := now - 1000

	for len(window) > 0 && window[0] < start {
		window = window[1:]
	}

	if len(window) >= maxQps {
		limit(rw, r)
		return
	}
	window = append(window, now)

	handler(rw, r)
}
