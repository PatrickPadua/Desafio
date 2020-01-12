package main

import (
	"os"
	"log"
	"time"
	. "app/infra/aws"
	. "app/domain/suggestion"
	"github.com/joho/godotenv"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var mongo = MongoDB{}

func init() {
	godotenv.Load()
	mongo.Server = os.Getenv("MONGODB_HOST")
	mongo.Database = os.Getenv("MONGODB_DATABASE")
	mongo.Connect()
}

func main() {
	
	ConsulURL := os.Getenv("CONSUL_HOST")
	ConsulEndpoint := os.Getenv("CONSUL_ENDPOINT")
	COLLECTION := os.Getenv("MONGODB_COLLECTION_NAME")

	sess := session.New(&aws.Config{
		MaxRetries:  aws.Int(5),
	})

	svc := sqs.New(sess)

	log.Printf("Worker is running...")
	for _ = range time.Tick(1 * time.Second) {
		go SQSReceiveAndDrop(svc, ConsulURL, ConsulEndpoint, COLLECTION)
	}
}