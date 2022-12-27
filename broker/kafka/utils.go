package kafka

import (
	kafkago "github.com/segmentio/kafka-go"
	"github.com/tx7do/kratos-transport/broker"
)

func kafkaHeaderToMap(h []kafkago.Header) broker.Headers {
	m := broker.Headers{}
	for _, v := range h {
		m[v.Key] = string(v.Value)
	}
	return m
}
