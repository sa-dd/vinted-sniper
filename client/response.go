package client


type Thumbnail struct {
	Url string `json:"url"`
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
	Url string `json:"url"`
}

type User struct {
	Name string `json:"login"`
	Pic Photo `json: "photo"`
}

type Item struct {
	Id            int        `json:"id"`
	Url           string 	 `json:"url"`
	FavoriteCount int        `json:"favourite_count"`
	Photo         Photo      `json:"photo"`
	Price         Price      `json:"price"`
	ServiceFee    ServiceFee `json:"service_fee"`
	Status        string     `json:"status"`
	Title         string     `json:"title"`
	SizeTitle     string     `json:"size_title"`
	User 	  	User 	`json: "user"`
}

type Response struct {
	Items []Item `json:"items"`
}
