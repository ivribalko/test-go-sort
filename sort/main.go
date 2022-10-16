package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type iSort interface {
	sort([]int)
}

type bubble struct{}

func (b bubble) sort(list []int) {
	for swapped := true; swapped; {
		swapped = false
		for i := 0; i < len(list) - 1; i++ {
			if list[i] > list[i + 1] {
				list[i], list[i + 1] = list[i + 1], list[i]
				swapped = true
			}
		}
	}
}

func main() {
	csv := os.Args[1]
	strs := strings.Split(csv, ",")
	ints := toInts(strs)

	getSorter().sort(ints)

	fmt.Println(ints)
}

func toInts(strings []string) []int {
	ints := []int{}
	for _, v := range strings {
		j, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, j)
	}
	return ints
}

func getSorter() iSort {
	var sorter iSort
	sorter = bubble{}
	return sorter
}
