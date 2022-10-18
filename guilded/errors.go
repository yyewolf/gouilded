package guilded

import "errors"

var (
	ErrGatewayAlreadyConnected = errors.New("gateway: already connected")
	ErrGatewayNotConnected     = errors.New("gateway: not connected")
)
