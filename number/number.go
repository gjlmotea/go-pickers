package number

import (
	"math/rand"
	"time"
)

type Picker struct {
	// 指定數字長度，預設長度為1，因可設超長長度，所以型態為string而非int
	Length int
	// 允許出現的Digit，預設允許全部
	AllowDigits []byte
	// TODO 允許出現前導0
	IsLeadingZero bool
}

// NewPicker returns a new Picker
func NewPicker() (p Picker) {
	p.Length = 1
	p.AllowDigits = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	return p
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Pick pick and return a random number.
// If Picker.AllowDigits is nil, it return empty string.
func (n *Picker) Pick() (number string) {
	if n.AllowDigits == nil {
		return
	}

	var bytes []byte
	for i := 0; i < n.Length; i++ {
		r := rand.Int() % len(n.AllowDigits)
		d := n.AllowDigits[r]
		bytes = append(bytes, d)
	}

	return string(bytes)
}
