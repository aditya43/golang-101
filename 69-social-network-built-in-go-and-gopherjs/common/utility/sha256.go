package utility

import (
	"crypto/sha256"
	"fmt"
)

func SHA256OfString(input string) string {

	sum := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x", sum)

}
