package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	//read input from stdin into an array of strings
	var lines []string
	for {
		var line string
		fmt.Scan(&line)
		if len(line) == 0 {
			break
		}
		lines = append(lines, line)
	}
	//regular expression that matches mul(\d{1,3},\d{1,3})
	var re = regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	sum := int64(0)
	//iterate over each line in the = 0
	for _, line := range lines {
		//match each occurance of the regular expression against the line
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			//convert matched numbers to integers and multiply them
			a, err := strconv.ParseInt(match[1], 10, 64)
			if err != nil {
				fmt.Println("Error parsing a: ", err)
				return
			}
			b, err := strconv.ParseInt(match[2], 10, 64)
			if err != nil {
				fmt.Println("Error parsing b: ", err)
				return
			}
			sum += a * b
		}
	}
	fmt.Println("Sum:  ", sum)
}
