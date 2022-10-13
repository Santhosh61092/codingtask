package transform

import (
	"regexp"
	"strings"
)
type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}

type skipString struct {
	originalText    string
	transformedRune []rune
}

func (ctx skipString) String() string {
	return string(ctx.transformedRune)
}

func NewSkipString(pos int, text string) skipString {
	instance := skipString{
		originalText: text,
	}
	instance.TransformRune(pos)
	return instance
}

func (ctx *skipString) TransformRune(pos int) {
	var str string
	for i, c := range strings.Split(ctx.originalText, "") {

		if i != 0 && (i+1)%3 == 0 {
			str = str + strings.ToUpper(c)
		} else {
			str = str + strings.ToLower(c)
		}
	}
	ctx.transformedRune = []rune(str)
}

func (ctx skipString) GetValueAsRuneSlice() []rune {
	return ctx.transformedRune
}

func CapitalizeEveryThirdAlphanumericChar(s string) string {
	var str string = strings.ToLower(s[0:2])

	re := regexp.MustCompile(`(\S{3})`)
	for _, charArray := range re.FindAllString(strings.ToLower(s)[2:], -1) {
		str = str + strings.Title(charArray)
	}
	return str
}