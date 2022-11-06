package src

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

func execExpression(str, repl string, exp string) string {
	return regexp.MustCompile(exp).ReplaceAllString(str, repl)
}

func getMinAndMaxLen(bits string, char string) (int, int) {
	min := math.MaxInt
	var max int
	for _, v := range regexp.MustCompile(fmt.Sprintf("%s+", char)).FindAllString(bits, -1) {
		if len(v)-1 > max || max == 0 {
			max = len(v)
		}
		if len(v) < min {
			min = len(v)
		}
	}
	return min, max
}

func replaceForDot(bits string, dotsLen int) string {
	for i := 0; i < len(bits) && strings.Contains(bits, "1"); i++ {
		if bits[i] != '1' {
			continue
		}

		j := strings.IndexAny(bits[i:], "0-. ")
		if j == -1 || i == len(bits)-dotsLen {
			if len(bits[i:]) == dotsLen {
				bits = bits[:i] + "."
			}
			break
		} else if j == dotsLen {
			tmp := bits
			bits = tmp[:i] + "." + tmp[j+i:]
		} else {
			i += j
		}
	}

	return bits
}

func DecodeBitsAdvanced(bits string) string {
	println(bits)

	bits = strings.Trim(bits, "0")

	var firstDash int
	for strings.Contains(bits, "1") {
		_, shortSpaceLen := getMinAndMaxLen(bits, "0")
		dots, dashes := getMinAndMaxLen(bits, "1")
		if dashes != dots && dashes != 0 {
			if firstDash == 0 {
				firstDash = dashes
			} else if firstDash/2 >= dashes {
				bits = execExpression(bits, ".", fmt.Sprintf("1{%d,}", dots))
			}
			bits = replaceForDot(bits, dots)
			bits = execExpression(bits, "-", fmt.Sprintf("1{%d,}", dashes))
		} else {
			if dots > shortSpaceLen+1 && shortSpaceLen != 0 {
				bits = execExpression(bits, "-", "1+")
			} else {
				bits = execExpression(bits, ".", "1+")
			}
			dashes = 0
		}

		checker, longSpaceLen := getMinAndMaxLen(bits, "0")
		if (checker != longSpaceLen || longSpaceLen > 5) && longSpaceLen > dots+2 && !strings.Contains(bits, "   ") {
			if dashes < longSpaceLen && !strings.Contains(bits, "   ") {
				longSpaceLen = dashes + 3
			}
			bits = execExpression(bits, "   ", fmt.Sprintf("0{%d,}", longSpaceLen))
		}

		_, shortSpaceLen = getMinAndMaxLen(bits, "0")
		if shortSpaceLen != 0 && (longSpaceLen > dots || (dots == 1 && dashes == 0 && shortSpaceLen != 1)) {
			if dashes <= shortSpaceLen && firstDash/2+1 > shortSpaceLen {
				shortSpaceLen = dashes + 1
			} else if dots+1 >= shortSpaceLen {
				shortSpaceLen = dashes
			}
			bits = execExpression(bits, " ", fmt.Sprintf("0{%d,}", shortSpaceLen))
		}
	}

	bits = execExpression(bits, "", "0")

	return strings.Trim(bits, " ")
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

var MORSE_CODE = map[string]string{"-": "T", "--": "M", "---": "O", "-----": "0", "----.": "9", "---..": "8", "---...": ":", "--.": "G", "--.-": "Q", "--..": "Z", "--..--": ",", "--...": "7", "-.": "N", "-.-": "K", "-.--": "Y", "-.--.": "(", "-.--.-": ")", "-.-.": "C", "-.-.--": "!", "-.-.-.": ";", "-..": "D", "-..-": "X", "-..-.": "/", "-...": "B", "-...-": "=", "-....": "6", "-....-": "-", ".": "E", ".-": "A", ".--": "W", ".---": "J", ".----": "1", ".----.": "'", ".--.": "P", ".--.-.": "@", ".-.": "R", ".-.-.": "+", ".-.-.-": ".", ".-..": "L", ".-..-.": "\"", ".-...": "&", "..": "I", "..-": "U", "..---": "2", "..--.-": "_", "..--..": "?", "..-.": "F", "...": "S", "...-": "V", "...--": "3", "...---...": "SOS", "...-..-": "$", "....": "H", "....-": "4", ".....": "5"}
