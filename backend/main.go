package main

import (
	_ "fmt"

	"./kafka_backend"  // because we have modules off
)

func main() {
	kafka_backend.Kafka_consumer()
	kafka_backend.Kafka_producer()

	// fmt.Println(consumer)
}