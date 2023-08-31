package main

import "fmt"

func main() {
	// append to slice
	strings0 := []string{"a", "b", "c"}
	dump(strings0)

	strings := append(strings0, "d") // cap 3 -> cap 6 (new slice)
	strings0[0] = "a_changed"
	dump(strings) // no changes

	strings = append(strings, "e", "f")
	dump(strings)

	more := []string{"g", "h"}
	strings = append(strings, more...)
	dump(strings)

	fmt.Println()
	dump(strings[0:1])
	dump(strings[0:2])
	dump(strings[0:3])

	fmt.Println()
	dump(strings[1:2])
	dump(strings[1:3])
	dump(strings[1:4])

	fmt.Println()

	// three-index slices
	nums := []string{"aa", "bb", "cc", "dd", "ee", "ff"}
	dump(nums)
	numsSlice := nums[0:4]
	dump(numsSlice)

	nums3i := nums[0:4:4]
	dump(nums3i)
	nums3iExtra := append(nums3i, "qq") // new slice, nums is not changed
	dump(nums3iExtra)

	numsSlice2 := append(numsSlice, "qq2")
	dump(numsSlice2)
	dump(numsSlice)
}

func dump(slice []string) {
	fmt.Printf("%v len: %v, cap: %v \n", slice, len(slice), cap(slice))
}
