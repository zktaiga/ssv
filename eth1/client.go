package eth1

import (
	"context"
	"math/big"

	"go.uber.org/zap"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/prysmaticlabs/prysm/v4/async/event"
)

//go:generate mockgen -package=eth1 -destination=./mock_client.go -source=./client.go

// Event represents an eth1 event log in the system
type Event struct {
	// Log is the raw event log
	Log types.Log
	// Name is the event name used for internal representation.
	Name string
	// Data is the parsed event
	Data interface{}
}

// SyncEndedEvent meant to notify an observer that the sync is over
type SyncEndedEvent struct {
	// Success returns true if the sync went well (all events were parsed)
	Success bool
	// Block is the block number of the last block synced
	Block uint64
}

// Client represents the required interface for eth1 client
type Client interface {
	EventsFeed() *event.Feed
	Start(logger *zap.Logger) error
	Sync(logger *zap.Logger, fromBlock *big.Int) error
	IsReady(ctx context.Context) (bool, error)
}
