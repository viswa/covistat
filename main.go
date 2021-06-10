// Covistat displays a short summary of Covid-19 statistics for Kerala
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
)

const TITLE = `
 ██████╗ ██████╗ ██╗   ██╗██╗███████╗████████╗ █████╗ ████████╗
██╔════╝██╔═══██╗██║   ██║██║██╔════╝╚══██╔══╝██╔══██╗╚══██╔══╝
██║     ██║   ██║██║   ██║██║███████╗   ██║   ███████║   ██║   
██║     ██║   ██║╚██╗ ██╔╝██║╚════██║   ██║   ██╔══██║   ██║   
╚██████╗╚██████╔╝ ╚████╔╝ ██║███████║   ██║   ██║  ██║   ██║   
 ╚═════╝ ╚═════╝   ╚═══╝  ╚═╝╚══════╝   ╚═╝   ╚═╝  ╚═╝   ╚═╝
`

// errExit checks err, displays msg and exits the program is err is not nil
func errExit(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("Error: %s\n", msg))
		os.Exit(1)
	}
}

// wg synchronizes fetchResource running in goroutines
var wg sync.WaitGroup

// fetchResource reads HTTP response from source and unmarshals
// it to summary
func fetchResource(source string, summary interface{}) {
	defer wg.Done()

	response, err := http.Get(source)
	errExit(err, "Could not connect to the internet.")
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		errExit(fmt.Errorf(""), "Failed to fetch resources.")
	}

	// Unmarshal body to summary and check for errors
	err = json.NewDecoder(response.Body).Decode(summary)
	errExit(err, "Invalid response.")
}

func main() {
	var summary Summary
	var vaccineSummary VaccineSummary

	wg.Add(2)
	go fetchResource(SummarySource, &summary)
	go fetchResource(VaccinationSummarySource, &vaccineSummary)

	fmt.Print(TITLE)
	wg.Wait()

	// Cases
	fmt.Printf("                         Last Updated: %s IST\n\n", summary.LastUpdated)
	fmt.Printf("               Total confirmed : %-12s [%+d]\n",
		localize(summary.Summary.Confirmed),
		summary.Delta.Confirmed)
	fmt.Printf("               Active Cases    : %-12s [%+d]\n",
		localize(summary.Summary.Active),
		summary.Delta.Active)
	fmt.Printf("               Recovered       : %-12s [%+d]\n",
		localize(summary.Summary.Recovered),
		summary.Delta.Recovered)
	fmt.Printf("               Deaths          : %-12s [%+d]\n\n",
		localize(summary.Summary.Deceased),
		summary.Delta.Deceased)

	// Vaccination stats
	fmt.Println("                      Vaccination Summary")
	fmt.Printf("                   First Dose  : %-12s [%+d]\n",
		localize(vaccineSummary.Summary.FirstDose),
		vaccineSummary.Delta.FirstDose)
	fmt.Printf("                   Second Dose : %-12s [%+d]\n",
		localize(vaccineSummary.Summary.SecondDose),
		vaccineSummary.Delta.SecondDose)
	fmt.Printf("                   Total       : %-12s\n",
		localize(vaccineSummary.Summary.FirstDose+vaccineSummary.Summary.SecondDose))
	fmt.Printf("     %.2f%% Population(3,54,89,000) of Kerala is vaccinated\n\n",
		float64(vaccineSummary.Summary.FirstDose)/KeralaPopulation*100)

	// Quarantine stats
	fmt.Println("                       Quarantine Summary")
	fmt.Printf("                Hospitalized   : %-12s [%+d]\n",
		localize(summary.Summary.HospitalObs),
		summary.Delta.HospitalObs)
	fmt.Printf("                Home Isolation : %-12s [%+d]\n",
		localize(summary.Summary.HomeObs),
		summary.Delta.HomeObs)
	fmt.Printf("                Total          : %-12s\n",
		localize(summary.Summary.TotalObs))
}
