package elastic

import (
	"fmt"
	"fx-golang-server/config"

	"github.com/rs/zerolog/log"

	"github.com/elastic/go-elasticsearch/v7"
)

func ESProvider(cfg *config.Config) *elasticsearch.Client {
	fmt.Println("fuck ESProvider", cfg.Elastic)
	client, err := NewES(cfg.Elastic)
	if err != nil {
		log.Panic().Err(err).Msg("connect to elastic error")
	}
	return client
}

func NewES(cfg config.ElasticConfig) (*elasticsearch.Client, error) {
	escfg := elasticsearch.Config{
		Addresses: []string{cfg.Addresses},
		Username:  cfg.Username,
		Password:  cfg.Password,
	}

	fmt.Println("fuck", cfg)

	es, err := elasticsearch.NewClient(escfg)
	if err != nil {
		log.Error().Err(err).Msg("elastic search connection err")
		return nil, err
	}
	// Ping Elasticsearch
	res, err := es.Ping()
	if err != nil {
		log.Error().Err(err).Msg("Error pinging Elasticsearch")
		return nil, err
	}
	// Check the response status
	if res.IsError() {
		log.Error().Err(err).Msg(fmt.Sprintf("Elasticsearch ping error: %s", res.String()))
	}
	return es, nil
}
