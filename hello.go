package main

import "fmt"

const (
	isAdmin = 1 << iota
	isHQ
	isEnna
	isTobi
)

func main() {
	role := isAdmin | isEnna

	fmt.Printf("isEnna %v \n", isEnna&role == isEnna)
	fmt.Printf("isTobi %v \n", isTobi&role == isTobi)
}
