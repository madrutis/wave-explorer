package cmd

const NOAA_BASE_URL = "https://www.ndbc.noaa.gov/data/realtime2/"

var DATA_TYPES = [...]string{
	"txt",       // standard meteorological data
	"drift",     // meteorological data from drifting buoys
	"cwind",     // continuous winds data
	"spec",      // spectral wave summaries
	"data_spec", // raw spectral wave data
	"swdir",     // spectral wave data (alpha1)
	"swdir2",    // spectral wave data (alpha2)
	"swr1",      // spectral wave data (r1)
	"swr2",      // spectral wave data (r2)
	"adcp",      // Acoustic Doppler Current Profiler
	"ocean",     // oceanographic data
	"tide",      // tide data
}

// TODO: add list of stations (will probably have to be fetched via script)
// var STATIONS = [...]string{}

var DATA_FOLDER = "data/"
