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


func read_env_consumer() (string, string, string, string, string, string) {
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

	return KAFKA_API_KEY, KAFKA_SECRET_API_KEY, KAFKA_GROUP_ID, BOOTSTRAP_SERVER, SECURITY_PROTOCOL, SASL_MECHANISM
}


func create_consumer(
		KAFKA_API_KEY string, 
		KAFKA_SECRET_API_KEY string, 
		KAFKA_GROUP_ID string, 
		BOOTSTRAP_SERVER string, 
		SECURITY_PROTOCOL string, 
		SASL_MECHANISM string) (*kafka.Consumer) {
			
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"sasl.username": KAFKA_API_KEY,
		"sasl.password": KAFKA_SECRET_API_KEY,
		"bootstrap.servers": BOOTSTRAP_SERVER,
		"group.id": KAFKA_GROUP_ID,
		"security.protocol": SECURITY_PROTOCOL,
		"sasl.mechanism": SASL_MECHANISM})


	if err != nil {
		log.Fatal("Failed to create the Consumer", err)
		os.Exit(1)
	}
	fmt.Println(consumer)

	return consumer
}


func Kafka_consumer() (*kafka.Consumer) {
	// We obtain first the environmental variables that we want
	KAFKA_API_KEY, KAFKA_SECRET_API_KEY, KAFKA_GROUP_ID, BOOTSTRAP_SERVER, SECURITY_PROTOCOL, SASL_MECHANISM := read_env_consumer()
	
	consumer := create_consumer(KAFKA_API_KEY, KAFKA_SECRET_API_KEY, KAFKA_GROUP_ID, 
								BOOTSTRAP_SERVER, SECURITY_PROTOCOL, SASL_MECHANISM)

	return consumer
}