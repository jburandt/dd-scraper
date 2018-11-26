package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"log"
	"os"
	"time"
)

var timeDuration = flag.IntP("time", "t", 10, "length of time in minutes")

func init() {
	RootCmd.AddCommand(eventCmd)
}

var eventCmd = &cobra.Command{
	Use:   "event",
	Short: "grab event info",
	Long:  "Fetch event information and status from datadog API",
	Run: func(cmd *cobra.Command, args []string) {

		now := time.Now()
		count := *timeDuration
		then := now.Add(time.Duration(-count) * time.Minute)

		intNow := int(now.Unix())
		intThen := int(then.Unix())

		getEvents(intThen, intNow, "Normal", "kubernetes", "")
	},
}

func getEvents(start, end int, priority, sources, tags string) {
	client, err := DatadogClient(os.Getenv("DD_API_KEY"), os.Getenv("DD_APP_KEY"))
	if err != nil {
		log.Fatalf("fatal: %s\n", err)
	}

	m, err := client.Datadog.GetEvents(start, end, priority, sources, tags)
	if err != nil {
		log.Fatalf("fatal: %s\n", err)
	}

	fmt.Printf("%v", m)
}
