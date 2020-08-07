package main
import "flag"

func main() {
	filename := flag.String("csv", "problems.csv", "a csv file in the format: 'question,answer")
	flag.Parse()
	_ = filename
}