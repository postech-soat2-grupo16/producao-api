package message

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/postech-soat2-grupo16/producao-api/entities"
	"log"
	"os"
)

type GatewayInterface interface {
	SendMessage(order *entities.ProducaoPedido) error
}

type Gateway struct {
	queueURL string
	queue    *sqs.SQS
}

type GatewayMock struct {
}

func NewGateway(queueClient *sqs.SQS) GatewayInterface {
	if queueClient == nil {
		return NewGatewayMock()
	}
	return &Gateway{
		queueURL: os.Getenv("QUEUE_URL"),
		queue:    queueClient,
	}
}

func NewGatewayMock() *GatewayMock {
	return &GatewayMock{}
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

	log.Printf("Enviando Mensagem - Produção: %v\n", message)
	messageResult, err := g.queue.SendMessage(message)
	log.Printf("Mensagem enviada Mensagem - Produção: %v\n", messageResult)

	return nil
}

func (g *GatewayMock) SendMessage(order *entities.ProducaoPedido) error {
	return nil
}
