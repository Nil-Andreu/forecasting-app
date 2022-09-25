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


func read_env_producer() (string, string, string, string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	KAFKA_API_KEY := os.Getenv("KAFKA_API_KEY")
	KAFKA_SECRET_API_KEY := os.Getenv("KAFKA_SECRET_API_KEY")
	BOOTSTRAP_SERVER := os.Getenv("BOOTSTRAP_SERVER")
	SECURITY_PROTOCOL := os.Getenv("SECURITY_PROTOCOL")
	SASL_MECHANISM := os.Getenv("SASL_MECHANISM")

	return KAFKA_API_KEY, KAFKA_SECRET_API_KEY, BOOTSTRAP_SERVER, SECURITY_PROTOCOL, SASL_MECHANISM
}


func create_producer(
		KAFKA_API_KEY string, 
		KAFKA_SECRET_API_KEY string, 
		BOOTSTRAP_SERVER string, 
		SECURITY_PROTOCOL string, 
		SASL_MECHANISM string) (*kafka.Producer) {
			
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"sasl.username": KAFKA_API_KEY,
		"sasl.password": KAFKA_SECRET_API_KEY,
		"bootstrap.servers": BOOTSTRAP_SERVER,
		"security.protocol": SECURITY_PROTOCOL,
		"sasl.mechanism": SASL_MECHANISM})


	if err != nil {
		log.Fatal("Failed to create the Producer", err)
		os.Exit(1)
	}
	fmt.Println(producer)

	return producer
}


func Kafka_producer() (*kafka.Producer) {
	// We obtain first the environmental variables that we want
	KAFKA_API_KEY, KAFKA_SECRET_API_KEY, BOOTSTRAP_SERVER, SECURITY_PROTOCOL, SASL_MECHANISM := read_env_producer()
	
	consumer := create_producer(KAFKA_API_KEY, KAFKA_SECRET_API_KEY, 
								BOOTSTRAP_SERVER, SECURITY_PROTOCOL, SASL_MECHANISM)

	return consumer
}