package plaid

import (
	"bytes"
	"encoding/json"
)

// Partial implementation of Plaid's info product
// Currently only supports POST /info/get

// InfoGet (POST /info/get) retrieves account holder information(names, emails, phone numbers and addresses) after the user's account has been linked
// See https://plaid.com/docs/legacy/api/#get-info-data
func (c *Client) InfoGet(accessToken string) (postRes *PostResponse, mfaRes *MFAResponse, err error) {

	jsonText, err := json.Marshal(infoGetJson{
		ClientID:    c.clientID,
		Secret:      c.secret,
		AccessToken: accessToken,
	})
	if err != nil {
		return nil, nil, err
	}
	return c.postAndUnmarshal("/info/get", bytes.NewReader(jsonText))
}

type infoGetJson struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}
