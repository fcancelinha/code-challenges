package ccwc

import (
	"bytes"
	"unicode/utf8"
)

type OperationFunc func(b []byte) int

func ByteCount() OperationFunc {
	return func(b []byte) int {
		c := len(b)
		return c
	}
}

func LineCount() OperationFunc {
	return func(b []byte) int {
		l := bytes.Count(b, []byte{'\n'})
		return l
	}
}

func WordCount() OperationFunc {
	return func(b []byte) int {
		f := bytes.Fields(b)
		return len(f)
	}
}

func CharCount() OperationFunc {
	return func(b []byte) int {
		c := utf8.RuneCount(b)
		return c
	}
}
