package main

import "fmt"

func restoreString(s string, indices []int) string {
	var temp []string
	var rst string

	for _, v := range indices {
		fmt.Println(v)
		temp = append(temp, s[v:v+1])
		fmt.Println(temp)
		rst += s[v : v+1]
	}

	/*m := make(map[int]string)
	for _, v := range indices {
		m[v] = s[v : v+1]
	}
	for i := 0; i < len(m); i++ {
		fmt.Println(m[i])
	}*/

	result := make([]byte, len(indices))
	for i, v := range indices {
		fmt.Println(result)
		result[v] = s[i]
	}
	fmt.Println(result)
	fmt.Println(string(result))

	return rst
}

func main() {
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	s := "codeleet"

	rst := restoreString(s, indices)
	fmt.Println(rst)
}
