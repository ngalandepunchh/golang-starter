package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type EventStore struct {
	BrokerURL string
	Producer  *kafka.Producer
	Consumer  *kafka.Consumer
}

// NewEventStore - constructor
func NewEventStore(url string) (*EventStore, error) {

	// producer config
	kakfaProdConfig := kafka.ConfigMap{"bootstrap.servers": url,}

	// init the producer
	producer, err := kafka.NewProducer(&kakfaProdConfig )

	if err != nil {
		fmt.Errorf("cannot get a kafka producer %v", err)
		return nil, err
	}

	kafkaConsumerConfig := kafka.ConfigMap{
		"bootstrap.servers": url,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	}

	// init the consumer
	consumer, err := kafka.NewConsumer(&kafkaConsumerConfig)

	if err != nil {
		fmt.Errorf("cannot get a kafka consumer %v", err)
		return nil, err
	}

	return &EventStore{BrokerURL: url, Producer: producer, Consumer: consumer}, nil
}

type EventHandler interface {
	Publish(msg, topic string , c chan string) error
	Subscribe(topic string, c chan string) (string, error)
}

func (e *EventStore) Publish(msg, topic string, c chan string) error {
	
	// create a message

	fmt.Println("Creating kafka message")

	kafkaMsg := kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value: []byte(msg)}

	// send the msg, delivery chan is null for now

	fmt.Printf("Sending message %s for topic %s  " , msg , topic)
	e.Producer.Produce(&kafkaMsg, nil)


	// Wait for message deliveries before shutting down
	e.Producer.Flush(15 * 1000)
	//
	//fmt.Printf("Publishing message to channel %s", msg)
	//
	//c <- msg

	fmt.Println("Finished publishing...")

	e.Producer.Close()

	return nil

}

func (e *EventStore) Subscribe(topic string, c chan string) (string, error) {

	fmt.Println("Consuming from topic " , topic)

	e.Consumer.SubscribeTopics([]string{topic, "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := e.Consumer.ReadMessage(-1)
		fmt.Println("Received msg")
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			return string(msg.Value),nil
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			return "",err
		}
	}

	//e.Consumer.Close()

	//res := <-c

}
