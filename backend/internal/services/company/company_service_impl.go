package CompanyService

import (
	"Groundwork/backend/internal/database"
	"Groundwork/backend/internal/domain"
	"context"
	"time"

	"github.com/google/uuid"
)

type companyService struct {
	queries *database.Database
}

func NewCompanyService(db *database.Database) CompanyService {
	return &companyService{
		queries: db,
	}
}

// CreateCompany implements [CompanyService].
func (c *companyService) CreateCompanyService(ctx context.Context, req *domain.CreateCompanyRequest) (*domain.CompanyResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	now := time.Now()

	// Construct the CompanyBase with defaults
	newCompanyBase := domain.CompanyBase{
		Name:     req.Name,
		Area:     req.Area,
		TANSSID:  req.TANSSID,
		IsActive: true,       // Default
		Comments: []string{}, // Default
	}

	// Override defaults if provided in the request
	if req.IsActive != nil {
		newCompanyBase.IsActive = *req.IsActive
	}
	if req.Comments != nil {
		newCompanyBase.Comments = *req.Comments
	}

	// Assemble the Domain Entity
	newCompany := &domain.Company{
		ID:          uuid.New(),
		CompanyBase: newCompanyBase,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := c.queries.CreateCompany(ctx, newCompany); err != nil {
		return nil, err
	}

	return newCompany.Sanitize(), nil
}

// DeleteCompany implements [CompanyService].
func (c *companyService) DeleteCompanyService(ctx context.Context, req *domain.DeleteCompanyRequest) error {
	// Check context for cancellation
	if err := ctx.Err(); err != nil {
		return err
	}
	// get timestamp
	now := time.Now()
	if err := c.queries.DeleteCompany(ctx, req.ID, now); err != nil {
		return err
	}
	// If deletion was successful, return nil
	return nil
}

// GetCompany implements [CompanyService].
func (c *companyService) GetCompanyService(ctx context.Context, req *domain.GetCompanyRequest) (*domain.CompanyResponse, error) {
	// Check context for cancellation
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	// try to retrieve company from database returning early on error
	company, err := c.queries.GetCompanyByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	// if successful, sanitize and return company
	return company.Sanitize(), nil
}

// ListCompanies implements [CompanyService].
func (c *companyService) ListCompaniesService(ctx context.Context) ([]*domain.CompanyResponse, error) {
	// Check context for cancellation
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// try to retrieve companies from database returning early on error
	companies, err := c.queries.ListCompanies(ctx)
	if err != nil {
		return nil, err
	}

	// if successful, sanitize and return companies
	sanitizedCompanies := make([]*domain.CompanyResponse, 0, len(companies))
	for _, company := range companies {
		sanitizedCompanies = append(sanitizedCompanies, company.Sanitize())
	}
	return sanitizedCompanies, nil
}

// ListCompanySubscriptions implements [CompanyService].
func (c *companyService) ListCompanySubscriptionsService(ctx context.Context, req *domain.ListCompanySubscriptionsRequest) ([]*domain.SubscriptionResponse, error) {
	// Check context for cancellation
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// try to retrieve subscriptions from database returning early on error
	subscriptions, err := c.queries.ListCompanySubscriptions(ctx, req.CompanyID)
	if err != nil {
		return nil, err
	}

	// if successful, sanitize and return subscriptions
	sanitizedSubscriptions := make([]*domain.SubscriptionResponse, 0, len(subscriptions))
	for _, subscription := range subscriptions {
		sanitizedSubscriptions = append(sanitizedSubscriptions, subscription.Sanitize())
	}
	return sanitizedSubscriptions, nil
}

// UpdateCompany implements [CompanyService].
func (c *companyService) UpdateCompanyService(ctx context.Context, req *domain.UpdateCompanyRequest) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	now := time.Now()
	if err := c.queries.UpdateCompany(ctx, req, now); err != nil {
		return err
	}

	return nil
}
