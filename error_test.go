package main

import (
	"testing"
	"time"

	"errors"
	"github.com/stretchr/testify/assert"
)

func TestIncErrorAndSleep(t *testing.T) {

	// fake error
	err := errors.New("FAKE")

	// decrease exponential back-off base retry delay.
	baseRetryDelay = 1 * time.Microsecond

	errCounts := []uint{0, 1, 2}
	for _, errCount := range errCounts {
		start := time.Now()
		incErrorCountAndSleep(err, &errCount)
		elapsed := time.Since(start)
		expectedErrorCount := errCount
		assert.Equal(t, errCount, expectedErrorCount)
		assert.Equal(t, true, elapsed >= (baseRetryDelay * (1 << errCount)))
	}

}
