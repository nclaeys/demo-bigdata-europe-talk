package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type Attendance struct {
	AttendeeID string
	SessionID  string
	StartTime  string
}

type Feedback struct {
	FeedbackID string
	AttendeeID string
	SessionID  string
	Rating     int
	Comments   string
}

// Predefined comment templates for different rating ranges
var commentTemplates = map[string][]string{
	"low": {
		"Content was too advanced for the target audience",
		"Needed more practical examples",
		"The pace was too fast to follow",
		"Could use more interactive elements",
		"Technical issues disrupted the flow",
		"Content felt disorganized",
		"More code examples would be helpful",
		"Slides were hard to read",
	},
	"medium": {
		"Good content but could use more real-world examples",
		"Interesting topic but needed more depth",
		"Decent presentation, could be more interactive",
		"Would benefit from more Q&A time",
		"Good overview but could use more detailed explanations",
		"Solid content but slides need improvement",
		"Nice presentation but too much theory",
		"Would be better with more hands-on exercises",
	},
	"high": {
		"Excellent content, maybe add more advanced topics",
		"Great presentation, could include more case studies",
		"Very informative, would love more deep dives",
		"Well structured, could add more interactive elements",
		"Fantastic session, consider adding workshop elements",
		"Really enjoyed it, would love more technical details",
		"Amazing content, perhaps add more time for questions",
		"Perfect pace, could include more industry examples",
	},
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Read attendance data
	attendances, err := readAttendance("../../../transform/data/raw_attendance.csv")
	if err != nil {
		log.Fatalf("Error reading attendance: %v", err)
	}

	// Generate feedback
	feedback := generateFeedback(attendances)

	// Write feedback to CSV
	err = writeFeedback(feedback, "../../../transform/data/raw_feedback.csv")
	if err != nil {
		log.Fatalf("Error writing feedback: %v", err)
	}

	fmt.Printf("Successfully generated feedback for %d sessions\n", len(feedback))
}

func readAttendance(filename string) ([]Attendance, error) {
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

	var attendances []Attendance
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		attendance := Attendance{
			AttendeeID: record[0],
			SessionID:  record[1],
			StartTime:  record[2],
		}
		attendances = append(attendances, attendance)
	}

	return attendances, nil
}

func generateFeedback(attendances []Attendance) []Feedback {
	var feedback []Feedback

	// Assume 80% of attendees provide feedback
	for i, attendance := range attendances {
		if rand.Float32() < 0.8 { // 80% chance of providing feedback
			rating := generateRating()
			feedback = append(feedback, Feedback{
				FeedbackID: fmt.Sprintf("FB%05d", i+1),
				AttendeeID: attendance.AttendeeID,
				SessionID:  attendance.SessionID,
				Rating:     rating,
				Comments:   generateComment(rating),
			})
		}
	}

	return feedback
}

func generateRating() int {
	// Weighted random distribution favoring middle-high ratings
	weights := []int{1, 2, 10, 25, 35, 27} // Weights for ratings 0-5
	totalWeight := 0
	for _, w := range weights {
		totalWeight += w
	}

	r := rand.Intn(totalWeight)
	for rating, weight := range weights {
		r -= weight
		if r < 0 {
			return rating
		}
	}
	return 5
}

func generateComment(rating int) string {
	var commentType string
	switch {
	case rating <= 2:
		commentType = "low"
	case rating <= 3:
		commentType = "medium"
	default:
		commentType = "high"
	}

	comments := commentTemplates[commentType]
	return comments[rand.Intn(len(comments))]
}

func writeFeedback(feedback []Feedback, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	err = writer.Write([]string{"feedback_id", "attendee_id", "session_id", "rating", "comments"})
	if err != nil {
		return err
	}

	// Write feedback entries
	for _, f := range feedback {
		err := writer.Write([]string{
			f.FeedbackID,
			f.AttendeeID,
			f.SessionID,
			fmt.Sprintf("%d", f.Rating),
			f.Comments,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
