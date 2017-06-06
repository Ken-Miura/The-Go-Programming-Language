// Copyright 2017 Ken Miura
package ex05

func RemoveAdjacentDuplication(strings []string) []string {
	before := strings[0]
	for i := 1; i < len(strings); {
		if before != strings[i] {
			before = strings[i]
			i++
			continue
		}
		strings = remove(strings, i)
	}
	return strings
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
