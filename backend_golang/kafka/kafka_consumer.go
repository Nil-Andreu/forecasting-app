package kafka

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func kafka_consumer() {
	want := "Hello world"
	fmt.Println(want)
}