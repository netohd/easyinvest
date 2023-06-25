package main

import (
	"easyinvest-microsservice/internal/infra/kafka"
	"easyinvest-microsservice/internal/market/dto"
	"easyinvest-microsservice/internal/market/entity"
	"easyinvest-microsservice/internal/market/transformer"
	"encoding/json"
	"fmt"
	"sync"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	ordersIn := make(chan *entity.Order)
	ordersOut := make(chan *entity.Order)
	wg := &sync.WaitGroup{}

	// É como se estivesse no final do código
	defer wg.Wait()

	kafkaMsgChan := make(chan *ckafka.Message)
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	}
	producer := kafka.NewKafkaProducer(configMap)
	kafka := kafka.NewConsumer(configMap, []string{"input"})

	go kafka.Consume(kafkaMsgChan) // Thread 2

	/* Recebe do canal do kafka, joga no input, processa,
	joga no output e depois publica no kafka */
	book := entity.NewBook(ordersIn, ordersOut, wg)
	go book.Trade() // Thread 3

	go func() {
		for msg := range kafkaMsgChan {
			wg.Add(1) // Wait Group faz um "done" p/ cada transação
			fmt.Println(string(msg.Value))
			tradeInput := dto.TradeInput{}
			err := json.Unmarshal(msg.Value, &tradeInput)
			if err != nil {
				panic(err)
			}
			order := transformer.TransformInput(tradeInput)
			ordersIn <- order
		}
	}()

	for res := range ordersOut {
		output := transformer.TransformOutput(res)
		outputJson, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			fmt.Println(err)
		}
		producer.Publish(outputJson, []byte("orders"), "output")
	}
}
