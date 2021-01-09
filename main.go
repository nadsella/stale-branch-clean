package main

import (
	"fmt"
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

	l, err := auth()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%v", l)
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

func auth() (bool, error) {
	g := os.Getenv("GITHUB_USERNAME")
	api := fmt.Sprintf("%s/%s", "https://api.github.com/users", g)

	req, err := http.NewRequest("GET", api, nil)

	if err != nil {
		log.Fatalln(err)
		return false, nil
	}

	//basic auth with personal token
	req.SetBasicAuth(g, os.Getenv("STALE_BRANCH_TOKEN"))

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalln(err)
		return false, nil
	}

	defer resp.Body.Close()

	if err != nil {
		log.Fatalln(err)
		return false, nil
	}

	return true, nil
}
