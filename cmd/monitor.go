package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zorkian/go-datadog-api"
	"log"
	"os"
	"strconv"
)

var Exitcode bool
var Message bool

func init() {
	RootCmd.PersistentFlags().BoolVarP(&Exitcode, "exitcode", "e", false, "Set output to exit code")
	RootCmd.PersistentFlags().BoolVarP(&Message, "message", "m", false, "Used to output message")
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "monitor",
	Short: "grab monitor info",
	Long:  "Fetch monitor information and status from datadog API",
	Run: func(cmd *cobra.Command, args []string) {

		if len(os.Args) < 4 {
			fmt.Println("Too few arguments provided")
			fmt.Println("./dd-scraper monitor [ID] [-e/-m]")
			os.Exit(1)
		}

		if IntCheck(os.Args[2]) == false {
			fmt.Println("Monitor ID should be an integer")
			os.Exit(1)
		}

		monitorID := os.Args[2]
		id, err := strconv.Atoi(monitorID)
		if err != nil {
			log.Fatalf("fatal: %s\n", err)
		}
		fmt.Println(os.Args[3])

		if Exitcode {
			getMonitorCode(id)
		} else if Message {
			getMonitorMessage(id)
		} else {
			fmt.Println("You must specify an output format using --output")
		}
	},
}

func getMonitorCode(i int) {
	client, err := DatadogClient(os.Getenv("DD_API_KEY"), os.Getenv("DD_APP_KEY"))
	if err != nil {
		log.Fatalf("fatal: %s\n", err)
	}

	m, err := client.Datadog.GetMonitor(*datadog.Int(i))
	if err != nil {
		log.Fatalf("fatal: %s\n", err)
	}

	fmt.Printf("Status for monitor %d returned with '%s'\n", i, *m.OverallState)

	switch *m.OverallState {
	case "OK":
		os.Exit(0)
	case "Warn":
		os.Exit(1)
	case "Alert":
		os.Exit(2)
	case "No Data":
		os.Exit(0)
	}
}

func getMonitorMessage(i int) {
	client, err := DatadogClient(os.Getenv("DD_API_KEY"), os.Getenv("DD_APP_KEY"))
	if err != nil {
		log.Fatalf("fatal: %s\n", err)
	}

	m, err := client.Datadog.GetMonitor(*datadog.Int(i))
	if err != nil {
		log.Fatalf("fatal: %s\n", err)
	}
	fmt.Printf("Status for monitor %d returned with '%s'\n", i, *m.OverallState)
	fmt.Println(*m.Message)
	return
}


