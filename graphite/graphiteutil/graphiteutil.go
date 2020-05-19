package graphiteutil

import (
	"github.com/zapr-oss/go-utils/graphite"
	"log"
)

func InitializeGraphite(config graphite.GraphiteConfig) error {
	graphite.InitDefaultGraphite(config)
	err := graphite.Start()
	if err != nil {
		log.Println("unable to initialize graphite_utils.", err)
		return err
	}
	return nil
}
