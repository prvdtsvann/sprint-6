package service

import (
	"log"

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
	log.Printf("Input: %s, isMorse: %v", input, isMorse)

	if isMorse {
		result := morse.ToText(input)
		log.Printf("Converted to text: %s", result)
		return result
	}

	result := morse.ToMorse(input)
	log.Printf("Converted to Morse: %s", result)
	return result
}
