package SubscriptionCommand

import (
	commandUtils "Groundwork/backend/internal/commands"
	"Groundwork/backend/internal/domain"
	SubscriptionService "Groundwork/backend/internal/services/subscription"
	"context"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type subscriptionCommand struct {
	service SubscriptionService.SubscriptionService
}

func NewSubscriptionCommand(svc SubscriptionService.SubscriptionService) SubscriptionCommand {
	return &subscriptionCommand{
		service: svc,
	}
}

// CreateSubscriptionCommand implements [SubscriptionCommand].
func (s *subscriptionCommand) CreateSubscriptionCommand(req *domain.CreateSubscriptionRequest) *commandUtils.Response {
	// Create context with timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Validate request
	if err := validate.Struct(req); err != nil {
		return commandUtils.ErrorResponse(err)
	}

	// If request is valid, proceed with creating subscription
	resp, err := s.service.CreateSubscriptionService(ctx, req)
	if err != nil {
		return commandUtils.ErrorResponse(err)
	}
	return commandUtils.SuccessResponse("Subscription created successfully", http.StatusCreated, resp)
}

// DeleteSubscriptionCommand implements [SubscriptionCommand].
func (s *subscriptionCommand) DeleteSubscriptionCommand(req *domain.DeleteSubscriptionRequest) *commandUtils.Response {
	// Create context with timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Validate request
	if err := validate.Struct(req); err != nil {
		return commandUtils.ErrorResponse(err)
	}

	// If request is valid, proceed with deleting subscription
	err := s.service.DeleteSubscriptionService(ctx, req)
	if err != nil {
		return commandUtils.ErrorResponse(err)
	}
	return commandUtils.SuccessResponse("Subscription deleted successfully", http.StatusOK, nil)
}

// GetSubscriptionCommand implements [SubscriptionCommand].
func (s *subscriptionCommand) GetSubscriptionCommand(req *domain.GetSubscriptionRequest) *commandUtils.Response {
	// Create context with timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Validate request
	if err := validate.Struct(req); err != nil {
		return commandUtils.ErrorResponse(err)
	}

	// If request is valid, proceed with getting subscription
	resp, err := s.service.GetSubscriptionService(ctx, req)
	if err != nil {
		return commandUtils.ErrorResponse(err)
	}
	return commandUtils.SuccessResponse("Subscription retrieved successfully", http.StatusOK, resp)
}

// ListSubscriptionsCommand implements [SubscriptionCommand].
func (s *subscriptionCommand) ListSubscriptionsCommand() *commandUtils.Response {
	// Create context with timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := s.service.ListSubscriptionsService(ctx)
	if err != nil {
		return commandUtils.ErrorResponse(err)
	}
	return commandUtils.SuccessResponse("Subscriptions retrieved successfully", http.StatusOK, resp)
}

// UpdateSubscriptionCommand implements [SubscriptionCommand].
func (s *subscriptionCommand) UpdateSubscriptionCommand(req *domain.UpdateSubscriptionRequest) *commandUtils.Response {
	// Create context with timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Validate request
	if err := validate.Struct(req); err != nil {
		return commandUtils.ErrorResponse(err)
	}

	// If request is valid, proceed with updating subscription
	if err := s.service.UpdateSubscriptionService(ctx, req); err != nil {
		return commandUtils.ErrorResponse(err)
	}
	return commandUtils.SuccessResponse("Subscription updated successfully", http.StatusOK, nil)
}
