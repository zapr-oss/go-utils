package mysqlconfig

type Config struct {
	User       string `json:"user"`
	Password   string `json:"password"`
	Database   string `json:"database"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	RetryCount int    `json:"retryCount"`
}
