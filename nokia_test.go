package main

import (
	"testing"

	"github.com/calmh/ipfix"
)

func TestIncludeNokiaFields(t *testing.T) {
	s := ipfix.NewSession()
	i := ipfix.NewInterpreter(s)
	includeNokiaFields(i)
}
