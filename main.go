package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/abel-aguilar-ph/rabbit-getter-go/rabbit"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatalln("Usage: ./RabbitGetter Environment(pro1|pro2|k8s) NameOfTheQueue NumberOfMessages Option(FullMessages|KibanaQuery|KibanaQueryErrors)")
	} else if os.Args[1] != "pro1" && os.Args[1] != "pro2" && os.Args[1] != "k8s" && os.Args[1] != "pre1" {
		log.Fatalln("Usage: ./RabbitGetter Environment(pro1|pro2|k8s) NameOfTheQueue NumberOfMessages Option(FullMessages|KibanaQuery|KibanaQueryErrors)")
	} else if os.Args[4] != "FullMessages" && os.Args[4] != "KibanaQuery" && os.Args[4] != "KibanaQueryErrors" {
		log.Fatalln("Usage: ./RabbitGetter Environment(pro1|pro2|k8s) NameOfTheQueue NumberOfMessages Option(FullMessages|KibanaQuery|KibanaQueryErrors)")
	}

	environment := os.Args[1]
	queueName := os.Args[2]
	numberOfMessages, err := strconv.Atoi(os.Args[3])
	option := os.Args[4]
	if err != nil {
		log.Fatalln("Error with the number of messages")
	}
	fmt.Println()
	fmt.Printf("Environmnet: %s, Name of the queue: %s, Number of messages: %d, Option: %s \n", environment, queueName, numberOfMessages, option)
	fmt.Println()

	if environment == "pro1" {
		var rabbitConfig rabbit.RabbitConfigure

		rabbitConfig.APIUser = os.Getenv("RABBITMQ_PRO1_USER")
		rabbitConfig.APIPwd = os.Getenv("RABBITMQ_PRO1_PWD")
		rabbitConfig.APIUrl = os.Getenv("RABBITMQ_PRO1_URL")
		rabbitConfig.APIVhost = os.Getenv("RABBITMQ_PRO1_VHOST")

		messagesRaw := rabbit.GetMessagesFromQueue(queueName, numberOfMessages, &rabbitConfig)
		rabbit.PrintOptionCondition(option, messagesRaw, queueName)
	} else if environment == "pro2" {
		var rabbitConfig rabbit.RabbitConfigure

		rabbitConfig.APIUser = os.Getenv("RABBITMQ_PRO2_USER")
		rabbitConfig.APIPwd = os.Getenv("RABBITMQ_PRO2_PWD")
		rabbitConfig.APIUrl = os.Getenv("RABBITMQ_PRO2_URL")
		rabbitConfig.APIVhost = os.Getenv("RABBITMQ_PRO2_VHOST")

		messagesRaw := rabbit.GetMessagesFromQueue(queueName, numberOfMessages, &rabbitConfig)
		rabbit.PrintOptionCondition(option, messagesRaw, queueName)

	} else if environment == "k8s" {
		var rabbitConfig rabbit.RabbitConfigure

		rabbitConfig.APIUser = os.Getenv("RABBITMQ_K8S_USER")
		rabbitConfig.APIPwd = os.Getenv("RABBITMQ_K8S_PWD")
		rabbitConfig.APIUrl = os.Getenv("RABBITMQ_K8S_URL")
		rabbitConfig.APIVhost = os.Getenv("RABBITMQ_K8S_VHOST")

		messagesRaw := rabbit.GetMessagesFromQueue(queueName, numberOfMessages, &rabbitConfig)
		rabbit.PrintOptionCondition(option, messagesRaw, queueName)

	} else if environment == "pre1" {
		var rabbitConfig rabbit.RabbitConfigure

		rabbitConfig.APIUser = os.Getenv("RABBITMQ_PRE1_USER")
		rabbitConfig.APIPwd = os.Getenv("RABBITMQ_PRE1_PWD")
		rabbitConfig.APIUrl = os.Getenv("RABBITMQ_PRE1_URL")
		rabbitConfig.APIVhost = os.Getenv("RABBITMQ_PRE1_VHOST")

		messagesRaw := rabbit.GetMessagesFromQueue(queueName, numberOfMessages, &rabbitConfig)
		rabbit.PrintOptionCondition(option, messagesRaw, queueName)
	}

}
