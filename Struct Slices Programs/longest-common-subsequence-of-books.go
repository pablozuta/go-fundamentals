package main

import (
	"fmt"
	"sort"
)

type HorrorBuk struct {
	Title  string
	Author string
}

func main() {
	horrorBooks := []HorrorBuk{
		{Author: "Stephen King", Title: "It"},
		{Author: "William Peter Blatty", Title: "The Exorcist"},
		{Author: "Stephen King", Title: "The Shining"},
		{Author: "Shirley Jackson", Title: "The Haunting of Hill House"},
		{Author: "Josh Malerman", Title: "Bird Box"},
		{Author: "Mark Z. Danielewski", Title: "House of Leaves"},
		{Author: "Stephen King", Title: "The Stand"},
		{Author: "Shirley Jackson", Title: "We Have Always Lived in the Castle"},
		{Author: "Stephen King", Title: "The Outsider"},
	}
	booksByAuthor := make(map[string][]string)
	for _, book := range horrorBooks {
		booksByAuthor[book.Author] = append(booksByAuthor[book.Author], book.Title)
	}

	var longestCommonSubsequences []string
	for _, titles := range booksByAuthor {
		sort.Strings(titles)
		if len(titles) > 1 {
			longestCommonSubsequence := findLongestCommonSubsequence(titles[0], titles[1:]...)
			longestCommonSubsequences = append(longestCommonSubsequences, longestCommonSubsequence)
		}
	}
	fmt.Println("Longest common subsequences of book titles for each author:")
	for i, lcs := range longestCommonSubsequences {
		fmt.Printf("%s: %s\n", horrorBooks[i*3].Author, lcs)
	}
}
func findLongestCommonSubsequence(s1 string, s2 ...string) string {
	var lcs string
	for i := range s1 {
		for _, s := range s2 {
			for j := range s {
				if s[j] == s1[i] {
					if len(lcs) == 0 {
						lcs += string(s1[i])
					} else if i > 0 && j > 0 {
						lcs = lcs[:len(lcs)-1] + string(s1[i])
					} else {
						lcs += string(s1[i])
					}
					break
				}
			}
		}
	}
	return lcs
}
