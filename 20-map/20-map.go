package main

import (
	"fmt"
	"math"
)

func main() {
	// map[key_type]value_type

	temp := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	fmt.Println(temp)

	t1 := temp["Earth"]
	fmt.Println(t1)

	temp["Earth"] = 16
	fmt.Println(temp)

	// new key
	temp["Moon"] = -20
	fmt.Println(temp)

	// wrong key -> 0
	t2 := temp["WRONG_KEY"]
	fmt.Println(t2)

	// The "comma ok" idiom
	if moonT, ok := temp["Moon"]; ok {
		fmt.Println(ok)
		fmt.Printf("Temp: %v *C\n", moonT)
	}

	if unknownT, found := temp["Unknown"]; found {
		fmt.Println("Not unknown", unknownT, found)
	} else {
		fmt.Println("Unknown", unknownT, found)
	}

	// maps are not copied when assigned
	temp2 := temp
	temp2["Moon"] = -21
	fmt.Println(temp["Moon"] == temp2["Moon"])

	// delete
	delete(temp, "Moon")
	fmt.Println(temp, temp2)

	// make map
	temp0 := make(map[string]int, 2)
	fmt.Println(temp0, len(temp0)) // 0
	temp0["a"] = 1
	temp0["b"] = 2
	temp0["c"] = 3
	fmt.Println(temp0)

	// count by map
	chars := []string{"a", "b", "c", "a", "b", "a"}
	freqs := make(map[string]int)
	for _, char := range chars {
		freqs[char]++
	}
	for key, val := range freqs {
		fmt.Printf("%v - %d times\n", key, val)
	}

	// grouping
	temps := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	groups := make(map[float64][]float64)
	for _, t := range temps {
		g := math.Trunc(t/10) * 10
		groups[g] = append(groups[g], t)
	}
	for group, groupTemps := range groups {
		fmt.Printf("%v: %v\n", group, groupTemps)
	}
}
