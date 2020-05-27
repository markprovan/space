package main

import (
	"fmt"
	"time"
)

// Launch represents an individual launch from https://launchlibrary.net/1.4/launch
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
	fmt.Println("🚀", l.Name, "⏰ ", localTime, streamLink)
}
