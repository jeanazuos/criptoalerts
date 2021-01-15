package cotacao

import "github.com/jeanazuos/criptoalerts/client"

type Endpoint struct {
	Url string
}

func CriptoRequest(moeda_type string) []byte {
	var criptoEndpoint *Endpoint
	criptoEndpoint = new(Endpoint)
	criptoEndpoint.Url = "https://www.mercadobitcoin.net/api/" + moeda_type + "/ticker/"

	cotacao := client.ClientRequest(criptoEndpoint)
	return cotacao
}
