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

// CreateSubscriptionCommand returns a specific pointer
func (s *subscriptionCommand) CreateSubscriptionCommand(req *domain.CreateSubscriptionRequest) *commandUtils.Response[*domain.SubscriptionResponse] {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := validate.Struct(req); err != nil {
		return commandUtils.ErrorResponseTyped[*domain.SubscriptionResponse](err)
	}

	resp, err := s.service.CreateSubscriptionService(ctx, req)
	if err != nil {
		return commandUtils.ErrorResponseTyped[*domain.SubscriptionResponse](err)
	}
	return commandUtils.SuccessResponse("Subscription created successfully", http.StatusCreated, resp)
}

// DeleteSubscriptionCommand returns [any] (nil data)
func (s *subscriptionCommand) DeleteSubscriptionCommand(req *domain.DeleteSubscriptionRequest) *commandUtils.Response[any] {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := validate.Struct(req); err != nil {
		return commandUtils.ErrorResponse(err)
	}

	err := s.service.DeleteSubscriptionService(ctx, req)
	if err != nil {
		return commandUtils.ErrorResponse(err)
	}
	return commandUtils.SuccessResponse[any]("Subscription deleted successfully", http.StatusOK, nil)
}

// GetSubscriptionCommand returns a specific pointer
func (s *subscriptionCommand) GetSubscriptionCommand(req *domain.GetSubscriptionRequest) *commandUtils.Response[*domain.SubscriptionResponse] {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := validate.Struct(req); err != nil {
		return commandUtils.ErrorResponseTyped[*domain.SubscriptionResponse](err)
	}

	resp, err := s.service.GetSubscriptionService(ctx, req)
	if err != nil {
		return commandUtils.ErrorResponseTyped[*domain.SubscriptionResponse](err)
	}
	return commandUtils.SuccessResponse("Subscription retrieved successfully", http.StatusOK, resp)
}

// ListSubscriptionsCommand returns a slice
func (s *subscriptionCommand) ListSubscriptionsCommand() *commandUtils.Response[[]*domain.SubscriptionResponse] {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := s.service.ListSubscriptionsService(ctx)
	if err != nil {
		return commandUtils.ErrorResponseTyped[[]*domain.SubscriptionResponse](err)
	}
	return commandUtils.SuccessResponse("Subscriptions retrieved successfully", http.StatusOK, resp)
}

// UpdateSubscriptionCommand returns [any] (nil data)
func (s *subscriptionCommand) UpdateSubscriptionCommand(req *domain.UpdateSubscriptionRequest) *commandUtils.Response[any] {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := validate.Struct(req); err != nil {
		return commandUtils.ErrorResponse(err)
	}

	if err := s.service.UpdateSubscriptionService(ctx, req); err != nil {
		return commandUtils.ErrorResponse(err)
	}
	return commandUtils.SuccessResponse[any]("Subscription updated successfully", http.StatusOK, nil)
}
