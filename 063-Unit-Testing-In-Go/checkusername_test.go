// To execute test suite: "go test -v"

package validationkit

import (
	"fmt"
	"testing"
)

// Test for minimum number of characters
func TestCheckUsernameSyntaxMinimumCharacterLength(t *testing.T) {
	result := CheckUsernameSyntax("")

	if result != false {
		t.Errorf("Failed the minimum character check.")
	}
}

// Test for maximum number of characters
func TestCheckUsernameSyntaxMaximumCharacterLength(t *testing.T) {
	result := CheckUsernameSyntax("aaaaaaaaaaaaaaaa")

	if result != false {
		t.Errorf("Failed the maximum character check.")
	}
}

func TestCheckUsernameSyntaxSymbols(t *testing.T) {
	result := CheckUsernameSyntax("g-@p!h#r")
	if result != false {
		t.Errorf("Failed the special character check.")
	}
}

func TestCheckUsernameSyntaxUnderscore(t *testing.T) {
	// Underscores are permissible
	result := CheckUsernameSyntax("the_gopher")
	if result != true {
		t.Errorf("Failed the check to allow underscore characters.")
	}
}

func TestCheckUsernameSyntaxAtSignInsideUsername(t *testing.T) {
	// The @ sign can only be placed at the start of the username string and is invalid anywhere else
	result := CheckUsernameSyntax("the@gopher")
	if result != false {
		t.Errorf("Failed the @ sign check. The @ sign was found in another place besides the start of the string.")
	}
}

func TestCheckUsernameSyntaxRandomUsernames(t *testing.T) {
	for i := 0; i < 10008; i++ {
		username := GenerateRandomUsername()
		//fmt.Println("username: ", username)
		result := CheckUsernameSyntax(username)
		if result != true {
			res := fmt.Sprintf("The random username '%v' failed to pass the username check.", username)
			t.Errorf(res)
			t.Fatal("Quitting!")
		}
	}
}
