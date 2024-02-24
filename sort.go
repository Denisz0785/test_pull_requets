package main

import (
	"fmt"
	"sort"
)

type author struct {
	name string
	age  int
}

func main() {
	author_slice := []author{{"Bill", 1967}, {"Sofia", 1985}, {"Alexandr", 1985}, {"Feona", 1957}}
	fmt.Println(author_slice)

	sort.Slice(author_slice, func(i, j int) bool {
		if author_slice[i].age != author_slice[j].age {
			return author_slice[i].age < author_slice[j].age

		} else if len(author_slice[i].name) != len(author_slice[j].name) {
			return len(author_slice[i].name) < len(author_slice[j].name)

		}

		return author_slice[i].name < author_slice[j].name

	})

	fmt.Println(author_slice)

}
