package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"regexp"
)

type User struct {
	FirstName string `structs:"first_name"`
	LastName  string `structs:"last_name"`
}

func main() {
	a := User{FirstName: `${file("google.json")}`, LastName: "Wick"}
	s := structs.New(a)
	m := s.Map()

	jsonByte, err := json.Marshal(m) // map to JSON Format([]byte)
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile(`[\/\\]`)
	//key := re.ReplaceAllString(string(jsonByte), "")
	//strings.Replace(string(jsonByte), '\"', "123", -1)

	//fmt.Println(key)
	TfJson, _ := FormatJSONPretty(jsonByte)
	key := re.ReplaceAllString(string(TfJson), "")

	fmt.Println(string(key))
}

func FormatJSONPretty(data []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", "    ")
	if err == nil {
		return out.Bytes(), err
	}
	return data, nil
}
