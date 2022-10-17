package main

import (
	"log"
	"strconv"
)

// Array of strings to array of ints.
func toInts(strings []string) []int {
	ints := make([]int, len(strings))
	for i, v := range strings {
		j, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		ints[i] = j
	}
	return ints
}
