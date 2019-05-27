package kafka

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEventStore(t *testing.T) {

	testEventStore,err := NewEventStore("localhost")

	if err != nil {
		assert.Fail(t,"cannot get an event store")
	}

	assert.Equal(t,nil,testEventStore)

}

func TestEventPublishSubscribe(t *testing.T){

	testEventStore,err := NewEventStore("localhost:19092,localhost:29092,localhost:39092")

	if err != nil {
		assert.Fail(t,"cannot get an event store")
	}

	queue := make(chan string,1)

	// send a foo
	 testEventStore.Publish("foo","test-topic", queue)

	// get it back
	actual,err := testEventStore.Subscribe("test-topic",queue)

	assert.Equal(t,"foo",actual)


}

