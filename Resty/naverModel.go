package main

type NaverResponse struct {
	LastBuildDate string  `json:"lastBuildDate"`
	Total         int     `json:"total"`
	Start         int     `json:"start"`
	Display       int     `json:"display"`
	Items         []Items `json:"items"`
}

type Items struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Image       string `json:"image"`
	LPrice      string `json:"lprice"`
	HPrice      string `json:"hprice"`
	MallName    string `json:"mallName"`
	ProductId   string `json:"productId"`
	ProductType string `json:"productType"`
	Brand       string `json:"brand"`
	Maker       string `json:"maker"`
	Category1   string `json:"category1"`
	Category2   string `json:"category2"`
	Category3   string `json:"category3"`
	Category4   string `json:"category4"`
}
