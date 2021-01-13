package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func task() {
	moedaType, cotacao := c.ClientRequest("btc")

	fmt.Println(moedaType, cotacao)
}

func main() {

	s1 := gocron.NewScheduler(time.UTC)

	s1.Every(3).Minutes().Do(task)
	s1.StartBlocking()
}
