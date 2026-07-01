// Every executable Go Program should contain a package called main.
// This tells the Go compiler to compile the package into an executable
// program rather than a shared library.
package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func searchForWord(filepath string, target string) {
	dat, err := os.ReadFile(filepath)
	check(err)
	input := string(dat)
	length := len(target) //get the word length of target word

	i := 1
	for i+length < len(input) {
		search := input[i : i+length] //looking through the string in sections of length
		if search == target {         // see if the current section of is the target word
			fmt.Printf("found "+target+"@ %v\n", i)
			i = i + length + 1 //increase by length to go over the word
		} else {
			i = i + 1 //cat not found, go to the very next letter
		}

	}
}

// The entry point of a Go program should be the main function of main package.
// When the executable is run, main() is automatically called.
func main() {
	fmt.Println("Hello World\n")
	searchForWord("/Users/mariagonzalez/Documents/cs-work/CS351/class-exercises/l1-l2/dictionary.txt", "fish")

	//input := "There once was a cat named Barry. He was a very good cat. This cat lived in Boston. He loved doing Boston-related activities (that were good for cats). He walked the esplanade. He shopped on Newbury. He ate at Tatte. He sometimes even went to TD Garden. Did you know that cats are not allowed in TD Garden?"

	// dat, err := os.ReadFile("/Users/mariagonzalez/Documents/cs-work/CS351/class-exercises/l1-l2/dictionary.txt")
	// check(err)
	// input := string(dat)

	// cat := "cat" //cat string to find
	// i := 1
	// for i+3 < len(input) {
	// 	search := input[i : i+3] //looking through the string in sections of 3
	// 	if search == cat {       // see if the current section of is the word cat
	// 		fmt.Printf("found cat @ %v\n", i)
	// 		i = i + 4 //increase by 4 to go over the word cat
	// 	} else {
	// 		i = i + 1 //cat not found, go to the very next letter
	// 	}

	// }
}
