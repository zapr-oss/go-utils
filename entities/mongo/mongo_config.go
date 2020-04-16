package mongo

// Config contains details of MongoDB
type Config struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Database   string `json:"database"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
}
