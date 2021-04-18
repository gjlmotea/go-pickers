package main

import (
	"github.com/gjlmotea/go-pickers/letter"
	"github.com/gjlmotea/go-pickers/number"
)

type Pickers struct {
	Letter letter.Picker
	Number number.Picker
}