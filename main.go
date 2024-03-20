package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/abel-aguilar-ph/rabbit-getter-go/rabbit"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Usage: [program] nameOfTheQueue numberOfMessages")
	}

	queueName := os.Args[1]
	numberOfMessages, _ := strconv.Atoi(os.Args[2])
	fmt.Println()
	fmt.Printf("Name of the queue: %s, Number of messages: %d \n", queueName, numberOfMessages)
	fmt.Println()
	rabbit.GetMessagesFromQueue(queueName, numberOfMessages)
	fmt.Println()
}
