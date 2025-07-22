package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(input string) string {
	if input == "" {
		return ""
	}

	isMorse := true
	for _, v := range input {
		if v != '.' && v != '-' && v != ' ' {
			isMorse = false
			break
		}
	}
	if isMorse {
		return morse.ToText(input)
	}
	return morse.ToMorse(input)
}
