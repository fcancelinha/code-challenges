package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/fcancelinha/code-challenge-ccwc/internal/ccwc"
)

func main() {
	var arg string

	help := flag.Bool("h", false, "Show usage")
	flag.Bool("c", false, "Output the number of bytes")
	flag.Bool("l", false, "Output the number of lines")
	flag.Bool("w", false, "Output the number of words")
	flag.Bool("m", false, "Output the number of characters")
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	var input io.Reader

	if len(flag.Args()) > 0 {
		arg = flag.Args()[0]

		file, err := os.Open(arg)
		if err != nil {
			fmt.Println(err)
			return
		}

		input = file
	} else {
		input = os.Stdin
	}

	var buffer bytes.Buffer

	_, err := io.Copy(&buffer, input)
	if err != nil && err != io.EOF {
		fmt.Println(err)
	}

	b := buffer.Bytes()

	p := func(f *flag.Flag) {
		if f.Name != "h" {
			err := ccwc.ProcessOperation(f.Name, b)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	if flag.NFlag() > 0 {
		flag.Visit(p)
	} else {
		flag.VisitAll(p)
	}

	fmt.Print(arg + "\n")
}
