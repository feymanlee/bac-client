package main

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	bac "github.com/feymanlee/bac-client"
)

func main() {
	client, err := bac.NewClient(
		os.Getenv("BAC_APP_KEY"),
		os.Getenv("BAC_APP_SECRET"),
		os.Getenv("BAC_DES_KEY"),
		bac.WithTimeout(15*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	instances, err := client.ListInstances(ctx, &bac.ListInstancesRequest{
		PageRequest: bac.PageRequest{Page: 1},
		Rows:        10,
	})
	if err != nil {
		var apiErr *bac.APIError
		if errors.As(err, &apiErr) {
			log.Fatalf("bac api error: code=%d message=%s ts=%s", apiErr.Code, apiErr.Message, apiErr.Timestamp)
		}
		log.Fatal(err)
	}

	log.Printf("instances: %+v", instances)
}
