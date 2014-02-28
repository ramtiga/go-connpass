package main

import (
        "github.com/ramtiga/go-connpass"
        "log"
        "fmt"
)

func main() {
        client := connpass.NewClient()

        params := map[string]interface{}{
                "keyword": "ruby",
                "ym": 201403,
        }
        results, err := client.Search(params)
        if err != nil {
                log.Fatal(err)
        }

        for _, res := range results {
                fmt.Println(res.Title)
        }
}
