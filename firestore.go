package favdb

import (
	"context"
	"log"

	firestore "cloud.google.com/go/firestore"
)

var client *firestore.Client

func NewDB() *firestore.Client {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "PROJECT_ID")
	if err != nil {
		log.Fatalf("Error initializing Cloud Firestore client: %v", err)
	}
	return client
}
