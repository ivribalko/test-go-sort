package main

type bubble struct{}

// Keep swapping adjacent until no swaps.
func (bubble) sort(list []int) {
	for swapped := true; swapped; {
		swapped = false
		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = true
			}
		}
	}
}
