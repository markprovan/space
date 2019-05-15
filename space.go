package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Launch struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Date    string   `json:"net"`
	Streams []string `json:"vidURLs"`
}

func (l Launch) print() {

	format := "Jan 2, 2006 15:04:05 MST"
	parsedTime, _ := time.Parse(format, l.Date)

	streamLink := ""

	if len(l.Streams) > 0 {
		streamLink = "📺 " + l.Streams[0]
	}

	localTime := parsedTime.Local().Format("Jan _2 15:04 MST 2006")
	fmt.Println("🚀", l.Name, "⏰ ", localTime, streamLink, "\n")
}

type LaunchResponse struct {
	Launches []Launch `json:"launches"`
}

func main() {
	resp, err := http.Get("https://launchlibrary.net/1.4/launch")

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var response LaunchResponse
	json.Unmarshal(body, &response)

	for _, launch := range response.Launches {
		launch.print()
	}
}
