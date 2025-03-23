package main

import "fmt"

//go:generate go build -o canton/gen/cantongen canton/gen/cantongen.go
//go:generate canton/gen/cantongen a b c

func main() {
	fmt.Printf("main\n")
}
