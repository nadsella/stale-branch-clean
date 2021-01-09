package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	//get list of repositories
	//loop through them
	//get list of branches
	//if branch is older than 1 hour delete & log
	//add skip message to log if branch doesn't need deleting
	//write go test for testing this

	login()
	// resp, err := http.Get("https://httpbin.org/get")

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// log.Println(string(body))
}

func login() (bool, error) {
	u := fmt.Sprintf("nadsella:%s", os.Getenv("STALE_BRANCH_TOKEN"))
	requestBody, err := json.Marshal(u)

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://api.github.com/users/nadsella", "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

	return true, nil
}
