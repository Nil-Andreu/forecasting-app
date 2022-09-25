package main

import (
	_ "fmt"

	"./kafka_backend"  // because we have modules off
)

func main() {
	kafka_backend.Kafka_consumer()

	// fmt.Println(consumer)
}