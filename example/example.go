package main

import (
	"fmt"
	"github.com/ramtiga/go-connpass"
	"log"
)

func main() {
	client := connpass.NewClient()

	params := map[string]interface{}{
		"keyword": "ruby",
		"ym":      201403,
	}
	results, err := client.Search(params)
	if err != nil {
		log.Fatal(err)
	}

	for _, res := range results {
		fmt.Println(res.Title)
	}
}
