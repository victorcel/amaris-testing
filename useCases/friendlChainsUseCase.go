package useCases

import (
	"errors"
	"fmt"
)

const (
	TEXTSTRINGLENGTH  = "text strings must be same length"
	THEYAREFRIENDS    = "Las cadenas '%s' y '%s' son amigas"
	THEYNOTAREFRIENDS = "Las cadenas '%s' y '%s' no son amigas"
)

func FriendlyChains(x, y string) (response string, err error) {

	if len(x) != len(y) {
		return "", errors.New(TEXTSTRINGLENGTH)
	}

	for index := 1; index < len(x); index++ {

		u := x[index:]

		v := x[:index]

		concatXY := fmt.Sprintf("%s%s", u, v)

		if y == concatXY {
			return fmt.Sprintf(THEYAREFRIENDS, x, y), nil
		}
	}

	return fmt.Sprintf(THEYNOTAREFRIENDS, x, y), nil
}
