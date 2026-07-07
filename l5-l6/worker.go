package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Map(key string, value string) []KVPair {

	output := make([]KVPair, 0)

	// TODO: This loop iterates over each line of the "value" string
	// You will want to parse out the date and temperature from each line and add it to the "output" slice
	for _, line := range strings.Split(strings.TrimSuffix(value, "\n"), "\n") {
		fmt.Println(line)
		slist := strings.Split(line, ",")    //split the line by commas to seperate city, date and temperature
		date := strings.Split(slist[1], "-") //split the date from the list

		year := date[0]  //get only the year
		temp := slist[2] //get only the temp

		item := KVPair{year, temp}    //make a keyvalue pair with year and temp
		output = append(output, item) //append to the output

	}

	return output
}

func Reduce(key string, value []string) float64 {
	// Converting from a string to float may be useful
	var max float64
	for _, x := range value { //iterate over the values
		val, err := strconv.ParseFloat(x, 64)
		if err != nil {
			fmt.Println("Error during conversion:", err)
		}
		if val > max {
			max = val
		}
	}
	return max
}
