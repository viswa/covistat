// Covistat displays a short summary of Covid-19 statistics for Kerala
package main

import (
	"fmt"
	"io"
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

	fmt.Printf("Last Updated: %s IST\n", summary.LastUpdated)
}
