package main

import (
	"errors"
	"sort"
)

type person struct {
	id int
}

func distinct(persons []person) ([]person, error) {
	if len(persons) == 0 {
		return nil, errors.New("no elements to check for duplicates")
	}

	sorted := ascendingSort(persons)
	result := []person{}
	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i].id != sorted[i+1].id {
			result = append(result, sorted[i])
		}
	}
	result = append(result, sorted[len(sorted)-1])

	return result, nil
}

func ascendingSort(structs []person) []person {
	sort.Slice(structs, func(i, j int) bool {
		return structs[i].id < structs[j].id
	})

	return structs
}
