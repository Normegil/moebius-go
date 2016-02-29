package input

import "unicode/utf8"

func insertRune(input *Input, r rune) {
	var buf [utf8.UTFMax]byte
	n := utf8.EncodeRune(buf[:], r)
	input.Content = insertByteSlice(input.Content, input.CursorPosition, buf[:n])
}

func insertByteSlice(text []byte, offset int, toInsert []byte) []byte {
	neededCapacity := len(text) + len(toInsert)
	text = increaseSizeOf(text, neededCapacity)
	text = text[:neededCapacity]
	copy(text[offset+len(toInsert):], text[offset:])
	copy(text[offset:], toInsert)
	return text
}

func increaseSizeOf(slice []byte, capacity int) []byte {
	if cap(slice) < capacity {
		newSlice := make([]byte, len(slice), capacity)
		copy(newSlice, slice)
		return newSlice
	}
	return slice
}

func removeBytesFrom(text []byte, from, to int) []byte {
	size := to - from
	copy(text[from:], text[to:])
	text = text[:len(text)-size]
	return text
}
