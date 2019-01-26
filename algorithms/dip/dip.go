package main

import (
	"fmt"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
)

func init() {
	fmt.Printf("Running w/ credentials [%v %v]\n", common.Credentials().ID, common.Credentials().Secret)
}

func main() {
	alpacaClient := alpaca.NewClient(common.Credentials())
	acct, err := alpacaClient.GetAccount()
	if err != nil {
		panic(err)
	}

	fmt.Println(*acct)
}
