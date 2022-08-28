package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "csv file in format question,answer")

	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		log.Fatalf("Could not open csv file -> %v", csvFile)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Could not parse csv file -> %s", err)
	}
	questions := parseLines(lines)

	correct := 0
	for i, q := range questions {
		fmt.Printf("Question #%d: %s = \n", i+1, q.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == q.a {
			correct++
		}
	}

	fmt.Printf("You got %d/%d", correct, len(questions))
}

func parseLines(lines [][]string) []question {
	r := make([]question, len(lines))
	for i, line := range lines {
		r[i] = question{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return r
}

type question struct {
	q string
	a string
}
