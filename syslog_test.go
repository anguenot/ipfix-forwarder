package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyslogWriter(t *testing.T) {
	writer := getSyslogWriter()
	assert.NotNil(t, writer)
}

func TestSendToSyslog(t *testing.T) {
	err := sendToSyslog("testing")
	assert.Nil(t, err)
}
