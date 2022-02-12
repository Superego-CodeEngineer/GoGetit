package main

func isUnique(astr string) bool {
	var result rune
	result = 0
	for _, ch := range astr {
		bit := ch - 97
		if result&(1<<bit) != 0 {
			return false
		}
		result |= 1 << bit
	}
	return true
}
