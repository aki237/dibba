package main

import (
	"fmt"
	"io"
	"os"

	"github.com/aki237/dibba"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprint(os.Stderr, "Program can only take 3 arguments in.\n")
		return
	}
	inputFile := os.Args[1]
	file := os.Args[2]
	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	d := dibba.NewDibbaReader(f)
	err = d.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}
	fd, err := d.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(os.Stdout, fd.GetReader())
	if err != nil {
		fmt.Println(err)
		return
	}
}
