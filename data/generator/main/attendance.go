package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
)

type Attendee struct {
	ID         string
	FirstName  string
	LastName   string
	Email      string
	Company    string
	Country    string
	TicketType string
}

type BasicSession struct {
	SessionId string
	Title     string
	Track     string
	StartTime string
}

type AttendanceLog struct {
	AttendeeID string
	SessionID  string
	StartTime  string
}

func main() {
	// Read attendees
	attendees, err := readAttendees("../../../transform/data/raw_attendees.csv")
	if err != nil {
		log.Fatalf("Error reading attendees: %v", err)
	}

	// Read sessions
	sessions, err := readSessions("../../../transform/data/raw_sessions.csv")
	if err != nil {
		log.Fatalf("Error reading sessions: %v", err)
	}

	// Generate attendance logs
	logs := generateAttendanceLogs(attendees, sessions)

	// Write attendance logs
	err = writeAttendanceLogs(logs, "../../../transform/data/raw_attendance.csv")
	if err != nil {
		log.Fatalf("Error writing attendance logs: %v", err)
	}

	fmt.Printf("Successfully generated attendance logs for %d attendees across %d sessions\n", len(attendees), len(sessions))
}

func readAttendees(filename string) ([]Attendee, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Skip header
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	var attendees []Attendee
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		attendee := Attendee{
			ID:        record[0],
			FirstName: record[1],
			LastName:  record[2],
			Email:     record[3],
			Company:   record[4],
		}
		attendees = append(attendees, attendee)
	}

	return attendees, nil
}

func readSessions(filename string) ([]BasicSession, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Skip header
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	var sessions []BasicSession
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		session := BasicSession{
			SessionId: record[0],
			Title:     record[1],
			Track:     record[2],
			StartTime: record[3],
		}
		sessions = append(sessions, session)
	}

	return sessions, nil
}

func generateAttendanceLogs(attendees []Attendee, sessions []BasicSession) []AttendanceLog {
	var logs []AttendanceLog

	// Group sessions by start time
	sessionsByTime := make(map[string][]BasicSession)
	for _, session := range sessions {
		sessionsByTime[session.StartTime] = append(sessionsByTime[session.StartTime], session)
	}

	// For each time slot, distribute attendees evenly across sessions
	for startTime, concurrentSessions := range sessionsByTime {
		// Calculate how many attendees per session
		attendeesPerSession := len(attendees) / len(concurrentSessions)
		remainingAttendees := len(attendees) % len(concurrentSessions)

		// Create a copy of attendees slice to shuffle
		shuffledAttendees := make([]Attendee, len(attendees))
		copy(shuffledAttendees, attendees)
		// Shuffle attendees (you might want to implement a proper shuffle function)
		sort.Slice(shuffledAttendees, func(i, j int) bool {
			return shuffledAttendees[i].ID < shuffledAttendees[j].ID
		})

		attendeeIndex := 0
		// Distribute attendees across sessions
		for _, session := range concurrentSessions {
			// Calculate number of attendees for this session
			numAttendees := attendeesPerSession
			if remainingAttendees > 0 {
				numAttendees++
				remainingAttendees--
			}

			// Assign attendees to this session
			for i := 0; i < numAttendees && attendeeIndex < len(shuffledAttendees); i++ {
				logs = append(logs, AttendanceLog{
					AttendeeID: shuffledAttendees[attendeeIndex].ID,
					SessionID:  session.SessionId,
					StartTime:  startTime,
				})
				attendeeIndex++
			}
		}
	}

	return logs
}

func writeAttendanceLogs(logs []AttendanceLog, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	err = writer.Write([]string{"attendee_id", "session_id", "start_time"})
	if err != nil {
		return err
	}

	// Write logs
	for _, log := range logs {
		err := writer.Write([]string{
			log.AttendeeID,
			log.SessionID,
			log.StartTime,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
