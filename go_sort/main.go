package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	csv := flag.String(
		"csv",
		"2,1,3",
		"comma separated integers")

	alg := flag.String(
		"algorithm",
		"bubble",
		"sorting algorithm")

	flag.Parse()

	fmt.Println(sorter(*alg).sort(sortee(*csv)))
}

func sortee(csv string) []int {
	return toInts(strings.Split(csv, ","))
}

func sorter(alg string) iSort {
	switch alg {
	case "bubble":
		return bubble{}
	case "selection":
		return selection{}
	default:
		panic("unknown sorting algorithm")
	}
}
