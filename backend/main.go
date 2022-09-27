package main

import (
	_ "fmt"

	"./kafka"  // because we have modules off
)

func main() {
	kafka.Kafka_consumer()
	kafka.Kafka_producer()

	// fmt.Println(consumer)
}