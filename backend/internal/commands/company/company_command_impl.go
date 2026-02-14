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

// CreateCompanyCommand implements [CompanyCommand].
func (c *companyCommand) CreateCompanyCommand(req *domain.CreateCompanyRequest) *commandUtils.Response {
	// Create context with timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Validate request
	if err := validate.Struct(req); err != nil {
		return commandUtils.ErrorResponse(err)
	}

	// If request is valid, proceed with creating company
	resp, err := c.service.CreateCompanyService(ctx, req)
	if err != nil {
		return commandUtils.ErrorResponse(err)
	}

	// If company is created successfully, return success response
	return commandUtils.SuccessResponse("Company created successfully", http.StatusCreated, resp)
}

// DeleteCompanyCommand implements [CompanyCommand].
func (c *companyCommand) DeleteCompanyCommand(req *domain.DeleteCompanyRequest) *commandUtils.Response {
	// Create context with timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Validate request
	if err := validate.Struct(req); err != nil {
		return commandUtils.ErrorResponse(err)
	}

	// If request is valid, proceed with deleting company
	err := c.service.DeleteCompanyService(ctx, req)
	if err != nil {
		return commandUtils.ErrorResponse(err)
	}

	// If company is deleted successfully, return success response
	return commandUtils.SuccessResponse("Company deleted successfully", http.StatusOK, nil)
}

// GetCompanyCommand implements [CompanyCommand].
func (c *companyCommand) GetCompanyCommand(req *domain.GetCompanyRequest) *commandUtils.Response {
	// Create context with timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Validate request
	if err := validate.Struct(req); err != nil {
		return commandUtils.ErrorResponse(err)
	}

	// If request is valid, proceed with getting company
	resp, err := c.service.GetCompanyService(ctx, req)
	if err != nil {
		return commandUtils.ErrorResponse(err)
	}

	// If company is retrieved successfully, return success response
	return commandUtils.SuccessResponse("Company retrieved successfully", http.StatusOK, resp)
}

// ListCompaniesCommand implements [CompanyCommand].
func (c *companyCommand) ListCompaniesCommand() *commandUtils.Response {
	// Create context with timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := c.service.ListCompaniesService(ctx)
	if err != nil {
		return commandUtils.ErrorResponse(err)
	}

	// If companies are retrieved successfully, return success response
	return commandUtils.SuccessResponse("Companies retrieved successfully", http.StatusOK, resp)
}

// ListCompanySubscriptionsCommand implements [CompanyCommand].
func (c *companyCommand) ListCompanySubscriptionsCommand(req *domain.ListCompanySubscriptionsRequest) *commandUtils.Response {
	// Create context with timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Validate request parameters
	if err := validate.Struct(req); err != nil {
		return commandUtils.ErrorResponse(err)
	}

	// If request parameters are valid, proceed with service call
	resp, err := c.service.ListCompanySubscriptionsService(ctx, req)
	if err != nil {
		return commandUtils.ErrorResponse(err)
	}

	// If subscriptions are retrieved successfully, return success response
	return commandUtils.SuccessResponse("Company subscriptions retrieved successfully", http.StatusOK, resp)
}

// UpdateCompanyCommand implements [CompanyCommand].
func (c *companyCommand) UpdateCompanyCommand(req *domain.UpdateCompanyRequest) *commandUtils.Response {
	// Create context with timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Validate request parameters
	if err := validate.Struct(req); err != nil {
		return commandUtils.ErrorResponse(err)
	}

	// If request parameters are valid, proceed with service call
	if err := c.service.UpdateCompanyService(ctx, req); err != nil {
		return commandUtils.ErrorResponse(err)
	}

	// If company is updated successfully, return success response
	return commandUtils.SuccessResponse("Company updated successfully", http.StatusOK, nil)
}
