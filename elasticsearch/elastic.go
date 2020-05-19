package elasticsearch

import (
	"log"
	"github.com/olivere/elastic"
	"github.com/zapr-oss/go-utils/elasticsearch/config"
)

// Returns an Elastic Entity struct with a client connection open to the ElasticSearch. On error, the error is returned.
func Connect(config elasticconfig.Config) (*Entity, error) {
	client, err := elastic.NewClient(elastic.SetURL(config.Addresses...))
	if err != nil {
		log.Println("Error connecting to ElasticSearch, Error: ", err)
		return nil, err
	}

	return &Entity{Client: client}, nil
}

type Entity struct {
	Client *elastic.Client
}
