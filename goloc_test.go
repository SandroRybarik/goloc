package main

import "testing"

func TestLocWithoutComments(t *testing.T) {
	input := `/**   
		 * Class X is doing only x.
		 */   
		class X {
			// method returns x
			method() {
				const x = 1;
				return x;
			}
		}
		/**
		 * End
		 */
		 // Comment
		 // Comment2


		 // Comment3
	`
	ans := LocWithoutComments(input)

	if ans != 6 {
		t.Errorf("LocWithoutComments(input) = %d; want 6", ans)
	}
}
