// To execute test suite: "go test -v"

package validationkit

import (
	"log"
	"math/rand"
	"regexp"
	"time"
)

const UsernameRegex string = `^@?(\w){1,15}$`

func CheckUsernameSyntax(username string) bool {
	validationResult := false
	r, err := regexp.Compile(UsernameRegex)
	if err != nil {
		log.Fatal(err)
	}
	validationResult = r.MatchString(username)
	return validationResult
}

func GenerateRandomUsername() string {
	rand.Seed(time.Now().UnixNano())

	usernameLength := rand.Intn(15) + 1

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_")
	b := make([]rune, usernameLength)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	randomUsername := string(b)

	zeroOrOne := rand.Intn(2)
	if zeroOrOne == 1 {
		randomUsername = "@" + randomUsername
	}
	return randomUsername
}
