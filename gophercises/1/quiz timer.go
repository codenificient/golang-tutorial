package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to Gophercises!")

	timer := time.NewTimer(50 * time.Second)
	done := make(chan bool)
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
	go func() {
		<-timer.C
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
		done <- true
	}()
	<-done
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
