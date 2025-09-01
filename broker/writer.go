package broker

import (
	"github.com/Yarik7610/library-backend-common/sharedconstants"
	"github.com/segmentio/kafka-go"
)

func NewWriter(topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:  kafka.TCP(sharedconstants.KAFKA_NODE_1_ADDR, sharedconstants.KAFKA_NODE_2_ADDR, sharedconstants.KAFKA_NODE_3_ADDR),
		Topic: topic,
	}
}
