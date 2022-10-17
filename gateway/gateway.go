package gateway

import (
	"context"
	"time"
)

type Status int

// IsConnected returns whether the Gateway is connected.
func (s Status) IsConnected() bool {
	switch s {
	case StatusWaitingForHello, StatusIdentifying, StatusWaitingForReady, StatusReady:
		return true
	default:
		return false
	}
}

// Indicates how far along the client is to connecting.
const (
	StatusUnconnected Status = iota
	StatusConnecting
	StatusWaitingForHello
	StatusIdentifying
	StatusResuming
	StatusWaitingForReady
	StatusReady
	StatusDisconnected
)

type (
	// EventHandlerFunc is a function that is called when an event is received.
	// EventHandlerFunc func(gatewayEventType EventType, sequenceNumber int, shardID int, event EventData)

	// CreateFunc is a type that is used to create a new Gateway(s).
	// CreateFunc func(token string, eventHandlerFunc EventHandlerFunc, closeHandlerFUnc CloseHandlerFunc, opts ...ConfigOpt) Gateway

	// CloseHandlerFunc is a function that is called when the Gateway is closed.
	CloseHandlerFunc func(gateway Gateway, err error)
)

type Gateway interface {
	// LastMessageReceived returns the last message id received by the Gateway.
	LastMessageReceived() *int

	// Open connects this Gateway to the Guilded API
	Open(ctx context.Context) error

	// Close gracefully closes the Gateway with the websocket.CloseNormalClosure code.
	// If the context is done, the Gateway connection will be killed.
	Close(ctx context.Context)

	// CloseWithCode closes the Gateway with the given code & message.
	// If the context is done, the Gateway connection will be killed.
	CloseWithCode(ctx context.Context, code int, message string)

	// Status returns the Status of the Gateway.
	Status() Status

	// Send sends a message to the Guilded gateway with the opCode and data.
	// If context is deadline exceeds, the message sending will be aborted.
	// TODO :
	// Send(ctx context.Context, op Opcode, data MessageData) error

	// Latency returns the latency of the Gateway.
	// This is calculated by the time it takes to send a heartbeat and receive a heartbeat ack by Guilded.
	Latency() time.Duration
}
