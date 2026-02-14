package CompanyCommand

import (
	commandUtils "Groundwork/backend/internal/commands"
	"Groundwork/backend/internal/domain"
)

type CompanyCommand interface {
	CreateCompanyCommand(req *domain.CreateCompanyRequest) *commandUtils.Response
	GetCompanyCommand(req *domain.GetCompanyRequest) *commandUtils.Response
	ListCompaniesCommand() *commandUtils.Response
	ListCompanySubscriptionsCommand(req *domain.ListCompanySubscriptionsRequest) *commandUtils.Response
	UpdateCompanyCommand(req *domain.UpdateCompanyRequest) *commandUtils.Response
	DeleteCompanyCommand(req *domain.DeleteCompanyRequest) *commandUtils.Response
}
