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

type FullMessage struct {
	//Exchange   string         `json:"exchange"`
	Payload    string         `json:"payload"` // JSON stringificado
	Properties FullProperties `json:"properties"`
	RoutingKey string         `json:"routing_key"`
}
type FullProperties struct {
	//Timestamp    int64       `json:"timestamp"`
	//MessageID string `json:"message_id"`
	//Priority  int    `json:"priority"`
	//DeliveryMode int         `json:"delivery_mode"`
	Headers FullHeaders `json:"headers"`
	//ContentType string      `json:"content_type"`
}
type FullHeaders struct {
	AccountingRoutingKey string `json:"accountingRoutingKey"`
	//B3                   string `json:"b3"`
	//ClearingHouse    string `json:"clearingHouse"`
	//CycleDate        string `json:"cycleDate"`
	//CycleId          string `json:"cycleId"`
	//DtdTraceTagInfo  string `json:"dtdTraceTagInfo"`
	Flow string `json:"flow"`
	//Forced           string `json:"forced"`
	//OrchestratorStep string `json:"orchestratorStep"`
	PayhubRoutingKey string `json:"payhubRoutingKey"`
	Scheme           string `json:"scheme"`
	Source           string `json:"source"`
	//SourceType       string `json:"source-type"`
	Target string `json:"target"`
	//TargetProtocol   string `json:"target-protocol"`
	//ExceptionMessage     string `json:"x-exception-message"`
	//ExceptionStacktrace  string `json:"x-exception-stacktrace"`
	OriginalExchange   string `json:"x-original-exchange"`
	OriginalRoutingKey string `json:"x-original-routingKey"`
}

// PUBLISH STRUCT
type FullMessageToSend struct {
	Payload    string                  `json:"payload"` // JSON stringificado
	Properties FullPropertiesToPublish `json:"properties"`
	RoutingKey string                  `json:"routing_key"`
}
type FullPropertiesToPublish struct {
	Headers FullHeadersToPublish `json:"headers"`
}
type FullHeadersToPublish struct {
	AccountingRoutingKey string `json:"accountingRoutingKey"`
	Flow                 string `json:"flow"`
	PayhubRoutingKey     string `json:"payhubRoutingKey"`
	Scheme               string `json:"scheme"`
	Source               string `json:"source"`
	Target               string `json:"target"`
	OriginalExchange     string `json:"x-original-exchange"`
	OriginalRoutingKey   string `json:"x-original-routingKey"`
}
