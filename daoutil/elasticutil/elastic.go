package elasticutil

import (
	"bitbucket.org/zapr/go-utils/entities/elastic_entity"
	graphite "bitbucket.org/zapr/graphite_go"
	"github.com/olivere/elastic"
	"log"
)

/*
	Returns an Elastic Entity struct with a client connection open to the ElasticSearch.
	On error, the error is returned.
*/
func Connect(config elastic_entity.Config) (*elastic_entity.Entity, error) {
	client, err := elastic.NewClient(elastic.SetURL(config.Addresses...))
	if err != nil {
		graphite.GetCounter("ElasticSearchConnectionError").Inc()
		log.Println("Error connecting to ElasticSearch, Error: ", err)
		return nil, err
	}

	return &elastic_entity.Entity{Client: client}, nil
}