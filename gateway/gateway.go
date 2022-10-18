package gateway

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/yyewolf/gouilded/guilded"
)

// Status is the state that the client is currently in.
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

type Gateway struct {
	config *Config
	token  string

	conn            *websocket.Conn
	connMu          sync.Mutex
	heartbeatTicker *time.Ticker
	status          Status

	heartbeatInterval     time.Duration
	lastHeartbeatSent     time.Time
	lastHeartbeatReceived time.Time
}

func (g *Gateway) Open(ctx context.Context) error {
	g.connMu.Lock()
	defer g.connMu.Unlock()

	// Check if we're already connected
	if g.conn != nil {
		return guilded.ErrGatewayAlreadyConnected
	}
	// Set the status to connecting
	g.status = StatusConnecting

	wsUrl := g.config.url
	// dial with custom headers
	headers := http.Header(make(map[string][]string))
	headers.Set("Authorization", fmt.Sprintf("Bearer %s", g.token))
	headers.Set("User-Agent", fmt.Sprintf("gouilded-%s", g.config.version))
	headers.Set("guilded-last-message-id", g.config.lastID)
	conn, rs, err := websocket.DefaultDialer.DialContext(ctx, wsUrl, headers)
	if err != nil {
		g.Close(ctx)
		// body := "null"
		if rs != nil && rs.Body != nil {
			defer func() {
				_ = rs.Body.Close()
			}()
			// rawBody, bErr := io.ReadAll(rs.Body)
			_, bErr := io.ReadAll(rs.Body)
			if bErr != nil {
				// TODO: log error while reading response body
			}
			// body = string(rawBody)
		}

		// TODO: log error while connecting to gateway
		return err
	}

	g.conn = conn
	g.status = StatusWaitingForHello

	// go g.readLoop(ctx)

	return nil
}

func (g *Gateway) Close(ctx context.Context) {
	g.CloseWithCode(ctx, websocket.CloseNormalClosure, "Shutting down")
}

func (g *Gateway) CloseWithCode(ctx context.Context, code int, message string) {
	if g.heartbeatTicker != nil {
		// TODO: Log closing heartbeat ticker
		g.heartbeatTicker.Stop()
		g.heartbeatTicker = nil
	}

	g.connMu.Lock()
	defer g.connMu.Unlock()
	if g.conn != nil {
		g.config.ratelimiter.Close(ctx)
		// TODO: Log closing connection with code & message
		if err := g.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(code, message)); err != nil && err != websocket.ErrCloseSent {
			// TODO: Log error while closing connection
		}
		_ = g.conn.Close()
		g.conn = nil

		// clear resume data as we closed gracefully
		if code == websocket.CloseNormalClosure || code == websocket.CloseGoingAway {
			g.config.lastID = ""
		}
	}
}

func (g *Gateway) Status() Status {
	g.connMu.Lock()
	defer g.connMu.Unlock()
	return g.status
}

func (g *Gateway) send(ctx context.Context, messageType int, data []byte) error {
	g.connMu.Lock()
	defer g.connMu.Unlock()
	if g.conn == nil {
		return guilded.ErrGatewayNotConnected
	}

	if err := g.config.ratelimiter.Wait(ctx); err != nil {
		return err
	}

	defer g.config.ratelimiter.Unlock()
	// TODO: Log sending message
	return g.conn.WriteMessage(messageType, data)
}

func (g *Gateway) reconnectTry(ctx context.Context, try int, delay time.Duration) error {
	if try >= g.config.maxReconnectTries-1 {
		return fmt.Errorf("failed to reconnect. exceeded max reconnect tries of %d reached", g.config.maxReconnectTries)
	}
	timer := time.NewTimer(time.Duration(try) * delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		timer.Stop()
		return ctx.Err()
	case <-timer.C:
	}

	// TODO: log reconnecting
	if err := g.Open(ctx); err != nil {
		if err == guilded.ErrGatewayAlreadyConnected {
			return err
		}
		// TODO: log error while reconnecting
		g.status = StatusDisconnected
		return g.reconnectTry(ctx, try+1, delay)
	}
	return nil
}
