package main
import (
	"encoding/csv"
	"flag"
	"fmt"
	"strings"
	"os"
	"time"
)

func main() {
	filename := flag.String("csv", "problems.csv", "a csv file in the format: 'question,answer")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	
	file, err := os.Open(*filename)
	if err!=nil {
		exit(fmt.Sprintf("Failed to open %s", *filename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	problems := parseLines(lines)
	
	timer:= time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0

  problemLoop:
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case answer := <- answerCh: 
			if answer == problem.a {
				correct++
			}
		}

	}
	fmt.Printf("You scored %d out of %d", correct, len(problems))
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}