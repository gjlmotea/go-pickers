package main

import (
	"fmt"

	"github.com/gjlmotea/go-pickers/letter"
)

func main() {
	g := letter.NewPicker()
	g.Length = 5
	g.Case = letter.CaseAlternating
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
}
