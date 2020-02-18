package iteration

import "fmt"

// Repeat repeats the provided character the given amount of times
func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

// ExampleRepeat provides an example for executing Repeat
func ExampleRepeat() {
	sum := Repeat("a", 5)
	fmt.Println(sum)
	// Output: aaaaa
}
