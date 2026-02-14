package CompanyService

type CompanyService interface {
	CreateCompany()
	GetCompany()
	ListCompanies()
	ListCompanySubscriptions()
	UpdateCompany()
	DeleteCompany()
}
