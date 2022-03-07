package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var cloudProvider string
var providerVersion interface{}

func providers() {
	fmt.Println("Please provide a cloud provider you want to use (eg. aws, google, azurerm)")
	fmt.Scanf("%s", &cloudProvider)
}

func getLatestProviderVersion() {
	url := fmt.Sprintf("https://registry.terraform.io/v1/providers/hashicorp/%s", cloudProvider)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	responseObject := make(map[string]interface{})
	err = json.Unmarshal(body, &responseObject)
	if err != nil {
		log.Fatalln(err)
	}
	providerVersion = responseObject["version"]
}
