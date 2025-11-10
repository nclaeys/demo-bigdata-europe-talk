package main

import (
	"encoding/csv"
	"fmt"
	"github.com/playwright-community/playwright-go"
	"log"
	"math/rand"
	"os"
	"strings"
)

type Session struct {
	SessionId   string
	Title       string
	Track       string
	StartTime   string
	Description string
	Speakers    []string
}

func main() {
	// Install playwright dependencies
	err := playwright.Install()
	if err != nil {
		log.Fatalf("Error installing playwright: %v", err)
	}

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("Error starting playwright: %v", err)
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	if err != nil {
		log.Fatalf("Error launching browser: %v", err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("Error creating new page: %v", err)
	}

	// Navigate to the conference page
	if _, err = page.Goto("https://events.pinetool.ai/3499/#sessions", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	}); err != nil {
		log.Fatalf("Error navigating to page: %v", err)
	}

	// Wait for the sessions container to be visible
	// You might need to adjust this selector based on the actual page structure
	page.WaitForSelector("div.sessions-container", playwright.PageWaitForSelectorOptions{
		State: playwright.WaitForSelectorStateVisible,
	})

	// Extract sessions data
	sessions, err := extractSessions(page)
	if err != nil {
		log.Fatalf("Error extracting sessions: %v", err)
	}

	// Create CSV file
	file, err := os.OpenFile("../../sessions.csv", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	header := []string{"Id", "Title", "Track", "Start Time", "Description", "Speakers"}
	if err := writer.Write(header); err != nil {
		log.Fatal(err)
	}

	// Write sessions to CSV
	for _, session := range sessions {
		record := []string{
			RandomString(8),
			session.Title,
			session.Track,
			session.StartTime,
			session.Description,
			strings.Join(session.Speakers, "; "),
		}
		if err := writer.Write(record); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Successfully extracted %d sessions to conference_sessions.csv\n", len(sessions))
}

func RandomString(n int) string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func extractSessions(page playwright.Page) ([]Session, error) {
	var sessions []Session

	// Get all session elements
	// Note: You'll need to adjust these selectors based on the actual HTML structure
	sessionElements, err := page.QuerySelectorAll("div.session")
	if err != nil {
		return nil, fmt.Errorf("error finding session elements: %v", err)
	}

	for _, element := range sessionElements {
		session := Session{}

		// Extract title
		if mainContent, err := element.QuerySelector(".main"); err == nil && mainContent != nil {
			if titleContent, err := element.QuerySelector(".head"); err == nil && titleContent != nil {
				content, err := titleContent.TextContent()
				if err == nil {
					content = strings.ReplaceAll(content, "nbsp", "")
					if strings.Contains(content, "|") {
						titleParts := strings.SplitN(content, "|", 2)
						session.Title = strings.TrimSpace(titleParts[1])
					} else {
						session.Title = strings.TrimSpace(content)
					}
				}
			}
			if descriptionContent, err := element.QuerySelector(".description"); err == nil && descriptionContent != nil {
				descriptionTxt, err := descriptionContent.TextContent()
				if err == nil {
					descriptionTxt = strings.ReplaceAll(descriptionTxt, "&nbsp;", "")
					session.Description = strings.TrimSpace(descriptionTxt)
				}
			}
		}

		// Extract date
		if dateElement, err := element.QuerySelector(".datetime"); err == nil && dateElement != nil {
			attribute, err := dateElement.GetAttribute("datetime")
			if err == nil {
				session.StartTime = attribute
			}
		}

		// Extract track
		if trackElement, err := element.QuerySelector(".theme"); err == nil && trackElement != nil {
			trackText, _ := trackElement.TextContent()
			session.Track = trackText
		}

		// Extract speakers
		speakerElements, err := element.QuerySelectorAll("nav.speakers")
		if err == nil && len(speakerElements) > 0 {
			allSpeakers, err := speakerElements[0].QuerySelector(".speakers-short")
			if err == nil && allSpeakers != nil {
				linkToSpeakers, err := allSpeakers.QuerySelectorAll("a")
				if err == nil && linkToSpeakers != nil && len(linkToSpeakers) > 0 {
					for _, speakerElement := range linkToSpeakers {
						speakerName, err := speakerElement.GetProperty("title")
						if err != nil {
							fmt.Printf("Error getting speaker name: %v", err)
						} else {
							if speakerName.String() != "" {
								name := strings.Split(speakerName.String(), "|")
								if len(name) > 0 {
									session.Speakers = append(session.Speakers, strings.TrimSpace(name[0]))
								}
							}
						}
					}
				}
			} else {
				fullSpeaker, err := speakerElements[0].QuerySelector(".speakers-full")
				if err == nil && fullSpeaker != nil {
					nameSelector, err := fullSpeaker.QuerySelector(".name")
					if err == nil && nameSelector != nil {
						speakerName, _ := nameSelector.TextContent()
						session.Speakers = append(session.Speakers, strings.TrimSpace(speakerName))
					} else {
						speakerName, _ := fullSpeaker.TextContent()
						speakerName = strings.TrimPrefix(speakerName, "Speakers")
						session.Speakers = append(session.Speakers, strings.TrimSpace(speakerName))
					}
				}
			}

		} else {
			fmt.Printf("Error finding speaker elements: %v", err)
		}

		sessions = append(sessions, session)
	}

	return sessions, nil
}
