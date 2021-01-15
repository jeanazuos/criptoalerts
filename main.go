package main

import (
	"fmt"
	"time"

	"github.com/jeanazuos/criptoalerts/cotacao"

	"github.com/go-co-op/gocron"
)

func task() {
	bitcoin := cotacao.CriptoRequest("btc")

	fmt.Println(bitcoin)
}

func main() {

	s1 := gocron.NewScheduler(time.UTC)

	s1.Every(3).Minutes().Do(task)
	s1.StartBlocking()
}
