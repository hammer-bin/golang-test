package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Gcp struct {
	Version   string     `json:"version"`
	Resources []Resource `json:"resources"`
}

type Resource struct {
	Mode      string     `json:"mode"`
	Type      string     `json:"type"`
	Name      string     `json:"name"`
	Instances []Instance `json:"instances"`
}

type Instance struct {
	SchemaVersion string    `json:"schema_version"`
	Attributes    Attribute `json:"attributes"`
}

type Attribute struct {
	InstanceId       string         `json:"instance_id"`
	Name             string         `json:"name"`
	NetworkInterface []NetInterface `json:"network_interface"`
}

type NetInterface struct {
	AccessConfig []AccessConfig `json:"access_config"`
}

type AccessConfig struct {
}

func main() {
	path, _ := os.Getwd()
	fmt.Println("Current Directory : ", path)
	// Open our jsonFile
	jsonFile, err := os.Open("./test01/GCP/gcp.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var gcp Gcp

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &gcp)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(gcp.Resources); i++ {
		fmt.Println("User Type: " + gcp.Resources[i].Type)
		fmt.Println("User Age: " + gcp.Resources[i].Name)
	}
}
