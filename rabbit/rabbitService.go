package rabbit

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Message struct {
	Properties Properties `json:"properties"`
	Payload string `json:"payload"` //JSON stringificado 
}

type Properties struct {
	Headers struct{
		ExceptionMessage string `json:"x-exception-message"`
	} `json:"headers"`
}

type Payload struct {
	ExecutionId struct {
		PaymentHubId string `json:"paymentHubId"`
	} `json:"executionId"`
}

func ExtractEntireMessages(messagesRaw string) ([]Message, error) {
	var messages []Message
	if err := json.Unmarshal([]byte(messagesRaw), &messages); err != nil {
		fmt.Printf("Error decodificando el JSON principal: %v\n", err)
		return nil, err
	}
	return messages, nil
}

func ExtractExceptionMessage(messages []Message) ([]string, error) {
	var exceptionMsgs []string
	for _, msg := range messages {
	
		if msg.Properties.Headers.ExceptionMessage != "" {
			exceptionMsgs = append(exceptionMsgs, msg.Properties.Headers.ExceptionMessage)
		}

		if len(exceptionMsgs) == 0 {
			return nil, fmt.Errorf("no se encontraron exceptionMessages")
		}
	}

	return exceptionMsgs, nil
}

func ShowExceptionMessages(excepMsgs []string, ids []string) {
	var res = make(map[string]([]string))
	if(len(excepMsgs)>0){
		for index, excepMsg := range excepMsgs { 
			value, ok := res[excepMsg] 
			if ok {                    
				res[excepMsg] = append(value, ids[index])
			} else {
				res[excepMsg] = []string{ids[index]}
			}
		}
		fmt.Println("X-Exception-Messages:\n")
		for key, value := range res {
			var idsSameError []string
			fmt.Println(key)
			fmt.Println("\t Payments with the error above (",len(value),"):")
			idsSameError = append(idsSameError,value...)
			/*for _, id := range value {
				idsSameError = append(idsSameError, id)
			}*/
			idsKibanaError,_ := KibanaOrQuery(idsSameError)
			fmt.Println(idsKibanaError)
			fmt.Println()
		}
	}else{
		fmt.Println()
		fmt.Println("x-excepcion-message field not found")
	}
}

func ExtractPayhubIds(messagesJson []Message) ([]string, error) {
	//fmt.Println(messagesPayload)
	var ids []string
	for _, msg := range messagesJson {
		var payload Payload
		if err := json.Unmarshal([]byte(msg.Payload), &payload); err != nil {
			fmt.Printf("Error decodificando el Payload: %v\n", err)
			return nil, err
		}
		//fmt.Println("JSON arreglado sin espacios:", msg.Payload)
		//fmt.Printf("Payload deserializado: %+v\n", payload)
		if payload.ExecutionId.PaymentHubId != "" {
			ids = append(ids, payload.ExecutionId.PaymentHubId)
		}
	}

	if len(ids) == 0 {
		return nil, fmt.Errorf("no se encontraron payhubIds")
	}

	return ids, nil
}

func KibanaOrQuery(ids []string) (string, error) {
	var formattedIds []string
	for _, id := range ids {
		formattedId := fmt.Sprintf("\"%s\"", id)
		formattedIds = append(formattedIds, formattedId)
	}

	query := strings.Join(formattedIds, " OR ")
	return query, nil
}
