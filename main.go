package main

import (
	"fmt"

	"github.com/gjlmotea/go-pickers/letter"
	"github.com/gjlmotea/go-pickers/number"
)

func main() {
	p1 := number.NewPicker()
	p1.Length = 12
	p1.InhabitLeadingZero = true
	fmt.Println(p1.Pick()) // 143492824677

	code := number.NewPicker()
	code.Dictionary = []byte{'0', '1'}
	code.Length = 20
	fmt.Println(code.Pick()) // 01010000000111010110

	p2 := letter.NewPicker()
	p2.Length = 8
	p2.Case = letter.CaseTitle
	fmt.Println(p2.Pick()) // Ipwzeigk

	gene := letter.NewPicker()
	gene.Length = 25
	gene.Dictionary = []byte{'A', 'T', 'C', 'G'}
	fmt.Println(gene.Pick()) //CGAGAACCCACTGTAGAGGAGGCTT
}
