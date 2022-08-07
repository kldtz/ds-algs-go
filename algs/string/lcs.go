package string

import "strings"

func LcsNaive(s1 string, s2 string) []string {
	maxLen := 0
	lcs := []string{}
	for i := 0; i < len(s1); i += 1 {
		for j := i + 1; j <= len(s1); j += 1 {
			substring := s1[i:j]
			if strings.Contains(s2, substring) {
				if len(substring) > maxLen {
					maxLen = len(substring)
					lcs = []string{substring}
				} else if len(substring) == maxLen {
					lcs = append(lcs, substring)
				}
			}
		}
	}
	return lcs
}

func LcsDynamicProgramming(s1 string, s2 string) []string {
	maxLen := 0
	lastCharIdx := []int{}
	counts := make([][]int, len(s1))
	for i := range counts {
		counts[i] = make([]int, len(s2))
	}
	for i, c1 := range s1 {
		for j, c2 := range s2 {
			if c1 == c2 {
				curr := 1
				if i != 0 && j != 0 {
					curr = counts[i-1][j-1] + 1
				}
				counts[i][j] = curr
				if curr > maxLen {
					maxLen = curr
					lastCharIdx = []int{i}
				} else if curr == maxLen {
					lastCharIdx = append(lastCharIdx, i)
				}
			}
		}
	}
	lcs := make([]string, 0, len(lastCharIdx))
	for _, i := range lastCharIdx {
		lcs = append(lcs, s1[i-maxLen+1:i+1])
	}
	return lcs
}
