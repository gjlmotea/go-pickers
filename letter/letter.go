package letter

import (
	"math/rand"
	"time"
)

const (
	CaseInsensitive        = iota // 不區分大小寫
	CaseUpper                     // `APPLE` 全部大寫
	CaseLower                     // `apple` 全部小寫
	CaseTitle                     // `Apple` 字首大寫，其餘小寫
	CaseTitleReverse              // `aPPLE` 字首小寫，其餘大寫
	CaseAlternating               // `aPpLe` 小寫大寫交錯
	CaseAlternatingReverse        // `ApPlE` 大寫小寫交錯
)

type Picker struct {
	// 全大寫、全小寫、字首大寫、字首小寫、大寫交錯、小寫交錯
	Case string
	// 指定數字長度，預設長度為1，因可設超長長度，所以型態為string而非int
	Length int
	// 允許出現的Digit，預設允許全部
	AllowDigits []byte
	// 禁止出現前導0
	InhabitLeadingZero bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewPicker returns a new Picker
func NewPicker() (p Picker) {
	p.Length = 1
	p.AllowDigits = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	return p
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

		// 如果index為第一項，並且 禁止前導0，需額外檢查AllowDigits是否只有'0'
		if index := indexOf(n.AllowDigits, '0'); i == 0 && n.InhabitLeadingZero && index != -1 {
			if len(n.AllowDigits) > 1 {
				r = rand.Int() % (len(n.AllowDigits) - 1)
				s := remove(n.AllowDigits, index)
				d = s[r]

			} else {
				return "`Picker.InhabitLeadingZero` is enabled, but `Picker.AllowDigits` only contains '0'!"
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
