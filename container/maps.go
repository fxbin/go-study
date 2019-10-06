package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"name":   "ccmouse",
		"course": "golang",
		"site":   "hhh",
		"cc":     "dd",
	}
	fmt.Println(m)

	// m2 == empty map
	m2 := make(map[string]int)
	fmt.Println(m2)

	// m3 == nil
	var m3 map[string]int
	fmt.Println(m3)

	fmt.Println("traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	if cause, ok := m["cause_"]; ok {
		fmt.Println(cause)
	} else {
		fmt.Println("key dose not exist")
	}

	fmt.Println("Deleting values")
	name1, ok := m["name"]
	fmt.Println(name1, ok)
}
