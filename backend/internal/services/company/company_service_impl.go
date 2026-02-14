package CompanyService

import (
	"Groundwork/backend/internal/database"
)

type companyService struct {
	queries *database.Database
}

func NewUserService(db *database.Database) CompanyService {
	return &companyService{
		queries: db,
	}
}

// CreateCompany implements [CompanyService].
func (c *companyService) CreateCompany() {
	panic("unimplemented")
}

// DeleteCompany implements [CompanyService].
func (c *companyService) DeleteCompany() {
	panic("unimplemented")
}

// GetCompany implements [CompanyService].
func (c *companyService) GetCompany() {
	panic("unimplemented")
}

// ListCompanies implements [CompanyService].
func (c *companyService) ListCompanies() {
	panic("unimplemented")
}

// ListCompanySubscriptions implements [CompanyService].
func (c *companyService) ListCompanySubscriptions() {
	panic("unimplemented")
}

// UpdateCompany implements [CompanyService].
func (c *companyService) UpdateCompany() {
	panic("unimplemented")
}
