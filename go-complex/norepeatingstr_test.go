package main

import "testing"

/*
   @Auth: menah3m
   @Desc:
*/

func TestGetMaxLengthOfNonRepeatingSubStrByRune(t *testing.T) {
	tests := []struct {
		s string
		i int
	}{
		{"sann", 3},
		{"三四五五四三", 3},
		{"abcabcbb", 3},
		{"", 0},
		{"b", 1},
		{"abcabcabcddsfsdfsadgdsagagdsgas", 5},
	}
	for _, tt := range tests {
		if actual := GetMaxLengthOfNonRepeatingSubStrByRune(tt.s); actual != tt.i {
			t.Errorf("GetMaxLengthOfNonRepeatingSubStrByRune(%s); get %d, expected %d", tt.s, actual, tt.i)
		}
	}
}

// 性能测试
func BenchmarkSubstr(b *testing.B) {
	s := "黑化非法所得了饭撒到发送到发送到发"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("len(s)=%d", len(s))
	b.ResetTimer()
	ans := 12

	for i := 0; i < b.N; i++ {
		if actual := GetMaxLengthOfNonRepeatingSubStrByRune(s); actual != ans {
			b.Errorf("GetMaxLengthOfNonRepeatingSubStrByRune(%s); get %d, expected %d", s, actual, ans)
		}
	}
}
