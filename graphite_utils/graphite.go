package graphite_utils

import (
	"log"
	graphite "bitbucket.org/zapr/graphite_go"
)

func InitializeGraphite(config graphite.GraphiteConfig) {
	graphite.InitDefaultGraphite(config)
	err := graphite.Start()
	if err != nil {
		log.Fatal("Unable to initialize graphite_utils.")
	}
}