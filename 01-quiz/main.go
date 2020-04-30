package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

const problemsFileName = "problems.csv"

func main() {
	// read from csv file via CSV PACKAGE
	f, err := os.Open(problemsFileName)
	if err != nil {
		fmt.Printf("Failed to open the file :%v\n", err)
		return
	}
	defer f.Close()

	// read csv file via csv package
	r := csv.NewReader(f)

	questions, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Failed to read the file %+v\n", err)
	}
    var correctAnswers int
	for i, record := range questions {
		question, correctAns := record[0], record[1]
		fmt.Printf("%d. %s ?\n", i+1,question)
		var answer string
		
		if _,err := fmt.Scan(&answer); err != nil {
			fmt.Printf("Filed to scan %v\n",err)
			return
		}
		
		
		if answer == correctAns {
			correctAnswers++
		}


		fmt.Printf("Result :  Correct answers : %d Total : %d\n",correctAnswers,len(questions))
	}
}
