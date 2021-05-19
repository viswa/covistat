package main

// API for latest summary of COVID-19 stats
const SummarySource = "https://keralastats.coronasafe.live/summary.json"

// Stats represents Covid-19 activity stats
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

// Summary represents present total and changes within a day in Stats
type Summary struct {
	Summary     Stats
	Delta       Stats
	LastUpdated string `json:"last_updated"`
}
