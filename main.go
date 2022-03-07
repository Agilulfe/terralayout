package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var filesList = []string{"backend.tf", "main.tf", "outputs.tf", "provider.tf", "variables.tf", "versions.tf"}

func main() {
	fmt.Println("Welcome to TerraLayout!")
	name, location, folders, git := getUserInput()
	chooseTerraformVersion()
	getLatestProviderVersion()
	createDirectories(name, location, folders, git)
}

func getUserInput() (string, string, []string, string) {
	fmt.Println("What is the name of your project/repository?")
	var name string
	fmt.Scanf("%s", &name)

	fmt.Println("Would you like to initialize git? Please answer with 'yes' or 'no'")
	var git string
	fmt.Scanf("%s", &git)

	fmt.Println("Where do you want to set up your project? Keep empty for current directory")
	var location string
	fmt.Scanf("%s", &location)

	fmt.Println("Please provide your different environments, separated by commas without spaces (eg. dev,sit,uat,prod)")
	var envs string
	fmt.Scanf("%s", &envs)
	folders := strings.Split(envs, ",")

	providers()

	return name, location, folders, git
}

func createDirectories(name string, location string, folders []string, git string) {
	var path string
	if len(location) > 0 {
		path = location + "/" + name + "/"
	} else {
		path = name + "/"
	}

	err := os.Mkdir(path, 0755)
	if err != nil {
		log.Fatalln(err)
	}

	if git == "yes" {
		cmd := exec.Command("git", "init")
		cmd.Dir = path
		_, err := cmd.Output()
		if err != nil {
			log.Fatalln(err)
		}
	}

	for i := 0; i < len(folders); i++ {
		layout := path + folders[i]
		err := os.Mkdir(layout, 0755)
		if err != nil {
			log.Fatalln(err)
		}
		createFiles(layout)
	}
}

func createFiles(layout string) {
	for i := 0; i < len(filesList); i++ {
		filePath := layout + "/" + filesList[i]
		file, e := os.Create(filePath)
		if e != nil {
			log.Fatalln(e)
		}
		if filesList[i] == "versions.tf" {
			writeVersionsFile(*file)
		}
		file.Close()
	}
}
