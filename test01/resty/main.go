package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
	"time"
)

func get() {
	client := resty.New()
	resp, err := client.R().SetQueryParams(map[string]string{
		"page_no": "1",
		"limit":   "20",
		"sort":    "name",
		"order":   "asc",
		"random":  strconv.FormatInt(time.Now().Unix(), 10),
	}).SetHeader("Accept", "application/json").Get("http://localhost:8080/")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}

func basic() {
	client := resty.New()
	resp, err := client.R().EnableTrace().Get("http://example.org")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode())
	fmt.Println(resp.Status())
	fmt.Println(resp.Proto())
	fmt.Println(resp.Time())
	fmt.Println(resp.ReceivedAt())

	ti := resp.Request.TraceInfo()
	fmt.Println(ti.DNSLookup)
	fmt.Println(ti.ConnTime)
	fmt.Println(ti.TCPConnTime)
	fmt.Println(ti.TLSHandshake)
	fmt.Println(ti.ServerTime)
	fmt.Println(ti.ResponseTime)
	fmt.Println(ti.TotalTime)
	fmt.Println(ti.IsConnReused)
	fmt.Println(ti.IsConnWasIdle)
	fmt.Println(ti.ConnIdleTime)
	fmt.Println(ti.RequestAttempt)
	fmt.Println(ti.RemoteAddr.String())
}
func main() {
	get()
}
