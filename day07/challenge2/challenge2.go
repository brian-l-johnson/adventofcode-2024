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

func doMath(terms []int64) []int64 {
	var answer []int64
	answer = append(answer, terms[0])
	for _, term := range terms[1:] {
		var newAnswers []int64
		for _, a := range answer {
			newAnswers = append(newAnswers, a+term)
			newAnswers = append(newAnswers, a*term)
			//ac, err := strconv.Atoi(fmt.Sprintf("%d%d", a, term))
			ac, err := strconv.ParseInt(fmt.Sprintf("%d%d", a, term), 10, 64)
			if err != nil {
				panic("could not perform concatination")
			}
			newAnswers = append(newAnswers, int64(ac))
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
			values := make([]int64, len(valuesStrings))
			for i, valueString := range valuesStrings {
				v, err := strconv.Atoi(valueString)
				if err != nil {
					fmt.Println(valueString)
					panic("unable to interpret value as int")
				}
				values[i] = int64(v)
			}
			//fmt.Println(answer, values)
			a := doMath(values)
			//fmt.Println("Answers: ", a)
			if slices.Contains(a, int64(answer)) {
				//fmt.Println(a)
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
