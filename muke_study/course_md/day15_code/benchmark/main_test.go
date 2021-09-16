package main

import "testing"

func IsPalindrome(str string) {
	_ = (str)
}
func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan. a canal: Panama")
	}
}
