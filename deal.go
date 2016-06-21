package hubspot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Deal struct {
	APIKey string `json:"-"`

	Associations struct {
		AssociatedCompanyIds []int `json:"associatedCompanyIds,omitempty"`
		AssociatedVids       []int `json:"associatedVids,omitempty"`
	} `json:"associations,omitempty"`

	PortalID int `json:"portalId,omitempty"`

	Properties []PropertyDeal `json:"properties,omitempty"`
}

func NewDeal(apiKey string) *Deal {
	return &Deal{
		APIKey: apiKey,
	}
}

func (h *Deal) Add(prop string, value interface{}) {
	h.Properties = append(h.Properties, PropertyDeal{prop, value})
}

func (h *Deal) Publish() {
	const (
		hubspotUrl = "http://api.hubapi.com/deals/v1/deal/?hapikey=%s"
	)

	url := fmt.Sprintf(hubspotUrl, h.APIKey)

	b, _ := json.Marshal(h)

	fmt.Println(string(b))

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
	fmt.Println("Hubspot body", string(x), resp)
}
