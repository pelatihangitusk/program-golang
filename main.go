package main

import (
	"bytes"
	"fmt"
)

func splitDemo() {
	fmt.Println("Split demo");
	a := []byte("a,b,c")
	b := bytes.Split(a, []byte(","))
	fmt.Printf("a before changing b[0][0]: %q\n", a)
	fmt.Printf("%q\n", b);

	b[0][0] = byte('*')
	fmt.Printf("a after changing b[0][0]:  %q\n", a)

	fmt.Printf("b[1] before appending to b[0]: %q\n", b[1])
	b[0] = append(b[0], 'd', 'e', 'f')
	fmt.Printf("b[1] after appending to b[0]:  %q\n", b[1])
	fmt.Printf("a after appending to b[0]: %q\n", a)
}

func appendDemo() {
	fmt.Println("\nAppend demo")
	s1 := make([]int, 2, 4)
	s1[0] = 1
	s1[1] = 2
	fmt.Printf("Initial address and value: %p: %[1]v\n", s1)
	s1 = append(s1, 3, 4)
	fmt.Printf("After first append:        %p: %[1]v\n", s1)
	s1 = append(s1, 5)
	fmt.Printf("After second append:       %p: %[1]v\n", s1)
}

func alwaysCopy() {
	fmt.Println("\nAppend and always copy")
	s1 := []int{1, 2, 3, 4}
	fmt.Printf("s1: %p: %[1]v\n", s1)
	s2 := make([]int, 4, 8)
	copy(s2, s1)
	fmt.Printf("s2: %p: %[1]v\n", s2)
	s2 = append(s2, 5, 6, 7, 8)
	fmt.Printf("s2: %p: %[1]v\n", s2)
}

func main() {
	splitDemo()
	// appendDemo()
	// alwaysCopy()
}
