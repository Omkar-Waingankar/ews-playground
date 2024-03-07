package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	ewsUrl := ""
	username := ""
	password := ""

	// Encode the credentials for basic authentication
	auth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))

	// SOAP request body
	soapBody := `<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"
   xmlns:t="http://schemas.microsoft.com/exchange/services/2006/types">
  <soap:Body>
    <GetFolder xmlns="http://schemas.microsoft.com/exchange/services/2006/messages"
               xmlns:t="http://schemas.microsoft.com/exchange/services/2006/types">
      <FolderShape>
        <t:BaseShape>Default</t:BaseShape>
      </FolderShape>
      <FolderIds>
        <t:DistinguishedFolderId Id="inbox"/>
      </FolderIds>
    </GetFolder>
  </soap:Body>
</soap:Envelope>`

	// Create a new HTTP request
	req, err := http.NewRequest("POST", ewsUrl, bytes.NewBuffer([]byte(soapBody)))
	if err != nil {
		log.Fatalf("Error creating request: %s", err)
	}

	// Set the necessary request headers
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("Authorization", "Basic "+auth)
	req.Header.Add("SOAPAction", "http://schemas.microsoft.com/exchange/services/2006/messages/GetFolder")

	// Initialize HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %s", err)
	}
	defer resp.Body.Close()

	// Read and print the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %s", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Println("Response Body:", string(body))
}
