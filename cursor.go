package cursor

import (
	"fmt"
	"strconv"
	"strings"
)

func Up(n int) {
	fmt.Printf("\033[%dA", n)
}
func Down(n int) {
	fmt.Printf("\033[%dB", n)
}
func Right(n int) {
	fmt.Printf("\033[%dC", n)
}
func Left(n int) {
	fmt.Printf("\033[%dD", n)
}
func Top() {
	fmt.Printf("\033[%dH", 0)
}
func Bottom() {
	fmt.Printf("\033[%dH", 9999)
}
func LineStart() {
	fmt.Printf("\033[%dG", 0)
}
func LineEnd() {
	fmt.Printf("\033[%dG", 9999)
}
func ClearLine() {
	fmt.Printf("\033[2K\r")
}
func ClearScreen() {
	fmt.Printf("\033[2J")
}
func SavePosition() {
	fmt.Printf("\033[s")
}
func RestorePosition() {
	fmt.Printf("\033[u")
}
func Hide() {
	fmt.Printf("\033[?25l")
}
func Show() {
	fmt.Printf("\033[?25h")
}
func Exec(cmd string) {
	lines := strings.Split(cmd, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		hasArg := len(parts) > 1
		var arg string
		if hasArg {
			arg = parts[1]
		}
		switch parts[0] {
		case "up":
			if hasArg {
				argInt, err := strconv.Atoi(arg)
				if err == nil {
					Up(argInt)
				}
			} else {
				Up(1)
			}
		case "down":
			if hasArg {
				argInt, err := strconv.Atoi(arg)
				if err == nil {
					Down(argInt)
				}
			} else {
				Down(1)
			}
		case "left":
			if hasArg {
				argInt, err := strconv.Atoi(arg)
				if err == nil {
					Left(argInt)
				}
			} else {
				Left(1)
			}
		case "right":
			if hasArg {
				argInt, err := strconv.Atoi(arg)
				if err == nil {
					Right(argInt)
				}
			} else {
				Right(1)
			}
		case "top":
			Top()
		case "bottom":
			Bottom()
		case "line-start":
			LineStart()
		case "line-end":
			LineEnd()
		case "clear-line":
			ClearLine()
		case "clear-screen":
			ClearScreen()
		case "save-position":
			SavePosition()
		case "restore-position":
			RestorePosition()
		case "hide":
			Hide()
		case "show":
			Show()
		}
	}
}
