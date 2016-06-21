package hubspot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Contact struct {
	APIKey     string     `json:"-"`
	Email      string     `json:"-"`
	Properties []Property `json:"properties"`
}

type ContactResp struct {
	Vid   int  `json:"vid"`
	IsNew bool `json:"isNew"`
}

func (h *Contact) Add(prop, value string) {
	h.Properties = append(h.Properties, Property{prop, value})
}

func NewContact(apiKey, email string) *Contact {
	return &Contact{
		APIKey: apiKey,
		Email:  email,
	}
}

// http://developers.hubspot.com/docs/methods/contacts/v2/get_contacts_properties
func (h *Contact) Publish() (cr *ContactResp) {
	const (
		hubspotUrl = "http://api.hubapi.com/contacts/v1/contact/createOrUpdate/email/%s/?hapikey=%s"
	)

	url := fmt.Sprintf(hubspotUrl, h.Email, h.APIKey)

	b, _ := json.Marshal(h)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	x, _ := ioutil.ReadAll(resp.Body)

	cr = &ContactResp{}
	err = json.Unmarshal(x, cr)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return
}
