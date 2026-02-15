package SubscriptionService

import (
	"Groundwork/backend/internal/domain"
	"context"
)

type SubscriptionService interface {
	CreateSubscriptionService(ctx context.Context, req *domain.CreateSubscriptionRequest) (*domain.SubscriptionResponse, error)
	GetSubscriptionService(ctx context.Context, req *domain.GetSubscriptionRequest) (*domain.SubscriptionResponse, error)
	ListSubscriptionsService(ctx context.Context) ([]*domain.SubscriptionResponse, error)
	UpdateSubscriptionService(ctx context.Context, req *domain.UpdateSubscriptionRequest) error
	DeleteSubscriptionService(ctx context.Context, req *domain.DeleteSubscriptionRequest) error
}
