package main

import (
	"os"
	"os/exec"
	"testing"
	"syscall"
)

func TestUsage(t *testing.T) {
	if os.Getenv("I_AM_USAGE") == "1" {
		usage()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestUsage")
	cmd.Env = append(os.Environ(), "I_AM_USAGE=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		waitStatus := e.Sys().(syscall.WaitStatus)
		if waitStatus.ExitStatus() == 2 {
			return
		}

	}
	t.Fatalf("err %v but expected exit status 2", err)
}

func TestParseOptions(t *testing.T) {
	if os.Getenv("I_AM_PARSE_OPTIONS") == "1" {
		parseOptions()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestParseOptions")
	cmd.Env = append(os.Environ(), "I_AM_PARSE_OPTIONS=1")
	err := cmd.Run()
	if err == nil {
		return
	}
	t.Fatalf("err %v but none expected ", err)
}
