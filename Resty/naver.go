package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

func naverPost(client *resty.Client) {
	target := readTarget()

	for _, keyword := range target {
		var rst NaverResponse

		_, err := client.R().SetQueryParams(map[string]string{
			"query":   keyword,
			"sort":    "sim",
			"exclude": "used:rental",
		}).
			SetHeader("X-Naver-Client-Id", "GON0C_n4Q47oj4xZbI2B").
			SetHeader("X-Naver-Client-Secret", "CMEvzkYqJh").
			SetResult(&rst).
			Get("https://openapi.naver.com/v1/search/shop.json")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Item Count: %d\n", rst.Total)

		if rst.Total > 0 {
			product := rst.Items[0]
			fmt.Printf("ProductName: |%s|%7s|%7s|%v\n", product.Title, product.LPrice, product.MallName, product.Brand)
		}

	}
}
