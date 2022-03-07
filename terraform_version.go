package main

import (
	"fmt"
	"log"
	"os"
)

var terraformVersion string

func chooseTerraformVersion() {
	fmt.Println("Which Terraform version would you like to use? Please choose any version >= 0.13.7")
	fmt.Scanf("%s", &terraformVersion)
	if terraformVersion < "0.13.7" {
		log.Fatalln("Please choose any Terraform version >= 0.13.7")
	}
}

func writeVersionsFile(f os.File) {
	file := fmt.Sprintf(`terraform {
  required_version = "%s"
  required_providers {
    source  = "hashicorp/%s"
	  version = "~> %s"
  }
}
	`, terraformVersion, cloudProvider, providerVersion)
	_, err := f.WriteString(file)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
