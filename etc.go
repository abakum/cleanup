//go:build !windows

package main

import "os/exec"

const sh = "bash"

func createNewConsole(*exec.Cmd) {

}
