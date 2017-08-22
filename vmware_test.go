package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNSXSegmentId01(t *testing.T) {
	const expected = 70010
	segmentID := getNSXSegmentID(72057594037997946)
	assert.Equal(t, expected, int(segmentID))
}

func TestGetNSXSegmentId02(t *testing.T) {
	const expected = 7015
	segmentID := getNSXSegmentID(72057594037934951)
	assert.Equal(t, expected, int(segmentID))
}
