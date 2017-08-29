package main

import (
	"time"
	"github.com/golang/glog"
)

// exponential back-off retry base delay
var (
	baseRetryDelay      = 5 * time.Second
	maxRetries     uint = 3
)

// increment error count and exponential back-off sleep for next retry.
func incErrorCountAndSleep(err error, errCount *uint) {
	if err != nil {
		glog.Errorln(err)
	}
	*errCount ++
	delay := baseRetryDelay * (1 << *errCount) // exponential back-off delay
	glog.Errorf(
		"Error #%d: sleeping for %s before retrying...", *errCount, delay)
	time.Sleep(delay)
}
