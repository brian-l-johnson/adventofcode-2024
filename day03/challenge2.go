package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	//read input from stdin into an array of strings
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	//fmt.Println("lines:", lines)
	//regular expression to match valid instructions
	//mul(\d{1,3}, \d{1,3})
	//do()
	//don't()
	var re = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don\'t\(\)`)
	var mul = true
	var sum = int64(0)

	for _, line := range lines {
		fmt.Println("line:", line)
		//match the instruction
		var match = re.FindAllStringSubmatch(line, -1)
		//fmt.Println("match:", match)
		for _, m := range match {
			fmt.Println("match: ", m)
			if m[0] == "do()" {
				mul = true
			} else if m[0] == "don't()" {
				mul = false
			} else {
				if mul {
					a, err := strconv.ParseInt(m[1], 10, 64)
					if err != nil {
						fmt.Println("error:", err)
						return
					}
					b, err := strconv.ParseInt(m[2], 10, 64)
					if err != nil {
						fmt.Println("error:", err)
						return
					}
					sum += a * b
				}
			}
		}
	}
	fmt.Println("sum:", sum)
}
