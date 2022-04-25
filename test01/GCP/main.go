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
	NetworkIp    string         `json:"network_ip"`
}

type AccessConfig struct {
	NatIp string `json:"nat_ip"`
}

type TotalInstance struct {
	inst []*ResourceInfo
}
type ResourceInfo struct {
	Name      string
	NetworkIp string
	NatIp     string
}

func (t *TotalInstance) addInstance(name, net_ip, nat_ip string) {
	r := &ResourceInfo{Name: name, NetworkIp: net_ip, NatIp: nat_ip}
	t.inst = append(t.inst, r)
}

func (t *TotalInstance) printInstance() {
	fmt.Println("-------[Instance Information]------")
	for _, value := range t.inst {
		fmt.Printf("%v\n", *value)
		fmt.Printf("%s\n", value)
		fmt.Printf("Instance Name: %s\nNetwork IP: %s\nNAT IP: %s\n", value.Name, value.NetworkIp, value.NatIp)
	}
	fmt.Println("-----------------------------------")
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

	q1 := TotalInstance{}
	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(gcp.Resources); i++ {
		var instanceName, networkIp, natIp string
		for j := 0; j < len(gcp.Resources[i].Instances); j++ {
			instanceName = gcp.Resources[i].Instances[j].Attributes.Name
			fmt.Println("name: ", instanceName)
			for l := 0; l < len(gcp.Resources[i].Instances[j].Attributes.NetworkInterface); l++ {
				networkIp = gcp.Resources[i].Instances[j].Attributes.NetworkInterface[l].NetworkIp
				fmt.Println("network_ip: ", networkIp)
				for t := 0; t < len(gcp.Resources[i].Instances[j].Attributes.NetworkInterface[l].AccessConfig); t++ {
					natIp = gcp.Resources[i].Instances[j].Attributes.NetworkInterface[l].AccessConfig[t].NatIp
					fmt.Println("nat_ip: ", natIp)
				}
			}
		}
		q1.addInstance(instanceName, networkIp, natIp)
	}
	q1.printInstance()

}
