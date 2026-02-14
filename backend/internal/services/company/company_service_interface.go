package CompanyService

import (
	"Groundwork/backend/internal/domain"
	"context"
)

type CompanyService interface {
	CreateCompanyService(ctx context.Context, req *domain.CreateCompanyRequest) (*domain.CompanyResponse, error)
	GetCompanyService(ctx context.Context, req *domain.GetCompanyRequest) (*domain.CompanyResponse, error)
	ListCompaniesService(ctx context.Context) ([]*domain.CompanyResponse, error)
	ListCompanySubscriptionsService(ctx context.Context, req *domain.ListCompanySubscriptionsRequest) ([]*domain.SubscriptionResponse, error)
	UpdateCompanyService(ctx context.Context, req *domain.UpdateCompanyRequest) error
	DeleteCompanyService(ctx context.Context, req *domain.DeleteCompanyRequest) error
}
