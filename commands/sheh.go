package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func Sheh(args []string) {
	fmt.Println(color.CyanString("Thanks and special credit to waxory/waxodium: https://github.com/waxodium/sheh "))
    fmt.Println("")

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
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Printf("%s error: %s\n", color.RedString("[-]"), err)
	}
}

func init() {
	Register("sheh", Sheh)
}
