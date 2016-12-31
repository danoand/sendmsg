package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	twilioAccountSID = "AC1661d55ca882016d8e4c038c47acc8bc"
	twilioAuthToken  = "62d5f0c0455961f085641e58963fc1d0"
	twilioNumber     = "+16302257108"
)

var (
	args []string

	// Set up the elements of the REST API URL
	twilioURLElement1         = "https://"
	twilioURLElement2         = ":"
	twilioURLElement3         = "@api.twilio.com/2010-04-01/Accounts/"
	twilioURLElement4         = "/Messages.json"
	twilioConstructMessageURL = twilioURLElement1 + twilioAccountSID + twilioURLElement2 + twilioAuthToken +
		twilioURLElement3 + twilioAccountSID + twilioURLElement4
)

// sendTwilioText triggers a Text message via Twilio
func sendTwilioText(phn string, msg string) bool {
	var (
		rErr   error
		retVal = true
		max    int
		twResp *http.Response
	)

	// Set the maxlength of the body to be sent
	if max = len(msg); max > 1600 {
		max = 1600
	}

	// Truncate the message text to a max length (if needed)
	msg = msg[0:max]

	// Construct the parameters to be sent via the HTTP POST
	urlVals := url.Values{}
	urlVals.Set("From", twilioNumber)
	urlVals.Set("To", phn)
	urlVals.Set("Body", msg)

	// Configure an HTTP POST
	if twResp, rErr = http.PostForm(twilioConstructMessageURL, urlVals); rErr != nil {
		log.Printf("ERROR: error posting text message to Twilio. See: %v\n", rErr)
		retVal = false
	}

	// Log the Twilo response
	log.Printf("INFO: Status of Twilio response: %v\n", twResp.Status)

	// Return to caller
	return retVal
}

// prtHelp prints a short help blurb to standard out
func prtHelp() {
	fmt.Printf("\n----- 'sendmsg' Help -----\n")
	fmt.Printf("01 Use 'sendmsg' to send a text message to a 10 digit phone number\n")
	fmt.Printf("02 Command format: ' sendmsg <phone number> <text message body> '\n")
	fmt.Printf("03 Send a text to a phone number: ' sendmsg \"4158675309\" \"Meet you at 10\" '\n")
	fmt.Printf("04 Send a text message to your phone: ' sendmsg \"Meet you at 10\" '\n")
	fmt.Printf("05 NOTE: If your parameters include embedded spaces, remember to enclose those parameters in quotes\n")
	fmt.Printf("----- 'sendmsg' Help -----\n\n")
}

func main() {
	// Grab the command line arguments
	args = os.Args

	// Missing command arguments? Print help.
	if len(args) == 1 {
		prtHelp()
		goto WrapUp
	}

	// User requesting 'help'
	if len(args) == 1 || args[1] == "help" || args[1] == "h" {
		prtHelp()
		goto WrapUp
	}

WrapUp:
}
