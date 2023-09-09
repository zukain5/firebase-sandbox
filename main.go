package main

import (
	"context"
	"log"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"
)

func addDocDataTypes(ctx context.Context, client *firestore.Client) error {
	doc := make(map[string]interface{})
	doc["stringExample"] = "Hello world!"
	doc["booleanExample"] = true
	doc["numberExample"] = 3.14159265
	doc["dateExample"] = time.Now()
	doc["arrayExample"] = []interface{}{5, true, "hello"}
	doc["nullExample"] = nil
	doc["objectExample"] = map[string]interface{}{
		"a": 5,
		"b": true,
	}

	_, err := client.Collection("data").Doc("one").Set(ctx, doc)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	return err
}

// City represents a city.
type City struct {
	Name       string   `firestore:"name,omitempty"`
	State      string   `firestore:"state,omitempty"`
	Country    string   `firestore:"country,omitempty"`
	Capital    bool     `firestore:"capital,omitempty"`
	Population int64    `firestore:"population,omitempty"`
	Regions    []string `firestore:"regions,omitempty"`
}

func addDocAsEntity(ctx context.Context, client *firestore.Client) error {
	city := City{
		Name:    "Los Angeles",
		Country: "USA",
	}
	_, err := client.Collection("cities").Doc("LA").Set(ctx, city)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	return err
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	conf := &firebase.Config{ProjectID: os.Getenv("PROJECT_ID")}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	addDocAsEntity(ctx, client)
}
