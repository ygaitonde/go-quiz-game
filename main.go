package main
import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("csv", "problems.csv", "a csv file in the format: 'question,answer")
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
	fmt.Println(problems)
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
			a: line[1],
		}
	}
	return ret
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}