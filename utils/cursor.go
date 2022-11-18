package utils

import "fmt"

func CursorHide() {
	fmt.Print("\033[?25l")
}

func CursorShow() {
	fmt.Print("\033[?25h")
}
