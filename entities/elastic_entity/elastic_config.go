package elastic_entity

type Config struct {
	Addresses    []string `json:"addresses"`
	ContentIndex string   `json:"contentIndex"`
	ArticleIndex string   `json:"articleIndex"`
}
