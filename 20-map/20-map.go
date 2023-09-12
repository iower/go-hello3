package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
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

	// Sets by maps
	set := make(map[float64]bool)
	for _, t := range temps {
		set[t] = true
	}
	if set[-28.0] {
		fmt.Println("is in the set")
	}
	fmt.Println(set)
	unique := make([]float64, 0, len(set))
	for t := range set {
		unique = append(unique, t)
	}
	sort.Float64s(unique)
	fmt.Println(unique)

	// freq
	text := "As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear — except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man."
	textMap := countWords(text)
	for key, freq := range textMap {
		if freq > 1 {
			fmt.Printf("%v - %v\n", freq, key)
		}
	}
}

func countWords(text string) map[string]int {
	words := strings.Fields(strings.ToLower(text))
	fmt.Println(words)
	textMap := make(map[string]int, len(words))
	for _, word := range words {
		word := strings.Trim(word, ",.;—")
		textMap[word] += 1
	}
	return textMap
}
