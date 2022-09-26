package kafka

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


func Kafka_consumer() {
	// We obtain first the environmental variables that we want
	KAFKA_API_KEY, KAFKA_SECRET_API_KEY, KAFKA_GROUP_ID, BOOTSTRAP_SERVER, SECURITY_PROTOCOL, SASL_MECHANISM := read_env_consumer()
	
	consumer := create_consumer(KAFKA_API_KEY, KAFKA_SECRET_API_KEY, KAFKA_GROUP_ID, 
								BOOTSTRAP_SERVER, SECURITY_PROTOCOL, SASL_MECHANISM)


	// And now we will consume the services
	err := consumer.SubscribeTopics([]string{"financials"}, nil)

	if err != nil {
		log.Fatal("Failed to suscribe to Topic")
	}

	// And now we run the consumer
	run := true

	for run {
		// Poll the new messages
		event := consumer.Poll(1)

		// No new events
		if event == nil {
			continue
		}

		// And now depending on the event type, we would do a different thing
		switch e := event.(type) {
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))

				if e.Headers != nil {
					fmt.Printf("%% Headers: %v\n", e.Headers)
				}

				_, err := consumer.StoreMessage(e)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%% Error storing offset after message %s:\n",
						e.TopicPartition)
				}

			case kafka.Error:
				// Print information about this error
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)

				// If all brokers are down, we stop consuming
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				fmt.Printf("Ignored %v\n", e)
		}
	}

	fmt.Printf("Closing consumer")
	consumer.Close()
}