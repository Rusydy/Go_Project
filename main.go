package main

import (
	nhiapi "NHI_Golang/nhiApi"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	//help benchmarking the request time
	now := time.Now()

	rosterFile, err := os.OpenFile("rosters.txt", os.O_RDWR|os.O_CREATE, 0666) // OpenFile need to receipt some flags

	if err != nil {
		log.Fatalf("error opening the file rosters.txt: %v", err)
	}

	defer rosterFile.Close()

	wrt := io.MultiWriter(os.Stdout, rosterFile)

	log.SetOutput(wrt)

	teams, err := nhiapi.GetAllTeams()
	if err != nil {
		log.Fatalf("error while getting all teams: %v", err)
	}

	var wg sync.WaitGroup

	wg.Add(len(teams))

	//unbuffered channel
	result := make(chan []nhiapi.Roster)

	for _, team := range teams {
		go func(team nhiapi.Team) {
			roster, err := nhiapi.GetRosters(team.ID)

			if err != nil {
				log.Fatalf("error getting roster: %v", err)
			}

			result <- roster

			wg.Done()
		}(team)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	display(result)

	log.Printf("took %v", time.Now().Sub(now).String())
}

func display(result chan []nhiapi.Roster) {
	for r := range result {
		for _, ros := range r {
			log.Println("----------------")
			log.Printf("ID: %d\n", ros.Person.ID)
			log.Printf("Name: %s\n", ros.Person.FullName)
			log.Printf("Position: %s\n", ros.Position.Abbreviation)
			log.Printf("Jersey: %s\n", ros.JerseyNumber)
			log.Println("----------------")
		}
	}
}
