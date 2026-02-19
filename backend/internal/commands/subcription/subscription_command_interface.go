package SubscriptionCommand

import (
	commandUtils "Groundwork/backend/internal/commands"
	"Groundwork/backend/internal/domain"
)

type SubscriptionCommand interface {
	// Returns the newly created subscription
	CreateSubscriptionCommand(req *domain.CreateSubscriptionRequest) *commandUtils.Response[*domain.SubscriptionResponse]

	// Returns a specific subscription
	GetSubscriptionCommand(req *domain.GetSubscriptionRequest) *commandUtils.Response[*domain.SubscriptionResponse]

	// Returns a list of all subscriptions
	ListSubscriptionsCommand() *commandUtils.Response[[]*domain.SubscriptionResponse]

	// Returns any (nil data) because you opted for no return upon updating
	UpdateSubscriptionCommand(req *domain.UpdateSubscriptionRequest) *commandUtils.Response[any]

	// Returns any (nil data) as is standard for deletion
	DeleteSubscriptionCommand(req *domain.DeleteSubscriptionRequest) *commandUtils.Response[any]
}
