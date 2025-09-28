package client

import (
	"fmt"
	"net/http"
)

var (
	AccessToken  = "eyJraWQiOiJFNTdZZHJ1SHBsQWp1MmNObzFEb3JIM2oyN0J1NS1zX09QNVB3UGlobjVNIiwiYWxnIjoiUFMyNTYifQ.eyJhY2NvdW50X2lkIjoyMjkzNjk5MDcsImFwcF9pZCI6NCwiYXVkIjoiZnIuY29yZS5hcGkiLCJjbGllbnRfaWQiOiJ3ZWIiLCJleHAiOjE3NTkwNzA5NjgsImlhdCI6MTc1OTA2Mzc2OCwiaXNzIjoidmludGVkLWlhbS1zZXJ2aWNlIiwibG9naW5fdHlwZSI6MywicHVycG9zZSI6ImFjY2VzcyIsInNjb3BlIjoidXNlciIsInNpZCI6IjE0ZTM4ZGFiLTE3NTg0MDMyMjYiLCJzdWIiOiIyODg0MDEzOTMiLCJjYyI6IlVLIiwiYW5pZCI6IjgyOGJmMjQ2LTljZWItNDY2Yi05MjBmLTlhYTdhMmM2YzkwNyIsImFjdCI6eyJzdWIiOiIyODg0MDEzOTMifX0.YKP8Pcda4-KZvzZbA9c8DzxXsR-85_bLxKVJMNWkLGEtszt1oycK9UdFRgrFWYbDl5Sc8Xj4K3MkmX7gBS5ibXF_NSJ1Q8Fn6VztvDnjpC3Ymxei0aG0vccSr5jRkm63G-DisgdfO1x1l2ghdi3YBkPEp3D6MDKvhhnJZOLAV6FUA3P9lUoclMnw03VEhpAukBM4dd_m-CfheeKXCsjnTWHmp8agidE1cFwLwqbIUigfDWc_ob7_1igrsG4gnoGLqgneWjkSb00rErnZOgpPbl-CZAM476e5nz4qpBY_SLfCEnPurIjLRGtA5Rzw4G2WKlFgGjmlPQ1xMq7uB5wu-g"
	RefreshToken = "eyJraWQiOiJFNTdZZHJ1SHBsQWp1MmNObzFEb3JIM2oyN0J1NS1zX09QNVB3UGlobjVNIiwiYWxnIjoiUFMyNTYifQ.eyJhY2NvdW50X2lkIjoyMjkzNjk5MDcsImFwcF9pZCI6NCwiYXVkIjoiZnIuY29yZS5hcGkiLCJjbGllbnRfaWQiOiJ3ZWIiLCJleHAiOjE3NTk2Njg1NjgsImlhdCI6MTc1OTA2Mzc2OCwiaXNzIjoidmludGVkLWlhbS1zZXJ2aWNlIiwibG9naW5fdHlwZSI6MywicHVycG9zZSI6InJlZnJlc2giLCJzY29wZSI6InVzZXIiLCJzaWQiOiIxNGUzOGRhYi0xNzU4NDAzMjI2Iiwic3ViIjoiMjg4NDAxMzkzIiwiY2MiOiJVSyIsImFuaWQiOiI4MjhiZjI0Ni05Y2ViLTQ2NmItOTIwZi05YWE3YTJjNmM5MDciLCJhY3QiOnsic3ViIjoiMjg4NDAxMzkzIn19.B7_LfXHK1mbEqgP8A9Fi7w7Nd4_WAtNCWdhMcUQhW7E2oZ7QQq2rqWynWw4Ez-53Xwopz6L8Y9OFGy2AJwNvj4JWvFGQXUPe8-OG5sXA_tvveo5j08eAtBL9X-LTYGBYSiqM-4MTNzrfMG_oNLYxFfslqlj8Io3zuc47uouU940EmmBB5xGmMdtYtrJNJnAMZqFzpkuVLosV3DPhbcyir8u4ItHGLV5kbIv7npv8j-b-zqZz7CTla_Tum6D5mGqbUGDoUZmg-lsO97RJe3BEX_iYmzmtyWgVsYlFLvgtMACkBNRP7J7dXug1B4usrIyIj8AKoO35u3DYSMMoh5lFhA"
)

func GetCookiesString() string {
	return fmt.Sprintf("access_token_web=%s; refresh_token_web=%s; cf_clearance=_pzjkEjDziR.c.2RJK5_rnYYh4Z6Gkl4qzneMW9TRCw-1758904946-1.2.1.1-rRsLuTI_ZSGkP4Afn9JpQ0sT5T0VV0PwmAs9uSD58TRQZZD1guoQf8Atd3ljpnSvQE9KZ9PMugu1Cd7594zWGUcYYJFLo6b9VAsUai09DEOCckbWh9KiMKkHQSJHCu.wscmfT1mmbvOX8jO3OgH78q4ZaDQW6c5NBSsg4kZQActi_v3aAl5kJPhvIzsUoSVQvqyCDs9KaRNDiw0D.K0SFq33UIRtcmiImWhKF9D46UU;", AccessToken, RefreshToken)
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
