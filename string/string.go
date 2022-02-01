package string

import (
	"fmt"
	"time"
)

type StringType int

const (
	Ascii StringType = iota
)

func NewStringSlice(typ StringType, sliceLen int, charLen int) []string {
	if sliceLen < 0 || charLen <= 0 {
		panic("Err: sliceLen or carLen must be positive value.")
	}
	arr := make([]string, sliceLen)
	s := uint(time.Now().UnixMilli())
	for i := 0; i < sliceLen; i++ {
		var str string
		for j := 0; j < charLen; j++ {
			var charByte uint
			switch typ {
			case Ascii:
				charByte = (uint(s) % 25) + 97
			default:
				panic("Err: not implemented yet.")
			}
			str += fmt.Sprintf("%c", charByte)
			s = randomizeNumber(s)
		}
		arr[i] = str
	}
	return arr
}
func NewAsciiStringWithCharLenSlice(sliceLen int, charLen int) []string {
	return NewStringSlice(Ascii, sliceLen, charLen)
}
func NewAsciiStringSlice(sliceLen int) []string {
	return NewAsciiStringWithCharLenSlice(sliceLen, 3)
}

func randomizeNumber(a uint) uint {
	return (a>>32)*a + (1<<62 - 1)
}
