package utils

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/renaldyhidayatt/redditgoent/ent"
	"github.com/spf13/viper"
)

func Database(c context.Context) (*ent.Client, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", viper.GetString("DB_USER"), viper.GetString("DB_PASSWORD"), viper.GetString("DB_HOST"), viper.GetString("DB_NAME"))

	client, err := ent.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	if err := client.Schema.Create(c); err != nil {
		log.Fatal(err.Error())
	}

	return client, err
}
