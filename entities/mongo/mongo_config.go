package mongo

// Mongo contains details of MongoDB
type Mongo struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Database   string `json:"database"`
	Collection string `json:"collection"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
}