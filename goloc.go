package goloc

func SkipToTheEndOfLine(dat string, datlen int, currpos int) int {
	for i := currpos; i < datlen; i++ {
		if !(i-1 >= 0 && dat[i-1] == '\\' && dat[i] == '\n') && dat[i] == '\n' {
			return i - currpos
		}
	}
	return datlen - currpos
}

func SkipMultilineComment(dat string, datlen int, currpos int) int {
	for i := currpos; i < datlen; i++ {
		if dat[i] == '*' && i+1 < datlen && dat[i+1] == '/' {
			return (i + 2) - currpos
		}
	}
	return datlen - currpos
}

// Counts line of code without comments
func LocWithoutComments(dat string) int {
	loc := 0
	datlen := len(dat)

	if datlen <= 0 {
		return 0
	}

	for i := 0; i < datlen; i++ {
		c := dat[i]

		// SKIP ALL spaces and tabs and empty lines before code
		if c == ' ' || c == '\t' || c == '\n' {
			continue
		}

		if c == '/' && i+1 < datlen && dat[i+1] == '/' {
			i += SkipToTheEndOfLine(dat, datlen, i)
		} else if c == '/' && i+1 < datlen && dat[i+1] == '*' {
			i += SkipMultilineComment(dat, datlen, i)
		} else {
			i += SkipToTheEndOfLine(dat, datlen, i)
			loc++
		}
	}
	return loc
}
