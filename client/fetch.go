package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchItems(client *http.Client, url string) ([]Item, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	req.Header = http.Header{
		"authority":          {"www.vinted.co.uk"},
		"accept":             {"application/json, text/plain, */*"},
		"accept-language":    {"en-uk-fr"},
		"cache-control":      {"max-age=0"},
		"cookie":             {"v_udt=L2JGTWk0RHJmVWN2QUYzVDZRaXFrVkFGTFBXci0tejE0bzlkb09USjVucHNEZi0tSk40aXhySHV6U0pBMEpQdXBvNERTUT09; anon_id=2c165bcb-8951-4dfc-8c70-7fab5d7847f7; anonymous-locale=en-uk-fr; anon_id=60cbd6a4-9b5c-4d35-81bc-957fc2373e88; OptanonAlertBoxClosed=2025-09-20T20:34:47.490Z; eupubconsent-v2=CQYCsRgQYCsRgAcABBENB9FgAAAAAEPgAAwIAAAWZABMNCogjLIgBCJQMAIEACgrCACgQBAAAkDRAQAmDApyBgAusJkAIAUAAwQAgABBgACAAASABCIAKACAQAAQCBQABgAQBAQAMDAAGACxEAgABAdAxTAggECwASMyqDTAlAASCAlsqEEgCBBXCEIs8AggREwUAAAIABQEAADwWAhJICViQQBcQTQAAEAAAUQIECKRswBBQGaLQXgyfRkaYBg-YJklMgyAJgjIyTYhN-Ew8UhRAAAA.YAAACHwAAAAA; OTAdditionalConsentString=1~; domain_selected=true; v_sid=afac40c7-1758400481; __eoi=ID=b7a26f62587f6978:T=1758461422:RT=1758461422:S=AA-AfjYzxoJxDzH3g-xwt4eCnWch; monetixads-user-session=250101575373614000053736516982151230; sharedid=7650805c-a780-4c94-9ea4-3c70a6754b05; sharedid_cst=A8%2BRZw%3D%3D; __cf_bm=iDt_W7LmNTCOrkKXnZ2Cn8jMrxrLqWMvxUk71DVFxa0-1758468456-1.0.1.1-1qbGf73jz6txMtrsmngZOsosAbXaAsviBYp7g1Rc7gv3dUjl.Y9M_5aj7UZPWKClEHEcjwDBdGeWqSoH1pV2BTYJcd63HHkStzjVya3oue0BqZDUraw9VIj8r0G.NgIb; cf_clearance=6UUWFSLVVfY5jgfYouIdHDijE4794RbsKIPtgkYfPuE-1758468456-1.2.1.1-EDB_0OKP9KCJYzWEOJparGNZm3fExYeFQT01C.JYUwzB_oSqNbDC_W.LHi8KK2.YLwONfstFL.SZ8bExj6IxheY8.CAUJ2ZNXS_CWCFcmCkPYPtV8D4TONjDBGGt8c2ROe0FOQYR28OYiAtn4XTOEH1ht3auwsEqy20s3YhySHd9aZtopYG8SXzDeXsmQXSN8k5fKRElcfZVXZCebsjLZkrUcequpI_TH08y.U_QP_Y; access_token_web=eyJraWQiOiJFNTdZZHJ1SHBsQWp1MmNObzFEb3JIM2oyN0J1NS1zX09QNVB3UGlobjVNIiwiYWxnIjoiUFMyNTYifQ.eyJhcHBfaWQiOjQsImF1ZCI6ImZyLmNvcmUuYXBpIiwiY2xpZW50X2lkIjoid2ViIiwiZXhwIjoxNzU4NDc1NjU3LCJpYXQiOjE3NTg0Njg0NTcsImlzcyI6InZpbnRlZC1pYW0tc2VydmljZSIsInB1cnBvc2UiOiJhY2Nlc3MiLCJzY29wZSI6InB1YmxpYyIsInNpZCI6ImFmYWM0MGM3LTE3NTg0MDA0ODEifQ.Hrpphn3qsRrGU-Z1GZKJ14R8DuuR-Tu1SSzE0C3G1Bab9lvaZUcjKYj_khWbGMt4dfMR3U8mocDcdVY5FK4lIP8J2S1zPSIdDcTx-RUyC3c2gSIuB6sDQwbm_nkywFCa08tk9qf7sghsA3F_PRbfkU9ZX76hPByYl2V9Ond8U1frYhAITtMdgWmUGdviFrOf4pY9WWx1sYoREdUMITvhkhF1a5KTr5N3CjXB_oafIaeqbFZ1mGttRkMfqjqDovbbmFK4dxHA0H4vgnN6I2x-yAjg86M5ZFdryd3lVb4fAKkCH27NOFjjwkuhxHtb9NQgb8EG2-z8lCfuFml6P0WSgw; refresh_token_web=eyJraWQiOiJFNTdZZHJ1SHBsQWp1MmNObzFEb3JIM2oyN0J1NS1zX09QNVB3UGlobjVNIiwiYWxnIjoiUFMyNTYifQ.eyJhcHBfaWQiOjQsImF1ZCI6ImZyLmNvcmUuYXBpIiwiY2xpZW50X2lkIjoid2ViIiwiZXhwIjoxNzU5MDczMjU3LCJpYXQiOjE3NTg0Njg0NTcsImlzcyI6InZpbnRlZC1pYW0tc2VydmljZSIsInB1cnBvc2UiOiJyZWZyZXNoIiwic2NvcGUiOiJwdWJsaWMiLCJzaWQiOiJhZmFjNDBjNy0xNzU4NDAwNDgxIn0.sVbp1PO3FNIQmbaRoHL5h5R_JF-blE58Xh8KxFv_jt-io8qeXRzOqFcnz79EOEIaFKHQIb6K3lIsNfF6T9ViPjNycIJ5__aQtSkytSH2octKSDf73CTr3vGvlerZaHMrVj9ftY2qMFRDcfFsw9SGYp71s2G4PFIYLwuQ0Slu1EL6bgae0dbSbLvhK2_Gpi_O80Xky4Rgh7z9bLo3C5Tv3lELM-aZBeT1s5Ga0IIHuVKaggrMpxYoikqBlDT38VLthJZl-BrLvQnF4jiMdiV5IYluHQ6Vke1K_KMbhJDUMuUyoCJ3vGZ4s3btVE_OLmu5E7KPR6ybQinVCaV9BgR5pw; homepage_session_id=a72b43e9-41d4-4522-a29d-b0189ba24f8a; viewport_size=637; OptanonConsent=isGpcEnabled=0&datestamp=Sun+Sep+21+2025+16%3A27%3A50+GMT%2B0100+(British+Summer+Time)&version=202508.2.0&browserGpcFlag=0&isIABGlobal=false&consentId=2c165bcb-8951-4dfc-8c70-7fab5d7847f7&identifierType=Cookie+Unique+Id&isAnonUser=1&hosts=&interactionCount=1&landingPath=NotLandingPage&groups=C0001%3A1%2CC0002%3A0%2CC0003%3A0%2CC0004%3A0%2CC0005%3A0%2CV2STACK42%3A0%2CC0015%3A0%2CC0035%3A0%2CC0038%3A0&genVendors=V2%3A0%2CV1%3A0%2C&intType=2&geolocation=GB%3BENG&AwaitingReconsent=false; banners_ui_state=SUCCESS; datadome=O1GbaPzX9QYC4cUNlcDg8UWAMTHsBwEl4ps4G3jBoC7L~kRqzrfeM4SbEnS4s~QyA88s0bJImzHgRMlghcZIpxyH3OirO7XpNof8KQCQsHQPmHwBZOJlOI2~yQ14n5Id; _vinted_fr_session=MUgxRHZmVS9xWFNCRmJlYzVqNEJXT0dNaXlrUGhOSHNUYktvdUpLWWJmVXZyemVBWVJJVEdYODQwSTduUjRuejVGazdySDh4dkJ1OXFtQ0l2eEsyb21iVXN2Z1V6ejhVS0R2REx6RUxYWWI2eUFuOEg5VkpiQnlqa1VyVml2SGJEMU5lN2FXc0lmOTc0aXpDaENyeXRDVm9sejh3VkZDby9QVmpDNlV3RTdob0VBa0pSQVZJODZrdENwR2NsbkVZQmZQVFVLYkJab0JUMGE0NmZTWWhSSjVPTm55WXBuZUpFUm5lem1FOE9GOD0tLTNGN3NobmY1UHRuNEhwWnJWTHpVUFE9PQ%3D%3D--1ecac438f16154bb3b2bffe8fd227941fab6cc65"}, // Get cookie from env
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

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("JSON unmarshal failed: %w", err)
	}

	return response.Items, nil
}

func FindLatestItems(latestItemId int, items []Item) []Item {
	var latestItems []Item

	for _, item := range items {
		if item.ID == latestItemId {
			break
		}
		latestItems = append(latestItems, item)

	}

	return latestItems
}

func PrintItems(items []Item) {
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
