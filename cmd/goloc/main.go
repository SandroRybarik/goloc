package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/SandroRybarik/goloc"
)

func usage() {
	fmt.Println("Usage goloc:\ngoloc <input_file>")
}

func main() {
	if len(os.Args) != 2 {
		usage()
	} else {
		dat, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		fmt.Println(goloc.LocWithoutComments(string(dat)))
	}
}
