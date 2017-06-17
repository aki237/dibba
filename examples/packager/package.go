package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aki237/dibba"
)

func main() {
	if len(os.Args) < 3 {
		return
	}
	output := os.Args[1]
	files := os.Args[2:]
	dbFile, err := os.Create(output)
	db := dibba.NewDibba(dbFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, val := range files {
		f, err := os.Open(val)
		if err != nil {
			fmt.Println(err)
			return
		}
		file := dibba.NewFile(filepath.Base(f.Name()), f)
		err = db.Add(file)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = db.Commit()
	if err != nil {
		fmt.Println(err)
		return
	}
}
