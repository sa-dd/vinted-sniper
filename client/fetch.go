package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type VintedResponse struct {
	Items []VintedItem `json:"items"`
}

type VintedItem struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	Status    string `json:"status"`
	SizeTitle string `json:"size_title"`
	User      struct {
		Login string `json:"login"`
	} `json:"user"`
	Photo struct {
		URL string `json:"url"`
	} `json:"photo"`
	Price struct {
		Amount       string `json:"amount"`
		CurrencyCode string `json:"currency_code"`
	} `json:"price"`
}

func FetchVintedItems(client *http.Client) ([]VintedItem, error) {
	const vintedURL = "https://www.vinted.co.uk/api/v2/catalog/items?page=1&per_page=96&time=1758397863&global_search_session_id=d9bc1e4f-6b8d-45a1-90d8-fb8a041ef637&search_text=uggs+slippers&catalog_ids=&order=newest_first&size_ids=&brand_ids=&status_ids=&color_ids=&material_ids="

	req, err := http.NewRequest("GET", vintedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	// Set required headers
	req.Header = http.Header{
		"authority":          {"www.vinted.co.uk"},
		"accept":             {"application/json, text/plain, */*"},
		"accept-language":    {"en-uk-fr"},
		"cache-control":      {"max-age=0"},
		"cookie":             {"v_udt=L2JGTWk0RHJmVWN2QUYzVDZRaXFrVkFGTFBXci0tejE0bzlkb09USjVucHNEZi0tSk40aXhySHV6U0pBMEpQdXBvNERTUT09; anon_id=2c165bcb-8951-4dfc-8c70-7fab5d7847f7; anonymous-locale=en-uk-fr; access_token_web=eyJraWQiOiJFNTdZZHJ1SHBsQWp1MmNObzFEb3JIM2oyN0J1NS1zX09QNVB3UGlobjVNIiwiYWxnIjoiUFMyNTYifQ.eyJhcHBfaWQiOjQsImF1ZCI6ImZyLmNvcmUuYXBpIiwiY2xpZW50X2lkIjoid2ViIiwiZXhwIjoxNzU4NDA3NjgxLCJpYXQiOjE3NTg0MDA0ODEsImlzcyI6InZpbnRlZC1pYW0tc2VydmljZSIsInB1cnBvc2UiOiJhY2Nlc3MiLCJzY29wZSI6InB1YmxpYyIsInNpZCI6ImFmYWM0MGM3LTE3NTg0MDA0ODEifQ.YH427er3ce_qRR1WLDzF2_-TRD7IkuIbSvKlaZi1Qc5SNEEb0vYRdA40z7kby22l7tVhYa5bPlxSlLti6jGDjV2ladxVe0lwxZZ8ZAl77dDIZZSpMJlmT7737fHE-Ux41gJoOACDzI_XD-Zw-5-YF7eQJEDwKWEKf6v0BtgWpb16UOW-CBRsDaA5z_huuKujvovjCGKAOLB0S2LhbMXggCag0bvll6D2ADdaUoj94jrTfDLnMCtKunzffmv53fE2RHxAilTdWzXYnjliX3RbSrTqG6sV4JK5x4qj6hY18IzkKlvSjTHuRdoJF4W6LruQN7CJ4cbEHr-qgOBaanuA9A; refresh_token_web=eyJraWQiOiJFNTdZZHJ1SHBsQWp1MmNObzFEb3JIM2oyN0J1NS1zX09QNVB3UGlobjVNIiwiYWxnIjoiUFMyNTYifQ.eyJhcHBfaWQiOjQsImF1ZCI6ImZyLmNvcmUuYXBpIiwiY2xpZW50X2lkIjoid2ViIiwiZXhwIjoxNzU5MDA1MjgxLCJpYXQiOjE3NTg0MDA0ODEsImlzcyI6InZpbnRlZC1pYW0tc2VydmljZSIsInB1cnBvc2UiOiJyZWZyZXNoIiwic2NvcGUiOiJwdWJsaWMiLCJzaWQiOiJhZmFjNDBjNy0xNzU4NDAwNDgxIn0.hTZysPc80sDo2mhOl58oLub8Et1zR_0w3g_g-CmV98MUQnOuBu1rMob4yZyE_cCDMqb0h7cCL0AV08rFuU78bekCBkqSlQ2e9sI2xlJbWQstCUA64WsFFU1w1kxOa4AHpdkggrQ6us2ZuexpxwkxGuXOLiwOic1RjlSGJs9s2M3swtsOu-l3QeBOs-JWbtzNfyUAIxsR6puyQ0mfxx3Zk7TxuHv0w8De7XhbEj3COpVVzGd34szBVZ_aTQwZPLLUZ9_FSFSHZzSmha8fR_aTB-8Lbw7iOXEHaARIW4mQyOJGvbbN47UDvl-X_QzIYCDB4ncDTQXr7lGG5y4cBiDQJQ; anon_id=60cbd6a4-9b5c-4d35-81bc-957fc2373e88; __cf_bm=3uVc.Hq8hPwwaMYxE87lsdOK.qVJLsiDODLNdp31fi8-1758400483-1.0.1.1-1FkC5Aw1MrjPyZhelO0vlujHaiHm66QGPPg_gpEN6W2OEjOezTMQ9jyqu7dvm8Y8TmuGenBaW1SZRBo_Zew0x9ItIcaDf81PJw1bZKw5U_RLs..pXGSISvJ8hD0i4opb; cf_clearance=SwVbLAWzqE03JbqDkHXOnlctuc2_ppOrtsbnNK_.PXA-1758400484-1.2.1.1-h0cQZcrrJQ1fWpKfRYPl26dP.Vh2LKYz.pnm2kT8ucI1B5YueLXp6i49RFHjXPenba7isVPWhUZa_OUAOtQd_.ZMr.pbiQoPrA3y44D70GubQlFi1Z9Kp.vNgh1o7p5pTqZml3DEQLtpi9ZC6apTkscvLCjSsbnZ1cqJkVyiIS1FQh9F0I_G0IcGh5EL5.s_l3Vb7Q7aXMS9nk5vCvbHD16gvPpE15RpyA3Yb7l8Gok; viewport_size=677; v_sid=e1dad7acf68ce1604789335976e4ce9e; OptanonAlertBoxClosed=2025-09-20T20:34:47.490Z; eupubconsent-v2=CQYCsRgQYCsRgAcABBENB9FgAAAAAEPgAAwIAAAWZABMNCogjLIgBCJQMAIEACgrCACgQBAAAkDRAQAmDApyBgAusJkAIAUAAwQAgABBgACAAASABCIAKACAQAAQCBQABgAQBAQAMDAAGACxEAgABAdAxTAggECwASMyqDTAlAASCAlsqEEgCBBXCEIs8AggREwUAAAIABQEAADwWAhJICViQQBcQTQAAEAAAUQIECKRswBBQGaLQXgyfRkaYBg-YJklMgyAJgjIyTYhN-Ew8UhRAAAA.YAAACHwAAAAA; OTAdditionalConsentString=1~; OptanonConsent=isGpcEnabled=0&datestamp=Sat+Sep+20+2025+21%3A35%3A14+GMT%2B0100+(British+Summer+Time)&version=202508.2.0&browserGpcFlag=0&isIABGlobal=false&consentId=2c165bcb-8951-4dfc-8c70-7fab5d7847f7&identifierType=Cookie+Unique+Id&isAnonUser=1&hosts=&interactionCount=1&landingPath=NotLandingPage&groups=C0001%3A1%2CC0002%3A0%2CC0003%3A0%2CC0004%3A0%2CC0005%3A0%2CV2STACK42%3A0%2CC0015%3A0%2CC0035%3A0%2CC0038%3A0&genVendors=V2%3A0%2CV1%3A0%2C&intType=2&geolocation=GB%3BENG&AwaitingReconsent=false; banners_ui_state=SUCCESS; datadome=nXJD0p9zPeWgCIWVEQDnSOntFurT2frb6YRE51DydF6P7R0kbKfYAtpmMx~Wfar8zhXLBfv1zC2FOd8n_LR9nYNyIuFVbXksFqD2XOFxOy3Yd5EqMPTXcjTC5d~hkPcQ; _vinted_fr_session=VFc3cWliZ0ZoWGdObkdwSit1dzB6MEZ3dDQvZmFxK2gzU3E5bHdJY2U3UEJjN3hsdDlBbmlERnkvS2JCU1plbVZVOUtaWlM2cy8wS2liWGtTTGEyazRBU1lVV2Y2b1FvU3BrT2wwZmorRmlZZk9uTmxBUlh6RWpnZ2pLQU9uMGdnUW5lV0x4WHdZdWtPanFtQVVpMnRZZGtEYkQ0MldtMEt0ejhNaUo4OGVaZWVJSGw0RVVlODhIY1J5YldHK20vaWZRY1p2WlRGWVBnNzNLK2Q4eWEweVN5c2U3THJ2VXdTVFMyenFCVis5TXVqS0Y5WEhCZFNwL1hwQmRhWlNpaS0telh2RXIvajR0VjZRVG0rOHk1S1Q0QT09--e0e081c7d4d66dc655160ea2dab44526af0315d7"}, // Get cookie from env
		"priority":           {"u=1, i"},
		"referer":            {"https://www.vinted.co.uk/catalog"},
		"sec-ch-ua":          {"\"Chromium\";v=\"136\", \"Microsoft Edge\";v=\"136\", \"Not.A/Brand\";v=\"99\""},
		"sec-ch-ua-mobile":   {"?0"},
		"sec-ch-ua-platform": {"\"Windows\""},
		"sec-fetch-dest":     {"empty"},
		"sec-fetch-mode":     {"cors"},
		"sec-fetch-site":     {"same-origin"},
		"user-agent":         {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36 Edg/136.0.0.0"},
		"x-anon-id":          {"2c165bcb-8951-4dfc-8c70-7fab5d7847f7"},
		"x-csrf-token":       {"75f6c9fa-dc8e-4e52-a000-e09dd4084b3e"},
		"x-money-object":     {"true"},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response failed: %w", err)
	}

	var response VintedResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("JSON unmarshal failed: %w", err)
	}

	return response.Items, nil
}

func PrintItems(items []VintedItem) {
	for _, item := range items {
		fmt.Printf("ID: %d\n", item.ID)
		fmt.Printf("Title: %s\n", item.Title)
		fmt.Printf("URL: %s\n", item.URL)
		fmt.Printf("Status: %s\n", item.Status)
		fmt.Printf("Size: %s\n", item.SizeTitle)
		fmt.Printf("User: %s\n", item.User.Login)
		fmt.Printf("Photo: %s\n", item.Photo.URL)
		fmt.Printf("Price: %s %s\n", item.Price.Amount, item.Price.CurrencyCode)
		fmt.Println("--------------------------------------------------")
	}
}
