package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func doMath(terms []int) []int {
	var answer []int
	answer = append(answer, 0)
	for _, term := range terms {
		var newAnswers []int
		for _, a := range answer {
			newAnswers = append(newAnswers, a+term)
			newAnswers = append(newAnswers, a*term)
		}
		answer = newAnswers
	}
	return answer
}

func main() {
	fmt.Println("Day 7 Challenge 1")

	scanner := bufio.NewScanner(os.Stdin)

	line_re := regexp.MustCompile(`^(\d)+: ((\d+) )+(\d+)$`)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line_re.MatchString(line) {
			fmt.Println("Line recognized: ", line)
			equation_parts := strings.Split(line, ": ")
			answerString := equation_parts[0]
			answer, err := strconv.Atoi(answerString)
			if err != nil {
				panic("unable to interpret answer as int")
			}
			valuesStrings := strings.Split(equation_parts[1], " ")
			//fmt.Println(valuesStrings)
			values := make([]int, len(valuesStrings))
			for i, valueString := range valuesStrings {
				v, err := strconv.Atoi(valueString)
				if err != nil {
					fmt.Println(valueString)
					panic("unable to interpret value as int")
				}
				values[i] = v
			}
			//fmt.Println(answer, values)
			a := doMath(values)
			//fmt.Println("Answers: ", a)
			if slices.Contains(a, answer) {
				fmt.Println("Answer found: ", answer)
				sum += answer
			}

		} else {
			fmt.Println("Line not recognized")
			panic("Invalid line")
		}
		//fmt.Println(line)
	}
	fmt.Println("Sum: ", sum)
}
