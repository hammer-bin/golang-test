package main

type key struct {
	v int
}

type value struct {
	v int
}

func main() {
	var m map[string]string
	var m1 map[int]string
	var m2 map[key]value
}
