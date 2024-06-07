package rabbit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// POST /api/queues/vhost/name/get

func GetMessagesFromQueue(queueName string, numberOfMessages int, rabbitConfig *RabbitConfigure) (messagesRaw string) {

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

	request.SetBasicAuth(rabbitConfig.APIUser, rabbitConfig.APIPwd)

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalf("Error al realizar la solicitud:  %s. %v", rabbitConfig.APIVhost, err)
	}
	defer resp.Body.Close()

	messageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error al leer la respuesta: %v", err)
	}

	//fmt.Println(messageBytes)
	messagesRaw = string(messageBytes)
	return messagesRaw
}

func PrintFullMessages(messagesRaw string, queueName string) {
	messagesJson, _ := ExtractEvidencesFullMessages(messagesRaw)
	for _, msg := range messagesJson {
		fmt.Println(msg)
	}
}

func PrintIDsKibanaQueryWithErrors(messagesRaw string, queueName string) {
	messagesJson, _ := ExtractParcialMessages(messagesRaw)
	extractedIds, _ := ExtractPayhubIds(messagesJson, queueName)
	extractedExceptionMsgs, _ := ExtractExceptionMessage(messagesJson)
	ShowExceptionMessages(extractedExceptionMsgs, extractedIds)
}

func PrintIDsKibanaQuery(messagesRaw string, queueName string) {
	messagesJson, _ := ExtractParcialMessages(messagesRaw)
	extractedIds, _ := ExtractPayhubIds(messagesJson, queueName)
	ShowKibanaOrQuery(extractedIds)
}
