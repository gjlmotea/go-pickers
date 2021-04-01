package number

import (
	"math/rand"
	"time"
)

var DefaultDigits = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

type Picker struct {
	// 指定數字長度，預設長度為1，因可設超長長度，所以型態為string而非int
	Length int
	// 允許出現的Digit，預設允許全部
	Dictionary []byte
	// 禁止出現前導0
	InhabitLeadingZero bool
}

// NewPicker returns a new Picker
func NewPicker() (n Picker) {
	rand.Seed(time.Now().UnixNano())
	n.Length = 1
	n.Dictionary = DefaultDigits
	return n
}

// Pick pick and return a random number.
// If Picker.Dictionary is nil, it return empty string.
func (n *Picker) Pick() (number string) {
	if n.Dictionary == nil {
		return
	}

	var bytes []byte
	for i := 0; i < n.Length; i++ {
		r := rand.Int() % len(n.Dictionary)
		d := n.Dictionary[r]

		// 如果index為第一項，並且 禁止前導0，需額外檢查AllowDigits是否只有'0'
		if index := indexOf(n.Dictionary, '0'); i == 0 && n.InhabitLeadingZero && index != -1 {
			if len(n.Dictionary) > 1 {
				r = rand.Int() % (len(n.Dictionary) - 1)
				s := remove(n.Dictionary, index)
				d = s[r]

			} else {
				return "`Picker.InhabitLeadingZero` is enabled, but `Picker.Dictionary` only contains '0'!"
			}
		}

		bytes = append(bytes, d)
	}

	return string(bytes)
}

func indexOf(s []byte, e byte) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}

func remove(s []byte, i int) []byte {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
