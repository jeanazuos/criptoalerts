package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
)

type moeda struct {
	Ticker map[string]interface{} `json:"Ticker"`
}

func task() map[string]interface{} {
	fmt.Println("I am running task.")
	url := "https://www.mercadobitcoin.net/api/BTC/ticker/"

	spaceClient := http.Client{
		Timeout: time.Second * 2, //Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	btc := moeda{}
	fmt.Println(string(body))

	jsonErr := json.Unmarshal(body, &btc)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// fmt.Println(people1.Number, res.Status)
	fmt.Println(btc.Ticker, res.Status)
	return btc.Ticker
}

func main() {

	s1 := gocron.NewScheduler(time.UTC)

	s1.Every(3).Minutes().Do(task)
	s1.StartBlocking()
}
