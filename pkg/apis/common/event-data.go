/*
Copyright 2020 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package common

import (
	"encoding/json"
	"github.com/minio/minio-go"
	"net/http"

	sqslib "github.com/aws/aws-sdk-go/service/sqs"
)

// AMQPEventData represents the event data generated by AMQP gateway.
// +k8s:openapi-gen=true
type AMQPEventData struct {
	// ContentType is the MIME content type
	ContentType string `json:"contentType"`
	// ContentEncoding is the MIME content encoding
	ContentEncoding string `json:"contentEncoding"`
	// Delivery mode can be either - non-persistent (1) or persistent (2)
	DeliveryMode int `json:"deliveryMode"`
	// Priority refers to the use - 0 to 9
	Priority int `json:"priority"`
	// CorrelationId is the correlation identifier
	CorrelationId string `json:"correlationId"`
	// ReplyTo is the address to reply to (ex: RPC)
	ReplyTo string `json:"replyTo"`
	// Expiration refers to message expiration spec
	Expiration string `json:"expiration"`
	// MessageId is message identifier
	MessageId string `json:"messageId"`
	// Timestamp refers to the message timestamp
	Timestamp string `json:"timestamp"`
	// Type refers to the message type name
	Type string `json:"type"`
	// AppId refers to the application id
	AppId string `json:"appId"`
	// Exchange is basic.publish exchange
	Exchange string `json:"exchange"`
	// RoutingKey is basic.publish routing key
	RoutingKey string `json:"routingKey"`
	// Body represents the messsage body
	Body []byte `json:"body"`
}

// SNSEventData represents the event data generated by SNS gateway.
// +k8s:openapi-gen=true
type SNSEventData struct {
	// Body represents the SNS message body
	Body []byte `json:"body"`
}

// SQSEventData represents the event data generated by SQS gateway.
// +k8s:openapi-gen=true
type SQSEventData struct {
	// A unique identifier for the message. A MessageId is considered unique across
	// all AWS accounts for an extended period of time.
	MessageId string `json:"messageId"`
	// Each message attribute consists of a Name, Type, and Value. For more information,
	// see Amazon SQS Message Attributes (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-message-attributes.html)
	// in the Amazon Simple Queue Service Developer Guide.
	MessageAttributes map[string]*sqslib.MessageAttributeValue `json:"messageAttributes"`
	// The message's contents (not URL-encoded).
	Body []byte `json:"body"`
}

// AzureEventsHubEventData represents to the event data generated by Azure Events Hub gateway.
// +k8s:openapi-gen=true
type AzureEventsHubEventData struct {
	// Id of the message
	Id string `json:"id"`
	// PartitionKey
	PartitionKey string `json:"partitionKey"`
	// Message body
	Body []byte `json:"body"`
}

// CalendarEventData represents the event data generated by the Calendar gateway.
// +k8s:openapi-gen=true
type CalendarEventData struct {
	// EventTime is time at which event occurred
	EventTime string `json:"eventTime"`
	// UserPayload if any
	// +optional
	UserPayload *json.RawMessage `json:"userPayload,omitempty"`
}

// EmitterEventData represents the event data generated by the Emitter gateway.
type EmitterEventData struct {
	// Topic name
	Topic string `json:"topic"`
	// Body represents the message body
	Body []byte `json:"body"`
}

// PubSubEventData represents the event data generated by the GCP PubSub gateway.
type PubSubEventData struct {
	// ID of the message
	ID string `json:"id"`
	// Body is the actual data in the message.
	Body []byte `json:"body"`
	// Attributes represents the key-value pairs the current message
	// is labelled with.
	Attributes map[string]string `json:"attributes"`
	// The time at which the message was published.
	PublishTime string `json:"publishTime"`
}

// GithubEventData represents the event data generated by the GitHub gateway.
type GithubEventData struct {
	// Body represents the message body
	Body []byte `json:"body"`
}

// GitLabEventData represents the event data generated by the GitLab gateway.
type GitLabEventData struct {
	// Body represents the message body
	Body []byte `json:"body"`
}

// KafkaEventData represents the event data generated by the Kafka gateway.
type KafkaEventData struct {
	// Topic refers to the Kafka topic
	Topic string `json:"topic"`
	// Partition refers to the Kafka partition
	Partition int `json:"partition"`
	// Body refers to the message value
	Body []byte `json:"value"`
	// Timestamp of the message
	Timestamp string `json:"timestamp"`
}

// MinioEventData represents the event data generated by the Minio gateway.
type MinioEventData struct {
	Notification []minio.NotificationEvent `json:"notification"`
}

// MQTTEventData represents the event data generated by the MQTT gateway.
type MQTTEventData struct {
	// Topic refers to the MQTT topic name.
	Topic string `json:"topic"`
	// MessageId is the unique ID for the message
	MessageId int `json:"messageId"`
	// Payload is the message payload.
	Body []byte `json:"payload"`
}

// NATSEventData represents the event data generated by the NATS gateway.
type NATSEventData struct {
	// Name of the subject.
	Subject string `json:"subject"`
	// Message data.
	Body []byte `json:"data"`
}

// NSQEventData represents the event data generated by the NSQ gateway.
type NSQEventData struct {
	// Body is the message data.
	Body []byte
	// Timestamp of the message.
	Timestamp string
	// NSQDAddress is the address of the nsq host.
	NSQDAddress string
}

// RedisEventData represents the event data generated by the Redis gateway.
type RedisEventData struct {
	// Subscription channel.
	Channel string `json:"channel"`
	// Message pattern
	Pattern string `json:"pattern"`
	// Message body
	Body string `json:"payload"`
}

// ResourceEventData represents the event data generated by the Resource gateway.
type ResourceEventData struct {
	// EventType of the type of the event.
	EventType string `json:"type"`
	// Resource body.
	Body []byte `json:"body"`
	// Resource group name.
	Group string `json:"group"`
	// Resource version.
	Version string `json:"version"`
	// Resource name.
	Resource string `json:"resource"`
}

// WebhookEventData represents the event data generated by the Webhook gateway.
type WebhookEventData struct {
	// Header is the http request header
	Header http.Header `json:"header"`
	// Body is http request body
	Body *json.RawMessage `json:"body"`
}
