package kafka_backend

import (
	"fmt"
	"os"
	"log"
	_ "os/signal"
	_ "syscall"

	"github.com/joho/godotenv"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Kafka_consumer() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	KAFKA_API_KEY := os.Getenv("KAFKA_API_KEY")
	KAFKA_SECRET_API_KEY := os.Getenv("KAFKA_SECRET_API_KEY")
	KAFKA_GROUP_ID := os.Getenv("KAFKA_GROUP_ID")
	BOOTSTRAP_SERVER := os.Getenv("BOOTSTRAP_SERVER")
	SECURITY_PROTOCOL := os.Getenv("SECURITY_PROTOCOL")
	SASL_MECHANISM := os.Getenv("SASL_MECHANISM")

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"sasl.username": KAFKA_API_KEY,
		"sasl.password": KAFKA_SECRET_API_KEY,
		"bootstrap.servers": BOOTSTRAP_SERVER,
		"group.id": KAFKA_GROUP_ID,
		"security.protocol": SECURITY_PROTOCOL,
		"sasl.mechanism": SASL_MECHANISM})

	fmt.Println(consumer)
}