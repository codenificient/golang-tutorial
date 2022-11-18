package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func main() {
	fmt.Println("Welcome to Gophercises!")

	file, err := os.Open("1/problems.csv")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	defer file.Close()

	rd := csv.NewReader(file)
	lines, err := rd.ReadAll()
	if err != nil {
		log.Fatal("Error reading CSV:", err)
	}
	problems := parseLines(lines)
	correct := 0
	for idx, problm := range problems {
		fmt.Printf("Problem #%d: %s = \n", idx+1, problm.question)
		var answer string
		fmt.Scanf("%s", &answer)
		if answer == problm.answer {
			correct++
		} else {
			correct--
		}
	}
	fmt.Printf("You got %d correct out of %d\n", correct, len(lines))
}

func parseLines(lines [][]string) []problem {
	retn := make([]problem, len(lines))
	for i, line := range lines {
		retn[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return retn
}
