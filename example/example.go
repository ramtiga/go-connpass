package main

import (
        "github.com/ramtiga/go-connpass"
        "log"
        "fmt"
)

func main() {
        client := connpass.NewClient()
        results, err := client.Search("ruby")
        if err != nil {
                log.Fatal(err)
        }

        for _, res := range results {
                fmt.Println(res.Title)
        }
}
