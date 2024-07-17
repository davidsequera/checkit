package main

import (
	"fmt"
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

func parseEmail(emailText string) Email {
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

func main() {
	emailText := `Message-ID: <19885966.1075841531676.JavaMail.evans@thyme>
Date: Sun, 27 May 2001 08:48:03 -0700 (PDT)
From: s..white@enron.com
To: asem.atta@enron.com
Subject: RE: Status?
Cc: diana.scholtes@enron.com, holden.salisbury@enron.com, chip.cox@enron.com, 
	david.steiner@enron.com, nikolay.kraltchev@enron.com, 
	william.crooks@enron.com
Mime-Version: 1.0
Content-Type: text/plain; charset=us-ascii
Content-Transfer-Encoding: 7bit
Bcc: diana.scholtes@enron.com, holden.salisbury@enron.com, chip.cox@enron.com, 
	david.steiner@enron.com, nikolay.kraltchev@enron.com, 
	william.crooks@enron.com
X-From: White, Michael S. </O=ENRON/OU=NA/CN=RECIPIENTS/CN=MWHITE3>
X-To: Atta, Asem </O=ENRON/OU=NA/CN=RECIPIENTS/CN=Aatta>
X-cc: Scholtes, Diana </O=ENRON/OU=NA/CN=RECIPIENTS/CN=Dscholt>, Salisbury, Holden </O=ENRON/OU=NA/CN=RECIPIENTS/CN=Hsalisbu>, Cox, Chip </O=ENRON/OU=NA/CN=RECIPIENTS/CN=Ccox3>, Steiner, David </O=ENRON/OU=NA/CN=RECIPIENTS/CN=Dsteine2>, Kraltchev, Nikolay </O=ENRON/OU=NA/CN=RECIPIENTS/CN=Nkraltc>, Crooks, William </O=ENRON/OU=NA/CN=RECIPIENTS/CN=Wcrooks>
X-bcc: 
X-Folder: \ExMerge - Salisbury, Holden\Read
X-Origin: SALISBURY-H
X-FileName: holden salisbury 6-26-02.PST

I found that the said icons where never moved out to Portland.  I have moved these and added the foster security group to them.  All they should have to do is refresh there start menu..  I checked the scripts and the installation packages and they are correct so everything should work properly.

If you have any problems feel free contact me about it.

Michael S White
713-853-9001

 -----Original Message-----
From: 	Atta, Asem  
Sent:	Friday, May 25, 2001 9:12 AM
To:	White, Michael S.
Cc:	Scholtes, Diana; Salisbury, Holden; Cox, Chip; Steiner, David; Kraltchev, Nikolay; Crooks, William
Subject:	FW: Status?

Michael,

	Holden Salisbury, trader in our PDX office, seems to be having problems [emails included below] with getting full access to the Foster project's applications. Seems that when he browses through his program menu's on his local machine he sees only Oasis Net Exports and Transmission Ticker and when he goes through TerminalServer he's got access to all of them, he needs to be able to access all apps [ Transmission Ticker, Net Exports. Transmission Offerings, Reservation Reports, Batch Scheduler, ATC reports] from his local machine.  IT IRM  contacts up on the Portland side are Chip Cox [503-703-0536] and David Steiner [503-703-1553]. Please do let whomever know what needs to be done in order to get this resolved.

thanks in advance
Asem
x31700
	

 -----Original Message-----
From: 	Salisbury, Holden  
Sent:	Thursday, May 24, 2001 02:47 PM
To:	Atta, Asem; Cox, Chip
Cc:	Steiner, David; Kane, Paul
Subject:	RE: Status?

I need all of the Foster applications on my local profile.  I have some of them on local and some of them on my terminal server profile.  I need the same Foster profile that Diana Scholtes has on her local machine.  As of this morning I still do not have it.  Please let me know when my profile gets updated.

Thank You,
Holden

 -----Original Message-----
From: 	Atta, Asem  
Sent:	Wednesday, May 23, 2001 7:18 AM
To:	Cox, Chip
Cc:	Steiner, David; Kane, Paul; Salisbury, Holden
Subject:	RE: Status?

Chip,

	The issue seems to be that of an Information Risk Management problem where they haven't pushed out all the application access to Holden's profile.. I had put in a ticket with (31411) under your name for this matter on Monday. I will call them and ask them to escalate this matter.

regards
Asem
x31700

 -----Original Message-----
From: 	Cox, Chip  
Sent:	Tuesday, May 22, 2001 05:33 PM
To:	Atta, Asem
Cc:	Steiner, David; Kane, Paul; Salisbury, Holden
Subject:	Status?

Hi Asem, just wondering if you've gotten any updates on the Foster Application issue for Holden Salisbury. 

Thanks,  Chip`

	email := parseEmail(emailText)
	fmt.Printf("%+v\n", email)
}
