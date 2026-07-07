// package main

// import (
// 	"fmt"
// 	"os"
// )

// type KVPair struct {
// 	key   string
// 	value string
// }

// func main() {
// 	path := "data/Moscow.txt"
// 	input, err := os.ReadFile(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	map_out := Map(path, string(input))
// 	fmt.Println(map_out)

// 	res := Reduce("2024", []string{"12.1", "-10.5", "32.1", "30.6"})
// 	fmt.Println(res)
// }

package main

import (
	"fmt"
	"os"
	"sync"
)

type KVPair struct {
	key   string
	value string
}

func main() {
	cities := []string{"Tokyo", "Delhi", "Shanghai", "Sao_Paulo", "Mexico_City", "Cairo", "Mumbai", "Beijing", "Dhaka", "Osaka", "New_York", "Karachi", "Buenos_Aires", "Istanbul", "Kolkata", "Lagos", "Moscow", "London", "Paris", "Los_Angeles"}

	//Reads in input file for each city and calls Map function
	ch := make(chan KVPair)
	var wg sync.WaitGroup
	for _, city := range cities {
		wg.Add(1)
		go func(city string) {
			defer wg.Done()
			path := "data/" + city + ".txt"
			input, err := os.ReadFile(path)
			if err != nil {
				panic(err)
			}
			map_out := Map(path, string(input))
			for _, item := range map_out {
				ch <- item
			}
		}(city)
	}

	// Goroutine waits to close the channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	// range over channel repeatedly reads from channel until it is closed
	kv_pairs := make(map[string][]string)
	fmt.Println(kv_pairs)

	for item := range ch {
		// TODO: correctly populate the map kv_pairs with the items read in on the channel "ch"
	}
}

// TODO: add calling Reduce tasks
