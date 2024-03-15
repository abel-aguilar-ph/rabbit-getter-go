package rabbit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type RabbitBasicAuth struct {
	APIUser string
	APIPwd  string
}

type RabbitConfigure struct {
	APIUrl   string //"Rabbit instance API URL"
	APIVhost string //"Rabbit instance VHOST"
}

// POST /api/queues/vhost/name/get
type GetMessagesBodyRequest struct {
	Count    int    `json:"count"`
	ACKMode  string `json:"ackmode"`
	Encoding string `json:"encoding"`
	//Truncate int    `json:"truncate"` //Opcional
}

func GetMessagesFromQueue(queueName string, numberOfMessages int) {
	var basicAuth RabbitBasicAuth
	basicAuth.APIUser = os.Getenv("RABBITMQ_PRO_USER")
	basicAuth.APIPwd = os.Getenv("RABBITMQ_PRO_PWD")

	var rabbitConfig RabbitConfigure
	rabbitConfig.APIUrl = os.Getenv("RABBITMQ_PRO_URL")
	rabbitConfig.APIVhost = os.Getenv("RABBITMQ_PRO_VHOST")

	urlFinal := fmt.Sprintf("%s/api/queues/%s/%s/get", rabbitConfig.APIUrl, rabbitConfig.APIVhost, queueName)
	//fmt.Println("Calling to -> ",urlFinal)

	requestBody := GetMessagesBodyRequest{
		Count:    numberOfMessages,
		ACKMode:  "ack_requeue_true",
		Encoding: "auto",
		//Truncate: 50000, // Asume un valor o ajusta según necesites
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Error al codificar el cuerpo de la solicitud: %v", err)
	}

	request, err := http.NewRequest(http.MethodPost, urlFinal, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalf("Error al crear la solicitud: %v", err)
	}

	client := &http.Client{}
	request.Header.Add("Content-Type", "application/json")

	request.SetBasicAuth(basicAuth.APIUser, basicAuth.APIPwd)

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalf("Error al realizar la solicitud: %v", err)
	}
	defer resp.Body.Close()

	messageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error al leer la respuesta: %v", err)
	}

	//fmt.Println(messageBytes)
	messagesRaw := string(messageBytes)
	//fmt.Println(string(messages))

	messagesJson, _ := ExtractEntireMessages(messagesRaw)

	extractedIds, _ := ExtractPayhubIds(messagesJson)
	extractedExceptionMsgs, _ := ExtractExceptionMessage(messagesJson)

	ShowExceptionMessages(extractedExceptionMsgs, extractedIds)

	queryRes, _ := KibanaOrQuery(extractedIds)

	fmt.Println()
	fmt.Println("Kibana Query OR:")
	fmt.Print(queryRes)
	fmt.Println()

}
