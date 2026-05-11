package main

import (
	"CGTerm/commands"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"github.com/fatih/color"
)

func main() {
	syscall.Setsid()
	fd := int(os.Stdin.Fd())
	syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TIOCSCTTY), 0)
	reader := bufio.NewReader(os.Stdin)
	

    fmt.Println("Welcome to CGTerm! Type 'exit' to quit;")
	

    for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		
        if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}
		

        input = strings.TrimSpace(input)
		if input == "" {
			continue
		}



        parts := strings.Split(input, " ")
		name := parts[0]
		args := parts[1:]

		cmd, exists := commands.Registry[name]
		if !exists {
			c := exec.Command(name, args...)
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			c.Stdin = os.Stdin
			

            if err := c.Run(); err != nil {
				fmt.Printf("%s error: %s\n", color.RedString("[-]"), err)
			}
			
            continue
		

        }

		cmd(args)
	}
}
