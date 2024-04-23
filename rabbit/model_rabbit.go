package rabbit

type RabbitConfigure struct {
	APIUrl   string //"Rabbit instance API URL"
	APIVhost string //"Rabbit instance VHOST"
	APIUser  string
	APIPwd   string
}

type GetMessagesBodyRequest struct {
	Count    int    `json:"count"`
	ACKMode  string `json:"ackmode"`
	Encoding string `json:"encoding"`
	//Truncate int    `json:"truncate"` //Opcional
}
