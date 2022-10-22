package main

import "fmt"

func main() {
	r := letterCombinations("23")

	fmt.Println(r)
}

func letterCombinations(digits string) []string {
	var result []string
	if len(digits) == 0 {
		return result
	}

	dict := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9' : "wxyz",
	}

	var dfs func(i int, cur []byte)
	dfs = func(i int, cur []byte) {
		if i == len(digits) {
			result = append(result, string(cur))
			return
		}

		search := dict[digits[i]]
		for idx:=0;idx<len(search);idx++{
			cur = append(cur, search[idx])
			dfs(i+1, cur)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(0, []byte{})

	return result
}
