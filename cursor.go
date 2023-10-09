package cursor

import "fmt"

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
