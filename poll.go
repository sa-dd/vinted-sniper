package vintsnipe

import "net/http"

type pol struct {
	Request      *http.Request
	Client       *http.Client
	ChannelID    string
	LatestItemID int
}

func (*pol) Init() *pol {
	return &pol{}
}

func (*pol) Create(channelID string, search string) (*pol, error) {

	client := &http.Client{}
	url := fmt.Sprintf("https://www.vinted.co.uk/api/v2/catalog/items?page=1&per_page=96&time=1758555430&global_search_session_id=8d83c2b9-8740-424c-a6c1-943c66c6f1f4&search_text=%s&currency=GBP&order=newest_first", search)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//fmt.Errorf("failed to create a new request: ", err)
		return nil, err
	}
	req.Header = Headers
	req.Header.Set("Cookie", GetCookiesString())

	var pol = &vintsnipe.pol{
		Request:      req,
		Client:       client,
		ChannelID:    channelID,
		LatestItemID: 0,
	}

	return pol, err

}

