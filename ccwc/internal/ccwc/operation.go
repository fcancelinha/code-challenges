package ccwc

import (
	"fmt"
)

func ProcessOperation(flag string, b []byte) error {
	op := map[string]OperationFunc{
		"c": ByteCount(),
		"l": LineCount(),
		"w": WordCount(),
		"m": CharCount(),
	}

	if fn, ok := op[flag]; ok {
		c := fn(b)
		fmt.Printf("%d ", c)
	}

	return nil
}
