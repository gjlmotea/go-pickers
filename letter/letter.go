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

var DefaultLetters = []byte{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
}

type Picker struct {
	// 全大寫、全小寫、字首大寫、字首小寫、大寫交錯、小寫交錯
	Case int
	// 指定數字長度，預設長度為1，因可設超長長度，所以型態為string而非int
	Length int
	// 允許出現的Letter，預設允許全部
	Dictionary []byte
}

// NewPicker returns a new Picker
func NewPicker() (l Picker) {
	rand.Seed(time.Now().UnixNano())
	l.Length = 1
	l.Dictionary = DefaultLetters
	return l
}

// Pick pick and return a random number.
// If Picker.Dictionary is nil, it return empty string.
func (l *Picker) Pick() (number string) {
	bytes := l.pick()
	switch l.Case {
	case CaseInsensitive:
		return string(bytes)
	case CaseUpper:
		for i := 0; i < len(bytes); i++ {
			if bytes[i] >= 'a' && bytes[i] <= 'z' {
				bytes[i] -= 32
			}
		}
		return string(bytes)
	case CaseLower:
		for i := 0; i < len(bytes); i++ {
			if bytes[i] >= 'A' && bytes[i] <= 'Z' {
				bytes[i] += 32
			}
		}
		return string(bytes)
	case CaseTitle:
		if len(bytes) > 0 && bytes[0] >= 'a' && bytes[0] <= 'z' {
			bytes[0] -= 32
		}
		for i := 1; i < len(bytes); i++ {
			if bytes[i] >= 'A' && bytes[i] <= 'Z' {
				bytes[i] += 32
			}
		}
		return string(bytes)
	case CaseTitleReverse:
		if len(bytes) > 0 && bytes[0] >= 'A' && bytes[0] <= 'Z' {
			bytes[0] += 32
		}
		for i := 1; i < len(bytes); i++ {
			if bytes[i] >= 'a' && bytes[i] <= 'z' {
				bytes[i] -= 32
			}
		}
		return string(bytes)
	case CaseAlternating:
		for i := 0; i < len(bytes); i++ {
			if i%2 == 0 && bytes[i] >= 'A' && bytes[i] <= 'Z' {
				bytes[i] += 32
			}
			if i%2 == 1 && bytes[i] >= 'a' && bytes[i] <= 'z' {
				bytes[i] -= 32
			}
		}
		return string(bytes)
	case CaseAlternatingReverse:
		for i := 0; i < len(bytes); i++ {
			if i%2 == 0 && bytes[i] >= 'a' && bytes[i] <= 'z' {
				bytes[i] -= 32
			}
			if i%2 == 1 && bytes[i] >= 'A' && bytes[i] <= 'Z' {
				bytes[i] += 32
			}
		}
		return string(bytes)
	default:
		return "Error"
	}
}

func (l *Picker) pick() (bytes []byte) {
	if l.Dictionary == nil {
		return
	}

	for i := 0; i < l.Length; i++ {
		r := rand.Int() % len(l.Dictionary)
		d := l.Dictionary[r]
		bytes = append(bytes, d)
	}

	return
}

// PickUp pick and return a unique random number.
// The only difference between PickUp and Pick is:
// PickUp will remember the value which had returned,
// and keep in a map structure to compare with the next PickUp value.
// It means that the same value won't be returned by PickUp twice or more,
// Just like that something(or choice) was picked up, and never put it back,
// until use func PutAllBack.
// TODO
func (l *Picker) PickUp() (number string) {
	return
}

// TODO
func PutAllBack() {

}
