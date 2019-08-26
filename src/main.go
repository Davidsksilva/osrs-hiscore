package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func fetchHiscore(playerName string) {

	var url = getQueriedURL(playerName)
	resp, err := http.Get(url)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	splittedString := strings.Split(string(body), "\n")

	fmt.Printf("Player %v \n", playerName)
	for index, value := range skills {
		skillValues := strings.Split(splittedString[index], ",")

		fmt.Printf("%v: level %v with %vxp ranked %v \n", value, skillValues[0], skillValues[1], skillValues[2])
	}

}

func main() {

	fetchHiscore(os.Args[1])
}
