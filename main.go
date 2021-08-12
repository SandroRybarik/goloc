package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Returns number of LOC skipping comments (//, /**/)
func LocWithoutComments(dat string) int {

	locCount := 0
	isMulti := false
	untilNewLine := false
	datlen := len(dat)

	for i := 0; i < datlen; i++ {
		if dat[i] == ' ' || dat[i] == '\t' {
			continue
		}

		if untilNewLine && dat[i] == '\n' {
			// SKIP LINE UNTIL '\n'
			untilNewLine = false
		} else if dat[i] == '/' && i+1 < datlen && dat[i+1] == '/' {
			// SINGLE LINE COMMENT
			untilNewLine = true
		} else if dat[i] == '/' && i+1 < datlen && dat[i+1] == '*' {
			// MULTI_LINE_COMMENT_START
			isMulti = true
			untilNewLine = true
		} else if dat[i] == '*' && i+1 < datlen && dat[i+1] == '/' {
			// MULTI_LINE_COMMENT_END
			isMulti = false
			untilNewLine = true
		} else if (dat[i] == '\n' && i+1 < datlen && dat[i+1] == '\n') || (i-1 >= 0 && dat[i-1] == '\n' && dat[i] == '\n') {
			// SKIP IMMIDIATE EMPTY LINES
		} else if !isMulti && dat[i] == '\n' {
			locCount++
		}
	}

	return locCount
}

func LocWithoutComments2(dat string) int {
	locCount := 0
	isMulti := false
	isSingle := false
	// untilNewLine := false
	isContent := false
	datlen := len(dat)

	for i := 0; i < datlen; i++ {
		if isSingle && dat[i] == '\n' {
			isSingle = false
		}
		if dat[i] == '/' && i+1 < datlen && dat[i+1] == '/' {
			// SINGLE LINE COMMENT
			isSingle = true
		} else if dat[i] == '/' && i+1 < datlen && dat[i+1] == '*' {
			// MULTI_LINE_COMMENT_START
			isMulti = true
		} else if dat[i] == '*' && i+1 < datlen && dat[i+1] == '/' {
			// MULTI_LINE_COMMENT_END
			isMulti = false
		} else if !(dat[i] == ' ' || dat[i] == '\n' || dat[i] == '\t') {
			isContent = true
			// line should be counted
		} else if !isMulti && !isSingle && isContent && dat[i] == '\n' {
			locCount++
		}
	}

	return locCount
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid number of arguments.")
	} else {
		dat, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		fmt.Println(LocWithoutComments(string(dat)))
	}
}
