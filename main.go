package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Player struct {
	Name string `json:"name"`
	Link string `json:"link"`
	Active string `json:"active"`
}

const (
	url = "https://www.basketball-reference.com/players/"
	html = ".html"
)

func CreatePlayerList(data [][]string) []Player {
	var players []Player
	for i, line := range data {
		if i > 0 { // omit header line
				var player Player
				for j, field := range line {
					switch j {
					case 0:
						player.Name = url + string(field[0]) + "/" + field + html
					case 1:
						player.Link = field
					case 2:
						player.Active = field		
					}
				}
				players = append(players, player)
		}
}
	return players
}

func main() {
    // open file
    f, err := os.Open("nba_players.csv")
    if err != nil {
        log.Fatal(err)
    }

    // remember to close the file at the end of the program
    defer f.Close()

    // 2. Read CSV file using csv.Reader
    csvReader := csv.NewReader(f)
    data, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    // 3. Assign successive lines of raw CSV data to fields of the created structs
    playerList := CreatePlayerList(data)

    // 4. Convert an array of structs to JSON using marshaling functions from the encoding/json package
    jsonData, err := json.MarshalIndent(playerList, "", "  ")
    if err != nil {
        log.Fatal(err)
    }

    err = ioutil.WriteFile("nba_players.json", jsonData, 0644)
		if err != nil {
			log.Fatal(err)
	}
}