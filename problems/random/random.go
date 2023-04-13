package random

import "fmt"

func RunRandom() {

	recursive_strings := []string{"aa2[hello3[world]]", "new1[project]"}

	fmt.Println("Recursive string")
	for _, v := range recursive_strings {
		fmt.Println(v, RecursiveString(v))
	}

}
