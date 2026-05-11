package commands

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/fatih/color"
)

func Sheh(args []string) {
	fmt.Println(color.BlueString("If Sheh is not working as expected. CGTerm may need to install correctly"))
	fmt.Println(color.CyanString("\nCheck: CGTerm manual install guide https://github.com/MasterArd/CGTerm/#README"))
	fmt.Println(" ")
	fmt.Println(color.CyanString("Thanks and special credit to waxory/waxodium: https://github.com/waxodium/ "))

	_, err := exec.LookPath("sheh")
	if err != nil {
		var shehNotInstalled string
		fmt.Print(color.RedString("[-]"), " sheh is not found on your system. Install it? (y/n): ")
		fmt.Scanln(&shehNotInstalled)
		if strings.ToLower(shehNotInstalled) == "y" {
			fmt.Println(color.BlueString("[*]"), "Installing sheh via npm...")
			installCmd := exec.Command("npm", "install", "-g", "@waxory/sheh")
			installCmd.Stdout = os.Stdout
			installCmd.Stderr = os.Stderr
			if err := installCmd.Run(); err != nil {
				fmt.Println(color.RedString("[-]"), "npm failed. Is it installed and in your PATH?")
				return
			}
			fmt.Println(color.GreenString("[+]"), "Installation complete.")
		} else {
			return
		}
	}

	cmd := exec.Command("sheh", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = nil
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}

	if err := cmd.Start(); err != nil {
		fmt.Println(color.RedString("[-]"), "Execution error:", err)
		return
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		s := <-sig
		cmd.Process.Signal(s)
	}()

	fmt.Println("-->")
}

func init() {
	Register("sheh", Sheh)
}
