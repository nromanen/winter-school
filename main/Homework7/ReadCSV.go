package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Open the CSV file
	file, err := os.Open("ds1.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	// Print each record to the console
	for _, record := range records {
		fmt.Printf("Employee ID: %s\n", record[0])
		fmt.Printf("Name: %s %s %s %s\n", record[1], record[2], record[3], record[4])
		fmt.Printf("Gender: %s\n", record[5])
		fmt.Printf("Email: %s\n", record[6])
		fmt.Printf("Father's Name: %s\n", record[7])
		fmt.Printf("Mother's Name: %s\n", record[8])
		fmt.Printf("Mother's Maiden Name: %s\n", record[9])
		fmt.Printf("Salary: %s\n", record[10])
		fmt.Println("--------------------")
	}
}
