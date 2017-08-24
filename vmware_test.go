package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/calmh/ipfix"
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

func TestIncludeVmwareNsxFields(t *testing.T) {
	s := ipfix.NewSession()
	i := ipfix.NewInterpreter(s)
	includeVmwareNsxFields(i)
}

func TestIncludeVmwareVcenterFields(t *testing.T) {
	s := ipfix.NewSession()
	i := ipfix.NewInterpreter(s)
	includeVmwareVDSFields(i)
}
