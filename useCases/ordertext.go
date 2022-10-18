package useCases

import (
	"errors"
	"sort"
	"strings"
)

func OrderText(text string) (response []string, err error) {
	invalidCharacters := ";:.'!@#$<>?[{]}\"'"

	if strings.ContainsAny(text, invalidCharacters) {
		return nil, errors.New("invalid character Example: \"victor,elias\"")
	}

	response = strings.Split(text, ",")

	sort.Slice(response, func(i, j int) bool {
		return response[i] < response[j]
	})

	return response, nil
}
