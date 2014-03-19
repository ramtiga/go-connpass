# go-connpass

Interface to connpass API written by golang.

## Usage:

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


## Install:

    go get github.com/ramtiga/go-connpass

## License:

MIT

## Author:

ramtiga
