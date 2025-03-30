//go:build windows

package main

import (
	"os/exec"
	"syscall"
)

const sh = "cmd"

var ping = []string{"/c", "pause"}

// cmd = exec.Command("cmd.exe", "/C", fmt.Sprintf(`start %s %s`, bin, opt))
func createNewConsole(cmd *exec.Cmd) {
	const CREATE_NEW_CONSOLE = 0x10
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags:    CREATE_NEW_CONSOLE,
		NoInheritHandles: true,
	}
}
