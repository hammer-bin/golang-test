package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/go-resty/resty/v2"
	"golang.org/x/net/html/charset"
	"log"
	"os"
)

type ProductSearchResponse struct {
	Req Products `xml:"Products"`
}

type Products struct {
	Arg     int       `xml:"TotalCount"`
	Product []Product `xml:"Product"`
}

type Product struct {
	ProductCode   string `xml:"ProductCode"`
	ProductName   string `xml:"ProductName"`
	ProductPrice  int64  `xml:"ProductPrice"`
	ProductImage  string `xml:"ProductImage"`
	Text1         string `xml:"Text1"`
	Text2         string `xml:"Text2"`
	SellerNick    string `xml:"SellerNick"`
	Seller        string `xml:"Seller"`
	SellerGrd     int    `xml:"SellerGrd"`
	Rating        int    `xml:"Rating"`
	SalePrice     int64  `xml:"SalePrice"`
	Delivery      string `xml:"Delivery"`
	DetailPageUrl string `xml:"DetailPageUrl"`
}

func getPost(client *resty.Client) {

	target := readTarget()

	for _, keyword := range target {
		var rst ProductSearchResponse

		_, err := client.R().SetQueryParams(map[string]string{
			"key":     "23a6d341bc2dfe5a99247681ebdac6f1",
			"apiCode": "ProductSearch",
			"keyword": keyword,
			"sortCd":  "L",
		}).SetHeader("Accept", "applicaiton/xml").SetAuthToken("23a6d341bc2dfe5a99247681ebdac6f1").
			SetResult(&rst).
			Get("http://openapi.11st.co.kr/openapi/OpenApiService.tmall")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Item Count: %d\n", rst.Req.Arg)

		//for _, product := range rst.Req.Product {
		//	fmt.Printf("ProductName: |%100s|%7d|%d|%v\n", product.ProductName, product.ProductPrice, product.SalePrice, product.Delivery)
		//}

		if rst.Req.Arg > 0 {
			product := rst.Req.Product[0]
			fmt.Printf("ProductName: |%s|%7d|%7d|%v\n", product.ProductName, product.ProductPrice, product.SalePrice, product.Delivery)
		}

	}

}

func readTarget() []string {
	getwd, err := os.Getwd()
	if err != nil {
		return nil
	}
	fmt.Println("getwd: ", getwd)
	path := "D:/gitHub/hammer-bin/golang-test/Resty/target.txt"
	file, _ := os.Open(path)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	rst := make([]string, 0)
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		rst = append(rst, fileScanner.Text())
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	//b1 := make([]byte, 100)
	//n1, err := file.Read(b1)
	//if err == nil {
	//	fmt.Printf("%d byte: %s\n", n1, string(b1))
	//}

	return rst

}

func main() {
	client := resty.New()
	//resp, err := client.R().EnableTrace().Get("http://example.org")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(resp)

	// Charset 처리
	// 처리하지 않을경우 SetResult 에서 결과값이 담기지 않는다.
	client.XMLUnmarshal = func(data []byte, v interface{}) error {
		decoder := xml.NewDecoder(bytes.NewReader(data))
		decoder.CharsetReader = charset.NewReaderLabel
		err := decoder.Decode(v)
		return err
	}

	getPost(client)
	naverPost(client)

}
