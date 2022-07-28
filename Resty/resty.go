package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/go-resty/resty/v2"
	"golang.org/x/net/html/charset"
)

type ProductSearchResponse struct {
	Req Products `xml:"Products"`
}

type Products struct {
	Arg     int     `xml:"TotalCount"`
	Product Product `xml:"Product"`
}

type Product struct {
	ProductCode   string `xml:"ProductCode"`
	ProductName   string `xml:"ProductName"`
	ProductPrice  string `xml:"ProductPrice"`
	ProductImage  string `xml:"ProductImage"`
	Text1         string `xml:"Text1"`
	Text2         string `xml:"Text2"`
	SellerNick    string `xml:"SellerNick"`
	Seller        string `xml:"Seller"`
	SellerGrd     int    `xml:"SellerGrd"`
	Rating        int    `xml:"Rating"`
	DetailPageUrl string `xml:"DetailPageUrl"`
}

func getPost(client *resty.Client) {

	var aaa ProductSearchResponse

	_, err := client.R().SetQueryParams(map[string]string{
		"key":     "23a6d341bc2dfe5a99247681ebdac6f1",
		"apiCode": "ProductSearch",
		"keyword": "놈모 크로마 게이밍",
	}).SetHeader("Accept", "applicaiton/xml").SetAuthToken("23a6d341bc2dfe5a99247681ebdac6f1").
		SetResult(&aaa).
		Get("http://openapi.11st.co.kr/openapi/OpenApiService.tmall")

	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(resp)

	fmt.Println(aaa.Req.Arg)
	fmt.Println(aaa.Req.Product.ProductCode)
	fmt.Println(aaa.Req.Product.ProductName)
	fmt.Println(aaa.Req.Product.SellerNick)

}

func main() {
	client := resty.New()
	resp, err := client.R().EnableTrace().Get("http://example.org")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

	// Charset 처리
	// 처리하지 않을경우 SetResult 에서 결과값이 담기지 않는다.
	client.XMLUnmarshal = func(data []byte, v interface{}) error {
		decoder := xml.NewDecoder(bytes.NewReader(data))
		decoder.CharsetReader = charset.NewReaderLabel
		err = decoder.Decode(v)
		return err
	}

	getPost(client)

	//client := resty.New()

}
