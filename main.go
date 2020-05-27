package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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
