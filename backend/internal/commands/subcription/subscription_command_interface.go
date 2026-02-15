package SubscriptionCommand

import (
	commandUtils "Groundwork/backend/internal/commands"
	"Groundwork/backend/internal/domain"
)

type SubscriptionCommand interface {
	CreateSubscriptionCommand(req *domain.CreateSubscriptionRequest) *commandUtils.Response
	GetSubscriptionCommand(req *domain.GetSubscriptionRequest) *commandUtils.Response
	ListSubscriptionsCommand() *commandUtils.Response
	UpdateSubscriptionCommand(req *domain.UpdateSubscriptionRequest) *commandUtils.Response
	DeleteSubscriptionCommand(req *domain.DeleteSubscriptionRequest) *commandUtils.Response
}
