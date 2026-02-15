package SubscriptionService

import (
	"Groundwork/backend/internal/database"
	"Groundwork/backend/internal/domain"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrUnauthorized   = errors.New("unauthorized")
	ErrInvalidRequest = errors.New("invalid request")
)

type subscriptionService struct {
	queries *database.Database
}

func NewSubscriptionService(db *database.Database) SubscriptionService {
	return &subscriptionService{
		queries: db,
	}
}

// CreateSubscriptionService implements [SubscriptionService].
func (s *subscriptionService) CreateSubscriptionService(ctx context.Context, req *domain.CreateSubscriptionRequest) (*domain.SubscriptionResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	now := time.Now()

	// Set Defaults
	base := domain.SubscriptionBase{
		Name:      req.Name,
		License:   req.License,
		Tier:      domain.TierCore,
		Status:    domain.StatusPending,
		StartDate: now,
		Features:  []string{},
	}

	// Apply overrides from request
	if req.Tier != nil {
		base.Tier = *req.Tier
	}
	if req.Status != nil {
		base.Status = *req.Status
	}
	if req.Features != nil {
		base.Features = *req.Features
	}
	if req.StartDate != nil {
		base.StartDate = *req.StartDate
	}
	if req.EndDate != nil {
		base.EndDate = req.EndDate
	}
	if req.AutoRenew != nil {
		base.AutoRenew = *req.AutoRenew
	}

	newSub := &domain.Subscription{
		ID:               uuid.New(),
		CompanyID:        req.CompanyID,
		SubscriptionBase: base,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	if err := s.queries.CreateSubscription(ctx, newSub); err != nil {
		return nil, err
	}

	return newSub.Sanitize(), nil
}

// DeleteSubscriptionService implements [SubscriptionService].
func (s *subscriptionService) DeleteSubscriptionService(ctx context.Context, req *domain.DeleteSubscriptionRequest) error {
	// Check context for cancellation
	if err := ctx.Err(); err != nil {
		return err
	}

	// get timestamp
	now := time.Now()
	if err := s.queries.DeleteSubscription(ctx, req.ID, now); err != nil {
		return err
	}

	// if deletion was successful, return nil
	return nil
}

// GetSubscriptionService implements [SubscriptionService].
func (s *subscriptionService) GetSubscriptionService(ctx context.Context, req *domain.GetSubscriptionRequest) (*domain.SubscriptionResponse, error) {
	// Check context for cancellation
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// try to retrive subscription from database returning early on error
	subscription, err := s.queries.GetSubscription(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	// if successful, return sanitized subscription
	return subscription.Sanitize(), nil
}

// ListSubscriptionsService implements [SubscriptionService].
func (s *subscriptionService) ListSubscriptionsService(ctx context.Context) ([]*domain.SubscriptionResponse, error) {
	// Check context for cancellation
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// try to retrieve subscriptions from database returning early on error
	subscriptions, err := s.queries.ListSubscriptions(ctx)
	if err != nil {
		return nil, err
	}

	// if successful, sanitize and return subscriptions
	sanitizedSubscriptions := make([]*domain.SubscriptionResponse, 0, len(subscriptions))
	for _, subscription := range subscriptions {
		sanitizedSubscriptions = append(sanitizedSubscriptions, subscription.Sanitize())
	}
	return sanitizedSubscriptions, nil
}

// UpdateSubscriptionService implements [SubscriptionService].
func (s *subscriptionService) UpdateSubscriptionService(ctx context.Context, req *domain.UpdateSubscriptionRequest) error {
	// Check context for cancellation
	if err := ctx.Err(); err != nil {
		return err
	}

	// get timestamp
	now := time.Now()
	if err := s.queries.UpdateSubscription(ctx, req, now); err != nil {
		return err
	}

	return nil
}
