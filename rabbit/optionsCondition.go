package rabbit

//FullMessages|KibanaQuery|KibanaQueryErrors)

func PrintOptionCondition(option string, messagesRaw string, queueName string) {
	if option == "FullMessages" {
		PrintFullMessages(messagesRaw, queueName)
	} else if option == "KibanaQuery" {
		PrintIDsKibanaQuery(messagesRaw, queueName)
	} else if option == "KibanaQueryErrors" {
		PrintIDsKibanaQueryWithErrors(messagesRaw, queueName)
	}
}
