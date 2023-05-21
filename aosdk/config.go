package aosdk

type Config struct {
	Endpoint string
	Version  string
}

var config Config

func InitConfig(endpoint, version string) {
	config = Config{
		Endpoint: endpoint, // such as "http://localhost:8080"
		Version:  version,  // v2, maybe v3 in the future
	}
}
