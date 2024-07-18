package models

import (
	"encoding/json"
	"log"
	"regexp"
)

type Email struct {
	MessageID               string
	Date                    string
	From                    string
	To                      string
	Cc                      string
	Subject                 string
	MimeVersion             string
	ContentType             string
	ContentTransferEncoding string
	XFrom                   string
	XTo                     string
	XCc                     string
	XBcc                    string
	XFolder                 string
	XOrigin                 string
	XFileName               string
	Body                    string
}

func ParseEmail(emailText string) Email {
	var email Email

	// Define regex patterns for each field
	patterns := map[string]string{
		"MessageID":               `(?m)^Message-ID:\s*(.+)$`,
		"Date":                    `(?m)^Date:\s*(.+)$`,
		"From":                    `(?m)^From:\s*(.+)$`,
		"To":                      `(?m)^To:\s*(.+)$`,
		"Cc":                      `(?m)^Cc:\s*(.+)$`,
		"Subject":                 `(?m)^Subject:\s*(.+)$`,
		"MimeVersion":             `(?m)^Mime-Version:\s*(.+)$`,
		"ContentType":             `(?m)^Content-Type:\s*(.+)$`,
		"ContentTransferEncoding": `(?m)^Content-Transfer-Encoding:\s*(.+)$`,
		"XFrom":                   `(?m)^X-From:\s*(.+)$`,
		"XTo":                     `(?m)^X-To:\s*(.+)$`,
		"XCc":                     `(?m)^X-cc:\s*(.*)$`,
		"XBcc":                    `(?m)^X-bcc:\s*(.*)$`,
		"XFolder":                 `(?m)^X-Folder:\s*(.+)$`,
		"XOrigin":                 `(?m)^X-Origin:\s*(.+)$`,
		"XFileName":               `(?m)^X-FileName:\s*(.+)$`,
	}

	// Compile and execute regex patterns
	for field, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(emailText)
		if len(matches) > 1 {
			switch field {
			case "MessageID":
				email.MessageID = matches[1]
			case "Date":
				email.Date = matches[1]
			case "From":
				email.From = matches[1]
			case "To":
				email.To = matches[1]
			case "Cc":
				email.Cc = matches[1]
			case "Subject":
				email.Subject = matches[1]
			case "MimeVersion":
				email.MimeVersion = matches[1]
			case "ContentType":
				email.ContentType = matches[1]
			case "ContentTransferEncoding":
				email.ContentTransferEncoding = matches[1]
			case "XFrom":
				email.XFrom = matches[1]
			case "XTo":
				email.XTo = matches[1]
			case "XCc":
				email.XCc = matches[1]
			case "XBcc":
				email.XBcc = matches[1]
			case "XFolder":
				email.XFolder = matches[1]
			case "XOrigin":
				email.XOrigin = matches[1]
			case "XFileName":
				email.XFileName = matches[1]
			}
		}
	}

	// Extract the body of the email
	bodyPattern := regexp.MustCompile(`(?m)^X-FileName:.+\r?\n\r?\n([\s\S]+)`)
	bodyMatches := bodyPattern.FindStringSubmatch(emailText)
	if len(bodyMatches) > 1 {
		email.Body = bodyMatches[1]
	}

	return email
}

func ConvertEmailtoJSON(e Email) string {
	jsonEmail, err := json.Marshal(e)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	return string(jsonEmail)
}
