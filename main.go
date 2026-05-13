package main

import (
	"CGTerm/commands"
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/fatih/color"
	"golang.org/x/sys/unix"
)

func main() {
	fd := int(os.Stdin.Fd())

	unix.Setpgid(0, 0)
	signal.Ignore(syscall.SIGTTOU)
	signal.Ignore(syscall.SIGTTIN)
    signal.Ignore(syscall.SIGINT)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to CGTerm! Type 'exit' or press Ctrl+D to quit.")

	for {
		fmt.Print("> ")

		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}

		input := strings.TrimSpace(line)
		if input == "" {
			continue
		}
		if input == "exit" {
			break
		}

		parts := strings.Fields(input)
		name := parts[0]
		args := parts[1:]

		cmdFunc, exists := commands.Registry[name]
		if exists {
			cmdFunc(args)
			continue
		}

		c := exec.Command(name, args...)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Stdin = os.Stdin
		c.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

		if err := c.Start(); err != nil {
			fmt.Fprintf(os.Stderr, "%s error: %s\n", color.RedString("[-]"), err)
			continue
		}

		unix.IoctlSetInt(fd, unix.TIOCSPGRP, c.Process.Pid)
		c.Wait()
		unix.IoctlSetInt(fd, unix.TIOCSPGRP, unix.Getpgrp())
	}
}
