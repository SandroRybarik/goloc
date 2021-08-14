package goloc

import (
	"testing"
)

func TestLocCodeWithSingleLineComment(t *testing.T) {
	input := `// Comment
const x = () => {}; // Another inline
// Another One
const y = 10;
`
	ans := LocWithoutComments(input)
	if ans != 2 {
		t.Errorf("LocWithoutComments(input) = %d; want 2", ans)
	}
}

func TestLocCodeWithMultilineComment(t *testing.T) {
	input := `/*
	* Comment */
const x = () => {}; /* Another inline */
// Another One /* nested one */
/**
 * y is constant
 */
const y = 10;
`
	ans := LocWithoutComments(input)
	if ans != 2 {
		t.Errorf("LocWithoutComments(input) = %d; want 2", ans)
	}
}

func TestLocCodeWithSkippedLines(t *testing.T) {
	input := `/*
	* Comment */

const x = () => {}; /* Another inline */
// Another One /* nested one */

const y = 11;   

/**
 * y is constant
 */
const z = 10;





`
	ans := LocWithoutComments(input)
	if ans != 3 {
		t.Errorf("LocWithoutComments(input) = %d; want 3", ans)
	}
}

func TestLocCodeWithSingleLineCommentEscaped(t *testing.T) {
	input := `// Single line comment\
that continues through 2 lines
const x = "123";
`

	ans := LocWithoutComments(input)
	if ans != 1 {
		t.Errorf("LocWithoutComments(input) = %d; want 1", ans)
	}
}

// Skip to the end of line (Steo)
func TestSkipToTheEndOfLine(t *testing.T) {
	input1 := `1234567890`

	ans1 := SkipToTheEndOfLine(input1, len(input1), 0)
	if ans1 != 10 {
		t.Errorf("Steol(input1) = %d; want 10", ans1)
	}

	input2 := `1234567890
hey`

	ans2 := SkipToTheEndOfLine(input2, len(input2), 0)
	if ans2 != 10 {
		t.Errorf("Steol(input2) = %d; want 10", ans2)
	}
}
