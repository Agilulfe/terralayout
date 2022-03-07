# Welcome to TerraLayout

This is a small command-line utility that allows you to bootstrap a Terraform project.

It does the following: 

    * Create the different directories you need (following the file layout best practices).
    * Create the different files needed to follow Terraform's best practices.
    * Initiliaze git if you need it.
    * Stamp the Terraform version of your choice (above or equal to 0.13.7).
    * Fetch the latest provider version of your choice (eg. aws, azurerm or google).

# How to use it ?

You should have Go installed on your machine (version >= 1.17) on your machine if you want to compile the code yourself and run the command `go build .`.
It is also possible to compile the code for a specific OS/Arch using the following command
    `env GOOS=linux GOARCH=arm go build -o linux-arm ` for Linux/ARM system for example.

I have created a few binaries (for the most common Os/Arch) under the binaries directories. Feel free to use one that corresponds to your machine.


# ToDo

    * Personalise backend.tf file (generate remote backend code depending on the cloud provider)
    * Write unit test code (*sigh*)
    * Think about more features
