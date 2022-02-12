package main

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) == len(s2) && transfor(s1)^transfor(s2) == 0 {
		return true
	}
	return false
}

func transfor(s string) (res rune) {
	for _, ch := range s {
		bit := ch - 96
		res += 1 << bit
	}
	return res
}
