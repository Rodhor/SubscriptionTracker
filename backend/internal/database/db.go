package database

import (
	"Groundwork/backend/internal/domain"
	"sync"

	"github.com/google/uuid"
)

type Database struct {
	mu            sync.RWMutex
	companies     map[uuid.UUID]*domain.Company
	subscriptions map[uuid.UUID]*domain.Subscription
}

func NewDB() *Database {
	return &Database{
		companies:     make(map[uuid.UUID]*domain.Company),
		subscriptions: make(map[uuid.UUID]*domain.Subscription),
	}
}
