package eventbus

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestPubSub(t *testing.T) {

	topic := "test_topic"
	eventMsg := "test-msg"
	var msgGroup []string

	PubSub().Sub(context.TODO(), topic, func(event interface{}) {
		msg := event.(string)

		fmt.Println("consumer 1 -- msg : ", msg)
		msgGroup = append(msgGroup, msg+"1")

	})

	PubSub().Sub(context.TODO(), topic, func(event interface{}) {
		msg := event.(string)
		fmt.Println("consumer 2 -- msg : ", msg)
		msgGroup = append(msgGroup, msg+"2")

	})

	PubSub().Pub(topic).Publish(eventMsg)

	time.Sleep(2 * time.Second)

	assert.Len(t, msgGroup, 2)
	sort.Strings(msgGroup)
	assert.True(t, reflect.DeepEqual(msgGroup, []string{eventMsg + "1", eventMsg + "2"}))
}
