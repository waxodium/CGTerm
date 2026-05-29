package main

import (
	"bufio"
	"cgterm/commands"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/fatih/color"
)

func main() {
	if err := ensureVersionFile(); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("File checked and initialized if needed.")
	}
	Firstlaunch()

	reader := bufio.NewReader(os.Stdin)
	signal.Ignore(syscall.SIGTERM, syscall.SIGINT)

	fmt.Println("Welcome to CGTerm! Type 'exit' or press Ctrl+D to quit.")

	for {
		fmt.Print("> ")

		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err.Error())
			break
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

		cmd := exec.Command(name, args...)
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Setsid: true, // new session, so Ctrl+C goes to child's process group
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		if err := cmd.Start(); err != nil {
			fmt.Fprintf(os.Stderr, "%s error: %s\n", color.RedString("[-]"), err)
			continue
		}

		cmd.Wait()
	}
}

func Firstlaunch() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("could not get home directory")
		return
	}
	marker := filepath.Join(home, ".CGTerm_init")

	// if file does NOT exist -> first run
	if _, err := os.Stat(marker); os.IsNotExist(err) {

		fmt.Println(color.CyanString("--FIRST-RUN--"))
		fmt.Println(color.GreenString("This is probably your first time using CGTerm"))
		fmt.Println(color.GreenString("Support this project on github: https://github.com/MasterArd/CGTerm/"))
		fmt.Println(color.GreenString("This message will only show once"))

		// create marker file
		file, err := os.Create(marker)
		if err != nil {
			fmt.Println("could not create marker file:", err)
			return
		}
		file.Close()
	}
}

func ensureVersionFile() error {
	file, err := os.OpenFile("version", os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)

	if err != nil {
		if errors.Is(err, os.ErrExist) {
			return nil 
		}
		return err 
	}
	defer file.Close()
	_, err = file.WriteString("1.3.1")
	return err
}
