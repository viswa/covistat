// Covistat displays a short summary of Covid-19 statistics for Kerala
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/logrusorgru/aurora/v3"
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

	fmt.Print(aurora.Index(189, TITLE))
	wg.Wait()
	vaccineSummary.CalcSecondDose()

	// Cases
	fmt.Printf("                         %v: %s IST\n\n",
		aurora.Index(69, "Last Updated"),
		summary.LastUpdated)
	fmt.Printf("               %v : %-12s [%+d]\n",
		aurora.Index(69, "Total confirmed"),
		localize(summary.Summary.Confirmed),
		summary.Delta.Confirmed)
	fmt.Printf("               %v    : %-12s [%+d]\n",
		aurora.Index(69, "Active Cases"),
		localize(summary.Summary.Active),
		summary.Delta.Active)
	fmt.Printf("               %v       : %-12s [%+d]\n",
		aurora.Index(69, "Recovered"),
		localize(summary.Summary.Recovered),
		summary.Delta.Recovered)
	fmt.Printf("               %v          : %-12s [%+d]\n\n",
		aurora.Index(69, "Deaths"),
		localize(summary.Summary.Deceased),
		summary.Delta.Deceased)

	// Vaccination stats
	fmt.Println("                      Vaccination Summary")
	fmt.Printf("                   %v  : %-12s [+%s]\n",
		aurora.Index(69, "First Dose"),
		localize(vaccineSummary.Summary.FirstDose),
		localize(vaccineSummary.Delta.FirstDose))
	fmt.Printf("                   %v : %-12s [+%s]\n",
		aurora.Index(69, "Second Dose"),
		localize(vaccineSummary.Summary.SecondDose),
		localize(vaccineSummary.Delta.SecondDose))
	fmt.Printf("                   %v      : %-12s\n",
		aurora.Index(69, "Total "),
		localize(vaccineSummary.Summary.TotalDose))
	fmt.Printf("     %.2f%% Population(3,54,89,000) of Kerala is vaccinated\n\n",
		vaccineSummary.VaccinatedPercent())

	// Quarantine stats
	fmt.Println("                       Quarantine Summary")
	fmt.Printf("                %v   : %-12s [%+d]\n",
		aurora.Index(69, "Hospitalized"),
		localize(summary.Summary.HospitalObs),
		summary.Delta.HospitalObs)
	fmt.Printf("                %v : %-12s [%+d]\n",
		aurora.Index(69, "Home Isolation"),
		localize(summary.Summary.HomeObs),
		summary.Delta.HomeObs)
	fmt.Printf("                %v          : %-12s\n",
		aurora.Index(69, "Total"),
		localize(summary.Summary.TotalObs))
}
