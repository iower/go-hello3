package main

import "fmt"

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

}
