package message

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/postech-soat2-grupo16/producao-api/entities"
	"os"
)

type Gateway struct {
	queueURL string
	queue    *sqs.SQS
}

func NewGateway(queueClient *sqs.SQS) *Gateway {
	return &Gateway{
		queueURL: os.Getenv("QUEUE_URL"),
		queue:    queueClient,
	}
}

func (g *Gateway) SendMessage(producaopedido *entities.ProducaoPedido) error {
	// Convert the struct to a JSON string
	jsonString, err := json.Marshal(producaopedido)
	if err != nil {
		fmt.Printf("Error parsing order to json string: %s\n", err)
		return err
	}

	stringMessage := string(jsonString)
	fmt.Printf("Sending message: %s\n", jsonString)

	//Build message
	message := &sqs.SendMessageInput{
		QueueUrl:    &g.queueURL,
		MessageBody: &stringMessage,
	}

	messageResult, err := g.queue.SendMessage(message)
	fmt.Printf("Message result: %s\n", messageResult)

	return nil
}
