package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "extract data from NOAA buoy data",
	Run:   extractData,
}

func extractData(cmd *cobra.Command, args []string) {
	time.Sleep(3 * time.Second)
	fmt.Printf("add http stuff here")
}
