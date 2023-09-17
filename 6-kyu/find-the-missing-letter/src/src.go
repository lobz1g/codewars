package src

func FindMissingLetter(chars []rune) rune {
	if chars[0]-chars[1] != 1 {
		return chars[0] + 1
	}
	return FindMissingLetter(chars[1:])
}
