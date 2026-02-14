package SubscriptionCommand

import (
	SubscriptionService "Groundwork/backend/internal/services/subscription"
)

type subscriptionCommand struct {
	service *SubscriptionService.SubscriptionService
}

func NewSubscriptionCommand(svc *SubscriptionService.SubscriptionService) SubscriptionCommand {
	return &subscriptionCommand{
		service: svc,
	}
}

// CreateSubscriptionCommand implements [SubscriptionCommand].
func (s *subscriptionCommand) CreateSubscriptionCommand() {
	panic("unimplemented")
}

// DeleteSubscriptionCommand implements [SubscriptionCommand].
func (s *subscriptionCommand) DeleteSubscriptionCommand() {
	panic("unimplemented")
}

// GetSubscriptionCommand implements [SubscriptionCommand].
func (s *subscriptionCommand) GetSubscriptionCommand() {
	panic("unimplemented")
}

// ListSubscriptionsCommand implements [SubscriptionCommand].
func (s *subscriptionCommand) ListSubscriptionsCommand() {
	panic("unimplemented")
}

// UpdateSubscriptionCommand implements [SubscriptionCommand].
func (s *subscriptionCommand) UpdateSubscriptionCommand() {
	panic("unimplemented")
}
