package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	csv := os.Args[1]
	strs := strings.Split(csv, ",")
	ints := toInts(strs)

	getSorter().sort(ints)

	fmt.Println(ints)
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
