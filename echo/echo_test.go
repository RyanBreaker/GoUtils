package main

import (
	"testing"
	"os/exec"
)

var args = []string{"run", "echo.go"}

func TestNothing(t *testing.T) {
	echo := exec.Command("go", args...)
	echoOut, _ := echo.Output()

	if string(echoOut) != "\n" {
		t.Fail()
	}
}

func TestFoo(t *testing.T) {
	echo := exec.Command("go", append(args, "foo")...)
	echoOut, _ := echo.Output()

	if string(echoOut) != "foo\n" {
		t.Fail()
	}
}

func TestJustN(t *testing.T) {
	echo := exec.Command("go", append(args, "-n")...)
	echoOut, _ := echo.Output()

	if string(echoOut) != "" {
		t.Fail()
	}
}

func TestNWithFoo(t *testing.T) {
	echo := exec.Command("go", append(args, "-n", "foo")...)
	echoOut, _ := echo.Output()

	if string(echoOut) != "foo" {
		t.Fail()
	}
}

func TestFoobar(t *testing.T) {
	echo := exec.Command("go", append(args, "foo", "bar")...)
	echoOut, _ := echo.Output()

	if string(echoOut) != "foo bar\n" {
		t.Fail()
	}
}
