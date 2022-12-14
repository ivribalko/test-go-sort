package main

type selection struct{}

// For all items find lowest among next items.
func (selection) sort(list []int) []int {
	for i, v := range list {
		j := i
		for k := i + 1; k < len(list); k++ {
			if list[k] < v {
				j = k
			}
		}
		if j != i {
			list[j], list[i] = list[i], list[j]
		}
	}
	return list
}
