package main

import (
	"log"
	"strconv"
)

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
