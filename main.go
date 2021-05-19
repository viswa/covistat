// Covistat displays a short summary of Covid-19 statistics for Kerala
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

func main() {
	fmt.Print(TITLE)
	response, err := http.Get(SummarySource)
	errExit(err, "Could not connect to the internet.")
	defer response.Body.Close()

	if response.StatusCode != 200 {
		errExit(fmt.Errorf(""), "Failed to fetch resources.")
	}

	body, err := io.ReadAll(response.Body)
	errExit(err, "Internal error.")

	var summary Summary
	json.Unmarshal(body, &summary)
	fmt.Printf("Last Updated: %s IST\n", summary.LastUpdated)
}
