package SubscriptionService

import (
	"Groundwork/backend/internal/database"
	"errors"
)

type subscriptionService struct {
	queries *database.Database
}

func NewSubscriptionService(db *database.Database) SubscriptionService {
	return &subscriptionService{
		queries: db,
	}
}

// CreateSubscription implements [SubscriptionService].
func (s *subscriptionService) CreateSubscription() {
	panic("unimplemented")
}

// DeleteSubscription implements [SubscriptionService].
func (s *subscriptionService) DeleteSubscription() {
	panic("unimplemented")
}

// GetSubscription implements [SubscriptionService].
func (s *subscriptionService) GetSubscription() {
	panic("unimplemented")
}

// ListSubscriptions implements [SubscriptionService].
func (s *subscriptionService) ListSubscriptions() {
	panic("unimplemented")
}

// UpdateSubscription implements [SubscriptionService].
func (s *subscriptionService) UpdateSubscription() {
	panic("unimplemented")
}

var (
	ErrUnauthorized   = errors.New("unauthorized")
	ErrInvalidRequest = errors.New("invalid request")
)
