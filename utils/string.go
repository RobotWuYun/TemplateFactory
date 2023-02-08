package utils

import (
	"strings"
	"unicode"
)

func Utf8Index(str, substr string) int {
	asciiPos := strings.Index(str, substr)
	if asciiPos == -1 || asciiPos == 0 {
		return asciiPos
	}
	pos := 0
	totalSize := 0
	reader := strings.NewReader(str)
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		totalSize += size
		pos++
		// 匹配到
		if totalSize == asciiPos {
			return pos
		}
	}
	return pos
}

func StringHasUpper(word string) (hasUpper bool) {
	return false
	for _, r := range word {
		if unicode.IsUpper(r) {
			hasUpper = true
			break
		}
	}
	return
}
