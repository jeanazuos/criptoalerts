package cotacao

import (
	"encoding/json"
	"log"

	"github.com/jeanazuos/criptoalerts/client"
)

type Endpoint struct {
	Url string
}

type Ativo struct {
	Ticker map[string]interface{} `json:"Ticker"`
}

func CriptoRequest(moeda_type string) map[string]interface{} {
	var criptoEndpoint *Endpoint
	criptoEndpoint = new(Endpoint)
	criptoEndpoint.Url = "https://www.mercadobitcoin.net/api/" + moeda_type + "/ticker/"

	response := client.ClientRequest(string(criptoEndpoint.Url))
	var ativo *Ativo
	ativo = new(Ativo)
	jsonErr := json.Unmarshal(response, &ativo)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return ativo.Ticker
}
