package reversegivenstring

func ReverseGivenString(item string) string {

	reverse := []rune(item)

	for i, j := 0, len(reverse)-1; i < j; i, j = i+1, j-1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}

	return string(reverse)
}
