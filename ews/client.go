package ews

import (
	"encoding/base64"
	"fmt"

	"github.com/Omkar-Waingankar/ews-playground/util"
)

type EWSClient struct {
	Url      string
	Username string
	Password string
}

func NewEWSClient(url, username, password string) *EWSClient {
	return &EWSClient{
		Url:      url,
		Username: username,
		Password: password,
	}
}

func (c *EWSClient) GetFolder(id string) (int, []byte, error) {
	// Encode the credentials for basic authentication
	auth := base64.StdEncoding.EncodeToString([]byte(c.Username + ":" + c.Password))

	// SOAP request body
	soapBody := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"
   xmlns:t="http://schemas.microsoft.com/exchange/services/2006/types">
  <soap:Body>
    <GetFolder xmlns="http://schemas.microsoft.com/exchange/services/2006/messages"
               xmlns:t="http://schemas.microsoft.com/exchange/services/2006/types">
      <FolderShape>
        <t:BaseShape>Default</t:BaseShape>
      </FolderShape>
      <FolderIds>
        <t:DistinguishedFolderId Id="%s"/>
      </FolderIds>
    </GetFolder>
  </soap:Body>
</soap:Envelope>`, id)

	return util.SendRequest(c.Url, []byte(soapBody), auth, "http://schemas.microsoft.com/exchange/services/2006/messages/GetFolder")
}

func (c *EWSClient) ListMessages() (int, []byte, error) {
	// Encode the credentials for basic authentication
	auth := base64.StdEncoding.EncodeToString([]byte(c.Username + ":" + c.Password))

	// SOAP request body
	soapBody := `<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"
	xmlns:t="http://schemas.microsoft.com/exchange/services/2006/types">
	<soap:Body>
	<FindItem xmlns="http://schemas.microsoft.com/exchange/services/2006/messages"
				xmlns:t="http://schemas.microsoft.com/exchange/services/2006/types"
				Traversal="Shallow">
		<ItemShape>
		<t:BaseShape>IdOnly</t:BaseShape>
		</ItemShape>
		<ParentFolderIds>
		<t:DistinguishedFolderId Id="inbox"/>
		</ParentFolderIds>
	</FindItem>
	</soap:Body>
</soap:Envelope>`

	return util.SendRequest(c.Url, []byte(soapBody), auth, "https://schemas.microsoft.com/exchange/services/2006/messages/FindItem")
}

func (c *EWSClient) GetMessage(id, changeKey string) (int, []byte, error) {
	// Encode the credentials for basic authentication
	auth := base64.StdEncoding.EncodeToString([]byte(c.Username + ":" + c.Password))

	// SOAP request body
	soapBody := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
	<soap:Envelope
		xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"
		xmlns:t="http://schemas.microsoft.com/exchange/services/2006/types">
		<soap:Body>
		<GetItem
			xmlns="http://schemas.microsoft.com/exchange/services/2006/messages"
			xmlns:t="http://schemas.microsoft.com/exchange/services/2006/types">
			<ItemShape>
				<t:BaseShape>AllProperties</t:BaseShape>
				<t:IncludeMimeContent>false</t:IncludeMimeContent>
			</ItemShape>
			<ItemIds>
				<t:ItemId Id="%s" ChangeKey="%s" />
			</ItemIds>
		</GetItem>
	</soap:Body>
</soap:Envelope>`, id, changeKey)

	return util.SendRequest(c.Url, []byte(soapBody), auth, "https://schemas.microsoft.com/exchange/services/2006/messages/GetItem")
}
