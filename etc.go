//go:build !windows

package main

import (
	"os/exec"
)

const sh = "bash"

var ping = []string{"-c", "ping 127.0.0.1"}

func createNewConsole(*exec.Cmd) {

}
