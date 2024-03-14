package util

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func SendRequest(url string, body []byte, auth string, soapAction string) (int, []byte, error) {
	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return 0, nil, fmt.Errorf("error creating request: %s", err)
	}

	// Set the necessary request headers
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("Authorization", "Basic "+auth)
	req.Header.Add("SOAPAction", soapAction)

	// Initialize HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("error sending request: %s", err)
	}
	defer resp.Body.Close()

	// Read the response body
	buf := new(bytes.Buffer)
	io.Copy(buf, resp.Body)

	return resp.StatusCode, buf.Bytes(), nil
}
