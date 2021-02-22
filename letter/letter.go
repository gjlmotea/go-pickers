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

var DefaultAplhabet = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
var Lowercase = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
var Uppercase = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

type Picker struct {
	// 全大寫、全小寫、字首大寫、字首小寫、大寫交錯、小寫交錯
	Case int
	// 指定數字長度，預設長度為1，因可設超長長度，所以型態為string而非int
	Length int
	// 允許出現的Digit，預設允許全部
	AllowLetters []byte
}

// NewPicker returns a new Picker
func NewPicker() (p Picker) {
	rand.Seed(time.Now().UnixNano())
	p.Length = 1
	p.AllowLetters = append(Lowercase, Uppercase...)
	return p
}

// Pick pick and return a random number.
// If Picker.AllowDigits is nil, it return empty string.
func (n *Picker) Pick() (number string) {
	if n.AllowLetters == nil {
		return
	}

	var bytes []byte
	for i := 0; i < n.Length; i++ {
		r := rand.Int() % len(n.AllowLetters)
		d := n.AllowLetters[r]
		bytes = append(bytes, d)
	}

	return string(bytes)
}

// PickUp pick and return a random number.
// The only difference between PickUp and Pick is:
// PickUp will remember the value which had returned,
// and keep in a map structure to compare with the next PickUp value.
// It means that the same value won't be returned by PickUp twice or more,
// Just like that something(or choice) was picked up, and never put it back,
// until use func PutAllBack.
func (n *Picker) PickUp() (number string) {
	return
}

func PutAllBack() {

}
