package main

func main() {
	//get list of repositories
	//loop through them
	//get list of branches
	//if branch is older than 1 hour delete & log
	//add skip message to log if branch doesn't need deleting
	//write go test for testing this

	// if err != nil {
	// 	log.Fatalln(err)
	// }

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

// func auth() (bool, error) {

// 	req, err := http.NewRequest("GET", api, nil)

// 	if err != nil {
// 		log.Fatalln(err)
// 		return false, nil
// 	}

// 	//basic auth with personal token
// 	req.SetBasicAuth(g, os.Getenv("STALE_BRANCH_TOKEN"))

// 	resp, err := http.DefaultClient.Do(req)

// 	if err != nil {
// 		log.Fatalln(err)
// 		return false, nil
// 	}

// 	defer resp.Body.Close()

// 	if err != nil {
// 		log.Fatalln(err)
// 		return false, nil
// 	}

// 	return true, nil
// }
