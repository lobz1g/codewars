package src

import (
	"fmt"
	"regexp"
	"strings"
)

var MORSE_CODE = map[string]string{"-": "T", "--": "M", "---": "O", "-----": "0", "----.": "9", "---..": "8",
	"---...": ":", "--.": "G", "--.-": "Q", "--..": "Z", "--..--": ",", "--...": "7", "-.": "N", "-.-": "K",
	"-.--": "Y", "-.--.": "(", "-.--.-": ")", "-.-.": "C", "-.-.--": "!", "-.-.-.": ";", "-..": "D", "-..-": "X",
	"-..-.": "/", "-...": "B", "-...-": "=", "-....": "6", "-....-": "-", ".": "E", ".-": "A", ".--": "W", ".---": "J",
	".----": "1", ".----.": "'", ".--.": "P", ".--.-.": "@", ".-.": "R", ".-.-.": "+", ".-.-.-": ".", ".-..": "L",
	".-..-.": "\"", ".-...": "&", "..": "I", "..-": "U", "..---": "2", "..--.-": "_", "..--..": "?", "..-.": "F",
	"...": "S", "...-": "V", "...--": "3", "...---...": "SOS", "...-..-": "$", "....": "H", "....-": "4", ".....": "5"}

func execExpression(str, repl string, exp string) string {
	return regexp.MustCompile(exp).ReplaceAllString(str, repl)
}

func DecodeBits(bits string) string {
	bits = strings.Trim(bits, "0")

	var dotsLen, dashLen int
	for _, v := range regexp.MustCompile("1+").FindAllString(bits, -1) {
		if dotsLen == 0 {
			dotsLen = len(v)
			continue
		}
		if len(v) > dotsLen {
			dashLen = len(v)
		} else if len(v) < dotsLen {
			dashLen = dotsLen
			dotsLen = len(v)
		}
	}

	if dashLen != 0 {
		bits = execExpression(bits, "   ", fmt.Sprintf("0{%d,}", dashLen+1))
		bits = execExpression(bits, "-", fmt.Sprintf("1{%d}", dashLen))
	}
	bits = execExpression(bits, " ", fmt.Sprintf("0{%d,}", dotsLen+1))

	zeroIndex := strings.Index(bits, "0")
	if zeroIndex == -1 {
		zeroIndex = 0
	}

	zerosLen := strings.Index(bits[zeroIndex:], "1")
	if zerosLen == dotsLen && dashLen == 0 {
		bits = execExpression(bits, "", fmt.Sprintf("0{%d}", dotsLen))
	} else {
		bits = execExpression(bits, "2", fmt.Sprintf("0{1,%d}", dotsLen))
	}

	if dashLen == 0 && dotsLen > 2 && strings.Contains(bits, "2") {
		bits = execExpression(bits, "-", fmt.Sprintf("1{%d}", dotsLen))
	} else {
		bits = execExpression(bits, ".", fmt.Sprintf("1{%d}", dotsLen))
	}

	bits = execExpression(bits, "", "2")

	return bits
}

func DecodeMorse(morseCode string) string {
	var result string
	for _, word := range strings.Split(morseCode, "   ") {
		if len(result) != 0 {
			result += " "
		}
		for _, char := range strings.Split(word, " ") {
			result += MORSE_CODE[char]
		}
	}
	return result
}
