package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	firstNames  = []string{"James", "John", "Robert", "Michael", "William", "David", "Mary", "Patricia", "Jennifer", "Linda", "Elizabeth", "Barbara", "Emma", "Olivia", "Ava", "Sophia", "Alexander", "Liam", "Noah", "Oliver"}
	lastNames   = []string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez", "Anderson", "Taylor", "Thomas", "Moore", "Jackson", "Martin", "Lee", "Thompson", "White", "Harris"}
	companies   = []string{"Tech Corp", "Global Systems", "Innovation Labs", "Digital Solutions", "Software Inc", "Data Systems", "Cloud Tech", "Dev Enterprise", "Smart Solutions", "Future Tech"}
	countries   = []string{"USA", "Canada", "UK", "Germany", "France", "Australia", "Japan", "Singapore", "Netherlands", "Sweden"}
	ticketTypes = []string{"VIP", "Early Bird", "Regular", "Student", "Speaker"}
)

func main() {
	rand.Seed(time.Now().UnixNano())

	file, err := os.OpenFile("../../../transform/data/attendees.csv", os.O_RDWR, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"attendee_id", "first_name", "last_name", "email", "company", "country", "ticket_type"}
	err = writer.Write(header)
	if err != nil {
		panic(err)
	}

	// Generate 1000 attendees
	for i := 1; i <= 1000; i++ {
		firstName := firstNames[rand.Intn(len(firstNames))]
		lastName := lastNames[rand.Intn(len(lastNames))]
		email := fmt.Sprintf("%s.%s@%s.com",
			strings.ToLower(firstName),
			strings.ToLower(lastName),
			strings.ToLower(strings.ReplaceAll(companies[rand.Intn(len(companies))], " ", "")))

		record := []string{
			fmt.Sprintf("ATT%05d", i),
			firstName,
			lastName,
			email,
			companies[rand.Intn(len(companies))],
			countries[rand.Intn(len(countries))],
			ticketTypes[rand.Intn(len(ticketTypes))],
		}

		err := writer.Write(record)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
}
