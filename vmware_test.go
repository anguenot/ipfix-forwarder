package main

import (
	"testing"
)

func TestGetNSXSegmentId01(t *testing.T) {
	const expected = 70010
	got := getNSXSegmentId(72057594037997946)
	if got != expected {
		t.Error(
			"Expected", expected,
			"got", got,
		)
	}
}

func TestGetNSXSegmentId02(t *testing.T) {
	const expected = 7015
	got := getNSXSegmentId(72057594037934951)
	if got != expected {
		t.Error(
			"Expected", expected,
			"got", got,
		)
	}
}
