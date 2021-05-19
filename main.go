// Covistat displays a short summary of Covid-19 statistics for Kerala
package main

import (
	"fmt"
	"io"
	"encoding/json"
	"net/http"
	"os"
)

const SummarySource = "https://keralastats.coronasafe.live/summary.json"

const TITLE = `
 ██████╗ ██████╗ ██╗   ██╗██╗███████╗████████╗ █████╗ ████████╗
██╔════╝██╔═══██╗██║   ██║██║██╔════╝╚══██╔══╝██╔══██╗╚══██╔══╝
██║     ██║   ██║██║   ██║██║███████╗   ██║   ███████║   ██║   
██║     ██║   ██║╚██╗ ██╔╝██║╚════██║   ██║   ██╔══██║   ██║   
╚██████╗╚██████╔╝ ╚████╔╝ ██║███████║   ██║   ██║  ██║   ██║   
 ╚═════╝ ╚═════╝   ╚═══╝  ╚═╝╚══════╝   ╚═╝   ╚═╝  ╚═╝   ╚═╝
`

type Stats struct {
	Confirmed     int `json:"confirmed"`
	Recovered     int `json:"recovered"`
	Active        int `json:"active"`
	Deceased      int `json:"deceased"`
	TotalObs      int `json:"total_obs"`
	HospitalObs   int `json:"hospital_obs"`
	HomeObs       int `json:"home_obs"`
	HospitalToday int `json:"hospital_today"`
}

type Summary struct {
	Summary     Stats
	Delta       Stats
	LastUpdated string `json:"last_updated"`
}

func main() {
	fmt.Print(TITLE)
	response, err := http.Get(SummarySource)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not connect to the internet.\n")
		os.Exit(1)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "Error: Failed to fetch resources.\n")
		os.Exit(1)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Internal error.\n")
		os.Exit(1)
	}

	var summary Summary
	json.Unmarshal(body, &summary)
	fmt.Printf("Last Updated: %s IST\n", summary.LastUpdated)
}
