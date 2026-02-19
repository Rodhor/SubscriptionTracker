package CompanyCommand

import (
	commandUtils "Groundwork/backend/internal/commands"
	"Groundwork/backend/internal/domain"
	CompanyService "Groundwork/backend/internal/services/company"
	"context"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type companyCommand struct {
	service CompanyService.CompanyService
}

func NewCompanyCommand(svc CompanyService.CompanyService) CompanyCommand {
	return &companyCommand{
		service: svc,
	}
}

// CreateCompanyCommand returns a specific pointer to CompanyResponse
func (c *companyCommand) CreateCompanyCommand(req *domain.CreateCompanyRequest) *commandUtils.Response[*domain.CompanyResponse] {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := validate.Struct(req); err != nil {
		return commandUtils.ErrorResponseTyped[*domain.CompanyResponse](err)
	}

	resp, err := c.service.CreateCompanyService(ctx, req)
	if err != nil {
		return commandUtils.ErrorResponseTyped[*domain.CompanyResponse](err)
	}

	return commandUtils.SuccessResponse("Company created successfully", http.StatusCreated, resp)
}

// DeleteCompanyCommand - Returns [any] because the data payload is nil on success
func (c *companyCommand) DeleteCompanyCommand(req *domain.DeleteCompanyRequest) *commandUtils.Response[any] {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := validate.Struct(req); err != nil {
		return commandUtils.ErrorResponse(err)
	}

	err := c.service.DeleteCompanyService(ctx, req)
	if err != nil {
		return commandUtils.ErrorResponse(err)
	}

	return commandUtils.SuccessResponse[any]("Company deleted successfully", http.StatusOK, nil)
}

// GetCompanyCommand returns a specific pointer
func (c *companyCommand) GetCompanyCommand(req *domain.GetCompanyRequest) *commandUtils.Response[*domain.CompanyResponse] {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := validate.Struct(req); err != nil {
		return commandUtils.ErrorResponseTyped[*domain.CompanyResponse](err)
	}

	resp, err := c.service.GetCompanyService(ctx, req)
	if err != nil {
		return commandUtils.ErrorResponseTyped[*domain.CompanyResponse](err)
	}

	return commandUtils.SuccessResponse("Company retrieved successfully", http.StatusOK, resp)
}

// ListCompaniesCommand returns a slice
func (c *companyCommand) ListCompaniesCommand() *commandUtils.Response[[]*domain.CompanyResponse] {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := c.service.ListCompaniesService(ctx)
	if err != nil {
		return commandUtils.ErrorResponseTyped[[]*domain.CompanyResponse](err)
	}

	return commandUtils.SuccessResponse("Companies retrieved successfully", http.StatusOK, resp)
}

// ListCompanySubscriptionsCommand returns a slice of SubscriptionResponse
func (c *companyCommand) ListCompanySubscriptionsCommand(req *domain.ListCompanySubscriptionsRequest) *commandUtils.Response[[]*domain.SubscriptionResponse] {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := validate.Struct(req); err != nil {
		return commandUtils.ErrorResponseTyped[[]*domain.SubscriptionResponse](err)
	}

	resp, err := c.service.ListCompanySubscriptionsService(ctx, req)
	if err != nil {
		return commandUtils.ErrorResponseTyped[[]*domain.SubscriptionResponse](err)
	}

	return commandUtils.SuccessResponse("Company subscriptions retrieved successfully", http.StatusOK, resp)
}

// UpdateCompanyCommand - Typically returns [any] since data is nil after update
func (c *companyCommand) UpdateCompanyCommand(req *domain.UpdateCompanyRequest) *commandUtils.Response[any] {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := validate.Struct(req); err != nil {
		return commandUtils.ErrorResponse(err)
	}

	if err := c.service.UpdateCompanyService(ctx, req); err != nil {
		return commandUtils.ErrorResponse(err)
	}

	return commandUtils.SuccessResponse[any]("Company updated successfully", http.StatusOK, nil)
}
