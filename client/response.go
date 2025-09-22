package client


type Thumbnail struct {
	URL string `json:"url"`
}

type Price struct {
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currency_code"`
}

type ServiceFee struct {
	Amount string `json:"amount"`
}

type Photo struct {
	Thumbnails []Thumbnail `json:"thumbnails"`
}

type Item struct {
	Id            int        `json:"favourite_count"`
	FavoriteCount int        `json:"favourite_count"`
	Photo         Photo      `json:"photo"`
	Price         Price      `json:"price"`
	ServiceFee    ServiceFee `json:"service_fee"`
	Status        string     `json:"status"`
	Title         string     `json:"title"`
	SizeTitle     string     `json:"size_title"`
}

type Response struct {
	Items []Item `json:"items"`
}


