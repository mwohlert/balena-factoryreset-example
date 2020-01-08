package constants

const (
	// BalenaAppID : App id of the openBalena application
	BalenaAppID string = "0"
	// BalenaAppIDEnv : Environment variable containing the app id
	BalenaAppIDEnv string = "BALENA_APP_ID"
	// BalenaSupervisorURIEnv : Environment variable containing the balena supervisor address
	BalenaSupervisorURIEnv string = "BALENA_SUPERVISOR_ADDRESS"
	// BalenaSupervisorAPIKeyEnv : Environment variable containing the balena supervisor api key
	BalenaSupervisorAPIKeyEnv string = "BALENA_SUPERVISOR_API_KEY"

	// Port : Port on which the webserver runs on
	Port string = "8080"
	// Host : Host on which the webserver listens
	Host string = "0.0.0.0"
)
