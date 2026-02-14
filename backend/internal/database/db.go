package database

import (
	"Groundwork/backend/internal/domain"
	"context"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Database struct {
	mu            sync.RWMutex
	companies     map[uuid.UUID]*domain.Company
	subscriptions map[uuid.UUID]*domain.Subscription
}

var ErrDuplicateCompany = errors.New("TANSSID already exists in database")
var ErrNotFound = errors.New("not found")

func NewDB() *Database {
	return &Database{
		companies:     make(map[uuid.UUID]*domain.Company),
		subscriptions: make(map[uuid.UUID]*domain.Subscription),
	}
}

// ============================================================================
// Company functionality
// ============================================================================

// Creates a new company in the database returning an error if the TANSSID already exists.
func (db *Database) CreateCompany(ctx context.Context, company *domain.Company) error {
	// Check for context cancellation
	if err := ctx.Err(); err != nil {
		return err
	}

	// proceed with the creation of the company
	db.mu.Lock()
	defer db.mu.Unlock()
	for _, c := range db.companies {
		if c.CompanyBase.TANSSID == company.CompanyBase.TANSSID {
			return ErrDuplicateCompany
		}
	}

	db.companies[company.ID] = company
	return nil
}

// Softdeletes a company in the database returning an error if the company does not exist.
func (db *Database) DeleteCompany(ctx context.Context, ID uuid.UUID, deletionTime time.Time) error {
	// Check for context cancellation
	if err := ctx.Err(); err != nil {
		return err
	}

	// proceed with the deletion of the company
	db.mu.Lock()
	defer db.mu.Unlock()
	company, ok := db.companies[ID]
	if !ok {
		return ErrNotFound
	}

	company.DeletedAt = &deletionTime
	company.UpdatedAt = deletionTime

	return nil
}

// Retrieves a company by its ID from the database returning an error if the company does not exist.
func (db *Database) GetCompanyByID(ctx context.Context, ID uuid.UUID) (*domain.Company, error) {
	// Check for context cancellation
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// proceed with the retrieval of the company
	db.mu.RLock()
	defer db.mu.RUnlock()
	company, ok := db.companies[ID]
	if !ok || company.DeletedAt != nil {
		return nil, ErrNotFound
	}

	return company, nil
}

// Retrieves a list of companies from the database.
func (db *Database) ListCompanies(ctx context.Context) ([]*domain.Company, error) {
	// Check for context cancellation
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// proceed with the retrieval of the companies
	db.mu.RLock()
	defer db.mu.RUnlock()
	companies := make([]*domain.Company, 0, len(db.companies))
	for _, company := range db.companies {
		if company.DeletedAt == nil {
			companies = append(companies, company)
		}
	}

	return companies, nil
}

// Retrieve a list of subscriptions for a company.
func (db *Database) ListCompanySubscriptions(ctx context.Context, ID uuid.UUID) ([]*domain.Subscription, error) {
	// Check for context cancellation
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// proceed with the retrieval of the company's subscriptions
	db.mu.RLock()
	defer db.mu.RUnlock()
	subscriptions := []*domain.Subscription{}
	for _, subscription := range db.subscriptions {
		if subscription.CompanyID == ID && subscription.DeletedAt == nil {
			subscriptions = append(subscriptions, subscription)
		}
	}

	return subscriptions, nil
}

// Update a company's information - Upon changing the TANSSID check if it's already in use by another company
func (db *Database) UpdateCompany(ctx context.Context, req *domain.UpdateCompanyRequest, updatingTime time.Time) error {
	// Check for context cancellation
	if err := ctx.Err(); err != nil {
		return err
	}

	// proceed with the update of the company
	db.mu.Lock()
	defer db.mu.Unlock()
	company, ok := db.companies[req.ID]
	if !ok || company.DeletedAt != nil {
		return ErrNotFound
	}

	// If TANSSID is provided, check if it's already in use by another company
	if req.TANSSID != nil {
		for _, c := range db.companies {
			if c.CompanyBase.TANSSID == *req.TANSSID && c.ID != req.ID {
				return ErrDuplicateCompany
			}
		}
		company.CompanyBase.TANSSID = *req.TANSSID
	}

	if req.Name != nil {
		company.CompanyBase.Name = *req.Name
	}
	if req.Area != nil {
		company.CompanyBase.Area = *req.Area
	}
	if req.Comments != nil {
		company.CompanyBase.Comments = *req.Comments
	}
	if req.IsActive != nil {
		company.CompanyBase.IsActive = *req.IsActive
	}

	company.UpdatedAt = updatingTime

	return nil
}

// ============================================================================
// Subscription functionality
// ============================================================================

func (db *Database) CreateSubscription() {
	panic("Not implemented")
}

func (db *Database) UpdateSubscription() {
	panic("Not implemented")
}

func (db *Database) DeleteSubscription() {
	panic("Not implemented")
}

func (db *Database) GetSubscription() {
	panic("Not implemented")
}
func (db *Database) ListSubscriptions() {
	panic("Not implemented")
}
