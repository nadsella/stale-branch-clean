package main

import (
	"fmt"
	"log"
	"os"
	"stale-branch-clean/config"
)

func main() {
	//get list of repositories
	//loop through them
	//get list of branches
	//if branch is older than 1 hour delete & log
	//add skip message to log if branch doesn't need deleting
	//write go test for testing this

	//fot now we just want to grab endpoints available for bb
	endpoints, err := config.GetEndpoints("bitbucket")

	if err != nil {
		log.Fatal(err)
	}


}
