package main

import (
	nhiapi "NHI_Golang/nhiApi"
	"io"
	"log"
	"os"
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

	for _, team := range teams {
		log.Println("------------------")
		log.Printf("Name %s", team.Name)
		log.Println("------------------")
	}

	log.Printf("took %v", time.Now().Sub(now).String())
}
