package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/abel-aguilar-ph/rabbit-getter-go/rabbit"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatalln("Usage: ./RabbitGetter Environment(pro1|pro2|k8s) NameOfTheQueue NumberOfMessages")
	} else if os.Args[1] != "pro1" && os.Args[1] != "pro2" && os.Args[1] != "k8s" {
		log.Fatalln("Usage: ./RabbitGetter Environment(pro1|pro2|k8s) NameOfTheQueue NumberOfMessages")
	}
	var basicAuth rabbit.RabbitBasicAuth
	var rabbitConfig rabbit.RabbitConfigure

	environment := os.Args[1]
	queueName := os.Args[2]
	numberOfMessages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatalln("Error with the number of messages")
	}
	fmt.Println()
	fmt.Printf("Environmnet: %s, Name of the queue: %s, Number of messages: %d \n", environment, queueName, numberOfMessages)
	fmt.Println()

	if environment == "pro1" {
		basicAuth.APIUser = os.Getenv("RABBITMQ_PRO1_USER")
		basicAuth.APIPwd = os.Getenv("RABBITMQ_PRO1_PWD")
		rabbitConfig.APIUrl = os.Getenv("RABBITMQ_PRO1_URL")
		rabbitConfig.APIVhost = os.Getenv("RABBITMQ_PRO1_VHOST")

		rabbit.GetMessagesFromQueue(queueName, numberOfMessages, basicAuth, rabbitConfig)

	} else if environment == "pro2" {
		basicAuth.APIUser = os.Getenv("RABBITMQ_PRO2_USER")
		basicAuth.APIPwd = os.Getenv("RABBITMQ_PRO2_PWD")
		rabbitConfig.APIUrl = os.Getenv("RABBITMQ_PRO2_URL")
		rabbitConfig.APIVhost = os.Getenv("RABBITMQ_PRO2_VHOST")

		rabbit.GetMessagesFromQueue(queueName, numberOfMessages, basicAuth, rabbitConfig)
	} else if environment == "k8s" {
		basicAuth.APIUser = os.Getenv("RABBITMQ_K8S_USER")
		basicAuth.APIPwd = os.Getenv("RABBITMQ_K8S_PWD")
		rabbitConfig.APIUrl = os.Getenv("RABBITMQ_K8S_URL")
		rabbitConfig.APIVhost = os.Getenv("RABBITMQ_K8S_VHOST")

		rabbit.GetMessagesFromQueue(queueName, numberOfMessages, basicAuth, rabbitConfig)
	}
	fmt.Println()
}
