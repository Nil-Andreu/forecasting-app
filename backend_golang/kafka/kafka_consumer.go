package main

import (
	"fmt"
	"os"
	"log"
	_ "os/signal"
	_ "syscall"

	"github.com/joho/godotenv"
	_ "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	KAFKA_API_KEY := os.Getenv("KAFKA_API_KEY")
	KAFKA_SECRET_API_KEY := os.Getenv("KAFKA_SECRET_API_KEY")
	KAFKA_GROUP_ID := os.Getenv("KAFKA_GROUP_ID")
	KAFKA_CONSUMER_ID := os.Getenv("KAFKA_CONSUMER_ID")
	KAFKA_PRODUCER_ID := os.Getenv("KAFKA_PRODUCER_ID")
	BOOTSTRAP_SERVER := os.Getenv("BOOTSTRAP_SERVER")
	SECURITY_PROTOCOL := os.Getenv("SECURITY_PROTOCOL")
	SASL_MECHANISM := os.Getenv("SASL_MECHANISM")

	fmt.Println(KAFKA_API_KEY)
}