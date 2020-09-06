package nhiapi // package NHIApi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Roster struct {
	Person struct {
		ID       int    `json:"id"`
		FullName string `json:"fullName"`
		Link     string `json:"link"`
	} `json:"person"`
	JerseyNumber string `json:"jerseyNumber"`
	Position     struct {
		Code         string `json:"code"`
		Name         string `json:"name"`
		Type         string `json:"type"`
		Abbreviation string `json:"abbreviation"`
	} `json:"position"`
}

type nhiRosterResponse struct {
	Roster []Roster `json:"roster"`
}

func GetRoster(teamID int) ([]Roster, error) {
	res, err := http.Get(fmt.Sprintf("%s/teams/%d/roster", baseURL, teamID))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var response nhiRosterResponse
	err = json.NewDecoder(res.Body).Decode(&response)

	return response.Roster, err
}
