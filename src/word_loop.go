package main

import "fmt"

func longestPalindrome(s string) string {
	if s == "" {
		return s
	}
	a := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		a[i] = make([]int, len(s))
		a[i][i] = 1
	}
	for i := 0; i < len(s)-1; i++ {
		if []byte(s)[i] == []byte(s)[i+1] {
			a[i][i+1] = 1
		}
	}

	for j := 1; j < len(s); j++ {
		for i := j - 1; i >= 0; i-- {
			if []byte(s)[i] == []byte(s)[j] && a[i+1][j-1] == 1 {
				a[i][j] = 1
			}
		}
	}
	max := 0
	ii := 0
	jj := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if a[i][j] == 1 && j-i > max {
				max = j - i
				ii = i
				jj = j
			}
		}

	}
	return string([]byte(s)[ii : jj+1])
}

func main()  {
	s := longestPalindrome("aasssddffddddddddddddffffffffdffffffffffffffffffffffffffdsfdddddddddddddddd")
	fmt.Println(s)
}