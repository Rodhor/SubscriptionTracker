package CompanyCommand

type CompanyCommand interface {
	CreateCompanyCommand()
	GetCompanyCommand()
	ListCompaniesCommand()
	ListCompanySubscriptionsCommand()
	UpdateCompanyCommand()
	DeleteCompanyCommand()
}
