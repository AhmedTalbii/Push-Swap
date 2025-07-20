package instructions

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	errClear := cmd.Run()
	if errClear != nil {
		fmt.Println("error clearing the terminal : ", errClear)
	}
}
