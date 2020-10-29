package config

const (
	defaultHTTPS         = false
	defaultDebug         = true
	defaultBaseURL       = "http://localhost"
	defaultRunMigrations = false
	defaultDatabaseDSN   = "user=librefitness password=librefitness dbname=librefitness sslmode=disable"
)

// Options holds apps configuration
type Options struct {
	Debug          bool
	BaseURL        string
	RunMigrations  bool
	DatabaseDSN    string
	DatabaseDriver string
}

func newOptions() *Options {
	return &Options{
		Debug:         defaultDebug,
		BaseURL:       defaultBaseURL,
		RunMigrations: defaultRunMigrations,
		DatabaseDSN:   defaultDatabaseDSN,
	}
}
