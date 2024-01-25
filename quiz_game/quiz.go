package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	filename := flag.String("f", "problems.csv", "Specify filename of csv files with problems")
	waitTime := flag.Int("t", 5, "Time in seconds to solve each problem")
	flag.Parse()
	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	fmt.Println("Press return to start the quiz")
	os.Stdin.Read(make([]byte, 1))
	correct, total := 0, 0
	input := make(chan string)
	go parseInput(input)
quiz:
	for record, err := r.Read(); record != nil; record, err = r.Read() {
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("Problem %s, you have %d seconds \n", record[0], *waitTime)
		wait := time.NewTimer(time.Duration(*waitTime) * time.Second)
		select {
		case answer := <-input:
			if answer == strings.TrimSpace(record[1]) {
				correct++
			}
		case <-wait.C:
			wait.Stop()
			fmt.Println("Time elapsed")
			break quiz
		}
		total++
	}
	fmt.Printf("Overall result is %d/%d \n", correct, total)
}

func parseInput(input chan<- string) {
	for {
		var answer string
		_, err := fmt.Scanf("%s\n", &answer)
		if err != nil {
			log.Println(err)
		}
		input <- answer
	}
}
