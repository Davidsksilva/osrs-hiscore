package main

import (
	"log"
	"net/url"
	"os"
)

func getQueriedURL(playerName string) string {

	base, err := url.Parse(baseHiscoreURL)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	params := url.Values{}
	params.Add("player", playerName)
	base.RawQuery = params.Encode()

	return base.String()
}
