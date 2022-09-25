package main

import (
	_ "fmt"

	"./kafka_backend"
)

func main() {
	kafka_backend.Kafka_consumer()
	// fmt.Println(want)
}