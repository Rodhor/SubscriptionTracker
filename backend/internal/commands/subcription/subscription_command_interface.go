package SubscriptionCommand

type SubscriptionCommand interface {
	CreateSubscriptionCommand()
	GetSubscriptionCommand()
	ListSubscriptionsCommand()
	UpdateSubscriptionCommand()
	DeleteSubscriptionCommand()
}
