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

type selection struct{}

func (bubble) sort(list []int) {
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

func (selection) sort(list []int) {
	for i, v := range list {
		j := i
		for k := i + 1; k < len(list); k++ {
			if list[k] < v {
				j = k
			}
		}
		if (j != i) {
			list[j], list[i] = list[i], list[j]
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
	if len(os.Args) > 2 {
		switch os.Args[2] {
		case "bubble":
			return bubble{}
		case "selection":
			return selection{}
		default:
			log.Fatal("unknown sorting algorythm")
		}
	}

	return bubble{}
}
