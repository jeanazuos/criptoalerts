package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type moeda struct {
	Ticker map[string]interface{} `json:"Ticker"`
}

func ClientRequest(moeda_type string) (string, map[string]interface{}) {
	endpoint := "https://www.mercadobitcoin.net/api/" + moeda_type + "/ticker/"

	client := http.Client{
		Timeout: time.Second * 2, //Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := client.Do(req)
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

	// fmt.Println(btc.Ticker, res.Status)
	return moeda_type, btc.Ticker

}
