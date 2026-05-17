package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func Firstlaunch() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("could not get home directory")
		return
	}
	marker := filepath.Join(home, ".CGTerm_init")

	// if file does NOT exist -> first run
	if _, err := os.Stat(marker); os.IsNotExist(err) {

		fmt.Println(color.RedString("WARNING:"))
		fmt.Println(color.RedString("Sudo command is currently bugged and may crash."))
		fmt.Println(color.RedString("Avoid using it until fixed."))

		// create marker file
		file, err := os.Create(marker)
		if err != nil {
			fmt.Println("could not create marker file:", err)
			return
		}
		file.Close()
	}
}
