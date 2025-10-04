package vintsnipe

import (
	"net/http"
	"fmt"
	"vintsnipe/client"
)

type Pol struct {
	Request      *http.Request
	Client       *http.Client
	ChannelID    string
	LatestItemID int
}

var PolSlice = []*Pol{}

func (*Pol) Init() *Pol {
	return &Pol{}
}

func Create(channelID string, search string) (*Pol, error) {

	iClient := &http.Client{}
	url := fmt.Sprintf("https://www.vinted.co.uk/api/v2/catalog/items?page=1&per_page=96&time=1758555430&global_search_session_id=8d83c2b9-8740-424c-a6c1-943c66c6f1f4&search_text=%s&currency=GBP&order=newest_first", search)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//fmt.Errorf("failed to create a new request: ", err)
		return nil, err
	}
	req.Header = client.Headers
	req.Header.Set("Cookie", client.GetCookiesString())

	var Pol = &Pol{
		Request:      req,
		Client:       iClient,
		ChannelID:    channelID,
		LatestItemID: 0,
	}

	return Pol, err

}

