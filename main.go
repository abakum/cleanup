// go mod init github.com/abakum/cleanup
// go mod tidy
package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/unixist/go-ps"
	"github.com/xlab/closer"
)

const (
	EL      = "\033[K"    // очистить строку
	DECTCEM = "\033[?25h" // показать курсор
)

func main() {
	log.SetPrefix("\r")
	log.SetFlags(log.Lshortfile)

	ctx, cancel := context.WithCancel(context.Background())
	defer closer.Close()
	cleanup := func() {
		log.Println("cleanup")
		<-ctx.Done()
		KidsDone(os.Getpid())
		log.Println("cleanup done" + DECTCEM + EL) // показать курсор, очистить строку
	}
	closer.Bind(cleanup)
	closer.Bind(cancel)

	shContext := exec.CommandContext(ctx, sh)
	createNewConsole(shContext)
	err := shContext.Start()
	log.Println(shContext, "context sh start", err)
	if err == nil {
		if shContext.Process != nil {
			log.Println(shContext, "context sh pid", shContext.Process.Pid)
		}
		go func() {
			log.Println(shContext, "context sh exit", shContext.Wait())
			closer.Close()
		}()
	}

	shell := exec.Command(sh)
	shell.Stdin = os.Stdin
	shell.Stdout = os.Stdout

	time.AfterFunc(time.Second*7, func() {
		closer.Close()
	})

	err = shell.Start()
	log.Println(shell, "sh start", err)
	log.Println("To exit, press <^C> or exit in any shell, otherwise exit will occur in 7 seconds")
	if err == nil {
		if shell.Process != nil {
			log.Println(shell, "sh pid", shell.Process.Pid)
		}
		log.Println("sh exit", shell, shell.Wait())
	} else {
		closer.Hold()
	}
}

// Завершает дочерние процессы
func KidsDone(ppid int) {
	if ppid < 1 {
		return
	}
	pes, err := ps.Processes()
	if err != nil {
		return
	}
	for _, p := range pes {
		if p == nil {
			continue
		}
		if p.PPid() == ppid && p.Pid() != ppid {
			PidDone(p.Pid())
		}
	}
}

// Завершает процесс с pid
func PidDone(pid int) {
	Process, err := os.FindProcess(pid)
	if err == nil {
		log.Println("pid", pid, "done", Process.Kill())
		return
	}
	log.Println("pid", pid, err)
}
