// A Go wapper for the connpass API
package connpass

import (
        "encoding/json"
        "net/http"
        "net/url"
)

const EndPoint = "http://connpass.com/api/v1/event/"

type Search struct {
        Results_returned   int64    `json:"results_returned"`
        Results_available  int64    `json:"results_available"`
        Results_start      int64    `json:"results_start"`
        Events             []Events `json:"events"`
}

type Events struct {
        Event_id            int64    `json:"event_id"`
        Title               string   `json:"title"`
        Catch               string   `json:"catch"`
        Description         string   `json:"description"`
        Event_url           string   `json:"event_url"`
        Hash_tag            string   `json:"hash_tag"`
        Started_at          string   `json:"started_at"`
        Ended_at            string   `json:"ended_at"`
        Limit               int64    `json:"limit"`
        Event_type          string   `json:"event_type"`
        Series              Series   `json:"series"`
        address             string   `json:"address"`
        Place               string   `json:"place"`
        Lat                 float64  `json:"lat"`
        Lon                 float64  `json:"lon"`
        Owner_id            int64    `json:"owner_id"`
        Owner_nickname      string   `json:"owner_nickname"`
        Owner_display_name  string   `json:"owner_display_name"`
        Accepted            int64    `json:"accepted"`
        Waiting             int64    `json:"waiting"`
        Updated_at          string   `json:"updated_at"`
}

type Series struct {
        Id      int64  `json:"id"`
        Title   string `json:"title"`
        Url     string `json:"url"`
}

type client http.Client

func NewClient() *client {
        return &client{}
}

func (c *client) Search(keyword string) ([]Events, error) {
        p := url.Values{}
        p.Set("keyword", keyword)

        res, err := http.Get(EndPoint + "?" + p.Encode()) 
        if err != nil {
                return nil, err
        }
        defer res.Body.Close()

        var events struct {
                Event []Events `json:"events"`
        }
        err = json.NewDecoder(res.Body).Decode(&events)
        if err != nil {
                return nil, err
        }
        return events.Event, nil
}

