package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Omkar-Waingankar/ews-playground/ews"
)

const ewsUrl = "https://west.EXCH092.serverdata.net/EWS/Exchange.asmx"
const username = "test@nylas.info"
const password = ""

func main() {
	ewsClient := ews.NewEWSClient(ewsUrl, username, password)

	// statusCode, body, err := ewsClient.ListMessages()
	statusCode, body, err := ewsClient.GetMessage(
		"AQAPAHRlc3RAbnlsYXMuaW5mbwBGAAAE/WS9WdWISpVv9iZzojjeBwColNEj/TvPTqUR8aHNdtOrAAACAQwAAAColNEj/TvPTqUR8aHNdtOrAAL4XvdiAAAA",
		"CQAAABYAAAColNEj/TvPTqUR8aHNdtOrAAL4iMT5",
	)
	if err != nil {
		log.Fatalf("Error getting folder: %s", err)
	}

	fmt.Printf("Status Code: %d\n", statusCode)

	// Create a new file to write the response body
	file, err := os.Create("response.xml")
	if err != nil {
		log.Fatalf("Error creating file: %s", err)
	}
	defer file.Close()

	// Write the response body to the file
	_, err = file.Write(body)
	if err != nil {
		log.Fatalf("Error writing to file: %s", err)
	}

	fmt.Println("Response body written to response.xml")
}
