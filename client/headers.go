package client

import (
	"fmt"
	"net/http"
)

var (
	AccessToken  = "eyJraWQiOiJFNTdZZHJ1SHBsQWp1MmNObzFEb3JIM2oyN0J1NS1zX09QNVB3UGlobjVNIiwiYWxnIjoiUFMyNTYifQ.eyJhcHBfaWQiOjQsImF1ZCI6ImZyLmNvcmUuYXBpIiwiY2xpZW50X2lkIjoid2ViIiwiZXhwIjoxNzU4OTA1NTUzLCJpYXQiOjE3NTg4OTgzNTMsImlzcyI6InZpbnRlZC1pYW0tc2VydmljZSIsInB1cnBvc2UiOiJhY2Nlc3MiLCJzY29wZSI6InB1YmxpYyIsInNpZCI6IjlhNDkwMmM4LTE3NTg0NjgzMTUifQ.xlJbOEuB-RQzcVEkq1MU5NqyMpIt_m258iKggCxf3ooz6xCPqKThD2jpFc3iv6GBcth2TqpzKJsNYFl9djLaq_ucyZpq1eKymuRxBvWmcB3yrU6qhv7SMJ6vXKnN4BXAqKMYWiKWPU4zx7SjhhglA--MNPBQPld_IS13z-Q3boOpvgaUDDONIdjOnRv4P5lPZaA4dLt7dM0CUz0iYwl55d7sAEegVIGJGOuqYIzzU7PsAMCRDoMvSzLPz98jkMk22qfBgXxqbDQipZPmLot5L43T5eE1_XdmYqqCCbBWV0x77AMSl1a7cXBkZA9PXLmw-ytwDJk25T5k3k3GKFUMTQ"
	RefreshToken = "eyJraWQiOiJFNTdZZHJ1SHBsQWp1MmNObzFEb3JIM2oyN0J1NS1zX09QNVB3UGlobjVNIiwiYWxnIjoiUFMyNTYifQ.eyJhcHBfaWQiOjQsImF1ZCI6ImZyLmNvcmUuYXBpIiwiY2xpZW50X2lkIjoid2ViIiwiZXhwIjoxNzU5NTAzMTUzLCJpYXQiOjE3NTg4OTgzNTMsImlzcyI6InZpbnRlZC1pYW0tc2VydmljZSIsInB1cnBvc2UiOiJyZWZyZXNoIiwic2NvcGUiOiJwdWJsaWMiLCJzaWQiOiI5YTQ5MDJjOC0xNzU4NDY4MzE1In0.tzhmsy6g28lDhntlLadiUu80L7_PgOtk_oogX3bAHKl6VhHlJ0_8p3Q1NyjiYNflKD4nYWtXLa7Ktb_9dFEvLq8mxoNrw1fKL1c-EolY7ICDPO5wvpf9Kz1XHVTFQOPJKutWRsSghAyKKFOu-0vMLDKo4NzxvMHuku7nNq6RSntm6B_KGU9EZMADpirt9_iGmMmgztN0vVa6fLTNUTHiTQfjMjVHXOrMV_EvXiNIiR-RabIEWXXz6xBZEl-lJZwNUsv42b-DshYpat6vPFHRra-Py5Sid9-8G0SSi7ODYCNNrrJ8xUCar_qatCiQAmcsHgr1w4G34cvqfB-9WBKlRg"
)

func GetCookiesString() string {
	return fmt.Sprintf("access_token_web=%s; refresh_token_web=%s; cf_clearance=5VfTn3zuMEnMf_r2DBWmilQ6fr9wxJVIJOaSwPuXNYc-1759380868-1.2.1.1-Sio_uFiYh8Oo0athtPtlW0BtS.OKNaMfldgbaOw5mUS3kIV7fKxuij4s29Li3HVsfKLQb41wpNNd2.YEFRHPL4XTYAeuOcAyFJy3mX9btOAlY64cVcjss.zR6_vzXqMawydpPqORefEqrQTgr.CdKdzGZArd.GuuBUzviLTiK9N4SS11XZMnN6bLPWoVBEwIy8Ms6OnkrkqSWgbe5LwW4OCdyCFV9UeBYswVbJmYlxU;", AccessToken, RefreshToken)
}

var Headers = http.Header{
	"Accept":          {"application/json, text/plain, */*"},
	"Accept-Language": {"en-uk-fr"},
	"Connection":      {"keep-alive"},
	"Host":            {"www.vinted.co.uk"},
	"Priority":        {"u=3"},
	"Referer":         {"https://www.vinted.co.uk/catalog?search_text=ugg%20tasman&page=1&time=1758555430&currency=GBP&order=newest_first"},
	"Sec-Fetch-Dest":  {"empty"},
	"Sec-Fetch-Mode":  {"cors"},
	"Sec-Fetch-Site":  {"same-origin"},
	"TE":              {"trailers"},
	"User-Agent":      {"Mozilla/5.0 (X11; Linux x86_64; rv:143.0) Gecko/20100101 Firefox/143.0"},
	"X-Anon-Id":       {"93508af9-cb8e-4f01-8352-c333f659e573"},
	"X-CSRF-Token":    {"75f6c9fa-dc8e-4e52-a000-e09dd4084b3e"},
	"X-Money-Object":  {"true"},
	"Cookie":          {GetCookiesString()},
}
