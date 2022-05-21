package reverse_string

func ReverseString(input string) (output string) {
	var sliceOfRunes []int32
	var reversedString string

	for _, rune := range input {
		sliceOfRunes = append(sliceOfRunes, rune)
	}

	for i := len(sliceOfRunes) - 1; i >= 0; i-- {
		reversedString += string(sliceOfRunes[i])
	}

	return reversedString
}
