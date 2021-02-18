package main

import (
	"fmt"

	"github.com/gjlmotea/go-pickers/number"
)

func main() {
	g := number.NewPicker()
	g.InhabitLeadingZero = true
	g.AllowDigits = []byte{'0'}
	g.Length = 30

	fmt.Println(g.Pick())
	fmt.Println(g.Pick())
	fmt.Println(g.Pick())
	fmt.Println(g.Pick())
	fmt.Println(g.Pick())
	fmt.Println(g.Pick())
	fmt.Println(g.Pick())
	fmt.Println(g.Pick())
	fmt.Println(g.Pick())
	fmt.Println(g.Pick())

	n := number.NewPicker()
	n.Length = 30

	fmt.Println(n.Pick())
	fmt.Println(n.Pick())
	fmt.Println(n.Pick())
	fmt.Println(n.Pick())
	fmt.Println(n.Pick())
}
