package CompanyCommand

import (
	commandUtils "Groundwork/backend/internal/commands"
	"Groundwork/backend/internal/domain"
)

type CompanyCommand interface {
	// Returns the newly created company
	CreateCompanyCommand(req *domain.CreateCompanyRequest) *commandUtils.Response[*domain.CompanyResponse]

	// Returns a specific company
	GetCompanyCommand(req *domain.GetCompanyRequest) *commandUtils.Response[*domain.CompanyResponse]

	// Returns a list of companies
	ListCompaniesCommand() *commandUtils.Response[[]*domain.CompanyResponse]

	// Returns a list of subscriptions for a company
	ListCompanySubscriptionsCommand(req *domain.ListCompanySubscriptionsRequest) *commandUtils.Response[[]*domain.SubscriptionResponse]

	// Returns the updated company (Best practice so frontend doesn't have to re-fetch)
	UpdateCompanyCommand(req *domain.UpdateCompanyRequest) *commandUtils.Response[any]

	// Returns a generic response (Data is usually nil/null after deletion)
	DeleteCompanyCommand(req *domain.DeleteCompanyRequest) *commandUtils.Response[any]
}
