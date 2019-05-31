package system

import (
	"fmt"
	"os/exec"
)

func GetTerminalDimension() {
	cmd := exec.Command("echo", "$(tput lines)")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	print(string(stdout))
}
