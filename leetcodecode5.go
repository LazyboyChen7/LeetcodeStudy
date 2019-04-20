package main

import (
	"fmt"
)

func manachar(s string) string {
	t := "$#"
	for _,v := range s {
		t += string(v)
		t += "#"
	}
	fmt.Println(t)
	arr := make([]int, len(t))
	var mx, idx, resLen, resCenter int
	for i:=1; i<len(t); i++ {
		if mx > i {
			if arr[2*idx - i] > mx-i {
				arr[i] = mx - i
			} else {
				arr[i] = arr[2*idx - i]
			}
		} else {
			arr[i] = 1
		}
		for i+arr[i] < len(t) && t[i+arr[i]] == t[i-arr[i]] {
			arr[i]++
		}
		if mx < i + arr[i] {
			mx = i + arr[i]
			idx = i
		}
		if resLen < arr[i] {
			resLen = arr[i]
			resCenter = i
		}
	}
	start := (resCenter - resLen)/2
	return s[start : start + resLen-1]
}

func main() {
	str1 := "abcacbcbcacbacd"
	str2 := manachar(str1)
	fmt.Printf("原始字符串: %s\n最长回文串: %s\n", str1, str2)
}