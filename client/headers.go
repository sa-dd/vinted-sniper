package client

import (
	"fmt"
	"net/http"
)

var (
	AccessToken  = "eyJraWQiOiJFNTdZZHJ1SHBsQWp1MmNObzFEb3JIM2oyN0J1NS1zX09QNVB3UGlobjVNIiwiYWxnIjoiUFMyNTYifQ.eyJhcHBfaWQiOjQsImF1ZCI6ImZyLmNvcmUuYXBpIiwiY2xpZW50X2lkIjoid2ViIiwiZXhwIjoxNzU5NTc4NDcyLCJpYXQiOjE3NTk1NzEyNzIsImlzcyI6InZpbnRlZC1pYW0tc2VydmljZSIsInB1cnBvc2UiOiJhY2Nlc3MiLCJzY29wZSI6InB1YmxpYyIsInNpZCI6IjhiYTlhODM4LTE3NTk1NzEyNzIifQ.ybHspUNdQ7f5o2TTJ7jLTqVJsDbcZEl77Z7Alyp_k_TqeBwa64aKpmQTvpX3nKkFnvp24S8dY6N1frYUM3Hs-tG13d0JUeqC_cYz3aPdjtszHxQBFAgxKBHG5IUl7rBLwYziFRX7ru91_yebhSKVhvKRM1b9eYzmimO1RKsKmTy3ea1xBK8qZdPSX0Kt_4yUvSuFwgGcHYKtHz8-MwgT6-HLId3AF8qb6kXfiTTHuPpDJl2SfIVylggaKdVPNkJnHaFKJMJQX0rJzWHpkOgcirrWdhTGn7nl8o4V0SoaXhLOLs7ejXGZwE8LmFTxuKAkj8lhNHe41X3ErRTDZmPlLA"
	RefreshToken = "eyJraWQiOiJFNTdZZHJ1SHBsQWp1MmNObzFEb3JIM2oyN0J1NS1zX09QNVB3UGlobjVNIiwiYWxnIjoiUFMyNTYifQ.eyJhcHBfaWQiOjQsImF1ZCI6ImZyLmNvcmUuYXBpIiwiY2xpZW50X2lkIjoid2ViIiwiZXhwIjoxNzYwMTc2MDcyLCJpYXQiOjE3NTk1NzEyNzIsImlzcyI6InZpbnRlZC1pYW0tc2VydmljZSIsInB1cnBvc2UiOiJyZWZyZXNoIiwic2NvcGUiOiJwdWJsaWMiLCJzaWQiOiI4YmE5YTgzOC0xNzU5NTcxMjcyIn0.WHNi4UaFSGADgTFaaRv0V_6gck1FA_qWlq5zqv6OXwjUT3VM3FXvGxBzcOVSDNrb_hHyluqF2ckU5qyNeZTvdyDkQWfHPEENPSF7Z4W_yCpnhnZyHoa15mURBSjPxs2Wj6VYYO7_GVLiO2RW4rNyCnReQKAzcpMeFVdBdmGqYnOw56svp562oFbio9xAihwLoI6p_F1xdQRnTVvnPe8unQhhHLmwruqJCjS1tQusELTink8PTwMRg-aeC5bkWVgYM9SP0C81zPAkUxFOExe2OcOSSCK1FFTm8Ow87BMp4RIKmPR-Zottyui-JpihSXpAVgu9Qoo_0TtPkiL-iz_Uzg"
)

func GetCookiesString() string {
	return fmt.Sprintf("access_token_web=%s; refresh_token_web=%s; cf_clearance=uLHZzmAvH.7gIT3dkAsWgzGQgb6tYbTuTy2_o5_bmTA-1759571276-1.2.1.1-tNctnl9h0kBSxYEaArfC.nIgi3VkZfPCWTqzQMjKPPHwABfS.RHATlMt1h3Z28OeoPaJBRGUe6.sD.HFjAVHelfY6h.VpE2kwRyu0BNP3qutafBGuRtAh6dSzRV_jKLVuRS5_dx8079kqSkjuMHBqm1VU3dj95XdMpjWbKi9oZCwirAxX0ZgJ_rEu8dqEZii58jH.s3tBiwMugKvAsbWbpbfxvOSUW5eX05UAQVUBPo;", AccessToken, RefreshToken)
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
	"X-Anon-Id":       {"def3ceb9-7e71-4e7b-a221-5cd29013b0bd"},
	"X-CSRF-Token":    {"75f6c9fa-dc8e-4e52-a000-e09dd4084b3e"},
	"X-Money-Object":  {"true"},
	"Cookie":          {GetCookiesString()},
}
