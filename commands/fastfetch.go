package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func Fastfetch(args []string) {
	cmd := exec.Command("fastfetch", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err, "is it installed?")
	}
}

func init() {
	Register("fastfetch", Fastfetch)
}
