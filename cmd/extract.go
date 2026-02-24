package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

type ExtractConfig struct {
	Station  string
	DataType string
}

var cfg ExtractConfig

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "extract data from NOAA buoy data",
	Run:   extractData,
}

// reference this:
// https://gobyexample.com/writing-files
func extractData(cmd *cobra.Command, args []string) {
	requestUrl := fmt.Sprintf("%s%s.%s", NOAA_BASE_URL, cfg.Station, cfg.DataType)
	fmt.Printf("requesting data from: %s", requestUrl)
	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	// fetch data
	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	responseBytes, err := io.ReadAll(res.Body)
	check(err)

	// write data to data folder with format <station>.<type>.<currentTimestamp>
	relativePath := fmt.Sprintf("%s%s.%s.%s", DATA_FOLDER, cfg.Station, cfg.DataType, time.Now().Format("2006-01-02T15-04-05"))
	cwd, err := os.Getwd()
	absolutePath := filepath.Join(cwd, relativePath)

	err = os.WriteFile(absolutePath, responseBytes, 0644)
	check(err)
}

func init() {

	// local flags
	extractCmd.Flags().StringVarP(&cfg.Station, "station", "s", "46025", "NOAA station to fetch data from")
	extractCmd.Flags().StringVarP(&cfg.DataType, "dtype", "d", "txt", "NOAA data file type to fetch")
	// extractCmd.Flags().StringSliceVarP(&cfg.Stations, "stations", "s", []string{"txt"}, "NOAA stations to fetch data from")
	// extractCmd.Flags().StringSliceVarP(&cfg.DataTypes, "dtypes", "d", []string{"46025"}, "NOAA data file types to fetch")
}
