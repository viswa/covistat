// Covistat displays a short summary of Covid-19 statistics for Kerala
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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

// Unmarshaller interface represents types that can unmarshal JSON data
type Unmarshaller interface {
	Unmarshal([]byte) error
}

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
func fetchResource(source string, summary Unmarshaller) {
	defer wg.Done()

	response, err := http.Get(source)
	errExit(err, "Could not connect to the internet.")
	defer response.Body.Close()

	if response.StatusCode != 200 {
		errExit(fmt.Errorf(""), "Failed to fetch resources.")
	}

	body, err := io.ReadAll(response.Body)
	errExit(err, "Internal error.")

	// Unmarshal body to summary and check for errors
	errExit(summary.Unmarshal(body), "Invalid response.")
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

// localize inserts commas to num based on Indian locale
func localize(num int) string {
	digits := []rune(fmt.Sprint(num))
	var builder strings.Builder
	var written int // no. of characters written to builder
	sep := 3        // no. of places between comma placement

	// digit characters are written to builder in reverse order
	for i := len(digits) - 1; i >= 0; i-- {
		if written == sep {
			builder.WriteString(",")
			// reset no. of written characters after each comma placed
			written = 0
			sep = 2
		}
		builder.WriteRune(digits[i])
		written++
	}

	// builder string is to be further reversed
	reversed := builder.String()
	digits = []rune(reversed)
	n := len(digits)
	for i := 0; i < n/2; i++ {
		digits[i], digits[n-1-i] = digits[n-1-i], digits[i]
	}
	return string(digits)
}
