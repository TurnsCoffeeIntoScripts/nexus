package system

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func GetTerminalDimension() (uint, uint) {
	heightStr := execCommand("tput", "lines")
	widthStr := execCommand("tput", "cols")

	height, errH := strconv.Atoi(heightStr)
	width, errW := strconv.Atoi(widthStr)

	if errH != nil {
		height = 0
		fmt.Println("Could not extract terminal height")
	}

	if errW != nil {
		width = 0
		fmt.Println("Could not extract terminal width")
	}

	return uint(height), uint(width)
}

func execCommand(cmdName string, params string) string {
	cmd := exec.Command(cmdName, params)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return strings.TrimSuffix(string(stdout), "\n")
}
