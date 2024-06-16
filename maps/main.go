package main

import (
	"fmt"
	"strings"
)

func Maps_demo() {
	m := make(map[string]int)

	m["Hello"] = 1

	fmt.Println(m["Hello"])

	delete(m, "Hello")

	fmt.Println(m)

	v, ok := m["Hello"]

	fmt.Printf("Value %d and presence %t\n", v, ok)

	m["Hello"] = 100

	u, ok := m["Hello"]

	fmt.Printf("Value %d and presence %t\n", u, ok)

	fmt.Println(maps_exercise("I am learning Go!"))
}

func maps_exercise(s string) map[string]int {
	m := make(map[string]int)

	splits := strings.Fields(s)

	for _, v := range splits {
		_, ok := m[v]

		if ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}

	return m
}
