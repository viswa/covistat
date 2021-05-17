// Covistat displays a short summary of Covid-19 statistics for Kerala
package main

import (
	"fmt"
)

const SUMMARY_SOURCE = "https://raw.githubusercontent.com/coronasafe/kerala-stats/master/summary.json"

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
}
