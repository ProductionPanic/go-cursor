package cursor

import "fmt"

type Cursor struct {
}

func (c *Cursor) Up(n int) {
	fmt.Printf("\033[%dA", n)
}
func (c *Cursor) Down(n int) {
	fmt.Printf("\033[%dB", n)
}
func (c *Cursor) Right(n int) {
	fmt.Printf("\033[%dC", n)
}
func (c *Cursor) Left(n int) {
	fmt.Printf("\033[%dD", n)
}
func (c *Cursor) Top() {
	fmt.Printf("\033[%dH", 0)
}
func (c *Cursor) Bottom() {
	fmt.Printf("\033[%dH", 9999)
}
func (c *Cursor) LineStart() {
	fmt.Printf("\033[%dG", 0)
}
func (c *Cursor) LineEnd() {
	fmt.Printf("\033[%dG", 9999)
}
func (c *Cursor) ClearLine() {
	fmt.Printf("\033[2K\r")
}
func (c *Cursor) ClearScreen() {
	fmt.Printf("\033[2J")
}
func (c *Cursor) SavePosition() {
	fmt.Printf("\033[s")
}
func (c *Cursor) RestorePosition() {
	fmt.Printf("\033[u")
}
func (c *Cursor) Hide() {
	fmt.Printf("\033[?25l")
}
func (c *Cursor) Show() {
	fmt.Printf("\033[?25h")
}

func GetCursor() *Cursor {
	return &Cursor{}
}
