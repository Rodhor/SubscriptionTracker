package CompanyCommand

import (
	CompanyService "Groundwork/backend/internal/services/company"
)

type companyCommand struct {
	service *CompanyService.CompanyService
}

func NewCompanyCommand(svc *CompanyService.CompanyService) CompanyCommand {
	return &companyCommand{
		service: svc,
	}
}

// CreateCompanyCommand implements [CompanyCommand].
func (c *companyCommand) CreateCompanyCommand() {
	panic("unimplemented")
}

// DeleteCompanyCommand implements [CompanyCommand].
func (c *companyCommand) DeleteCompanyCommand() {
	panic("unimplemented")
}

// GetCompanyCommand implements [CompanyCommand].
func (c *companyCommand) GetCompanyCommand() {
	panic("unimplemented")
}

// ListCompaniesCommand implements [CompanyCommand].
func (c *companyCommand) ListCompaniesCommand() {
	panic("unimplemented")
}

// ListCompanySubscriptionsCommand implements [CompanyCommand].
func (c *companyCommand) ListCompanySubscriptionsCommand() {
	panic("unimplemented")
}

// UpdateCompanyCommand implements [CompanyCommand].
func (c *companyCommand) UpdateCompanyCommand() {
	panic("unimplemented")
}
