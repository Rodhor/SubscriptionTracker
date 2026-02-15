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
var ErrDuplicateSubscription = errors.New("SubscriptionLicense already exists in database")
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

	// Cascade delete subscriptions
	for _, subscription := range db.subscriptions {
		if subscription.CompanyID == company.ID {
			if subscription.DeletedAt == nil {
				subscription.DeletedAt = &deletionTime
				subscription.UpdatedAt = deletionTime
			}
		}
	}

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

// Creates a new subscription - errors out if the license already exists
func (db *Database) CreateSubscription(ctx context.Context, sub *domain.Subscription) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	// 1. Verify Company exists and is not deleted
	company, ok := db.companies[sub.CompanyID]
	if !ok || company.DeletedAt != nil {
		return ErrNotFound
	}

	// 2. Uniqueness check for License
	for _, existing := range db.subscriptions {
		if existing.SubscriptionBase.License == sub.SubscriptionBase.License && existing.DeletedAt == nil {
			return ErrDuplicateSubscription
		}
	}

	db.subscriptions[sub.ID] = sub
	return nil
}

// Updates an existing subscription - errors out if the updated license already exists
func (db *Database) UpdateSubscription(ctx context.Context, req *domain.UpdateSubscriptionRequest, updatingTime time.Time) error {
	// Check for context cancellation
	if err := ctx.Err(); err != nil {
		return err
	}

	// proceed with the update of the company
	db.mu.Lock()
	defer db.mu.Unlock()
	subscription, ok := db.subscriptions[req.ID]
	if !ok || subscription.DeletedAt != nil {
		return ErrNotFound
	}

	// If License is provided, check if it's already in use by another subscription
	if req.License != nil {
		for _, s := range db.subscriptions {
			if s.SubscriptionBase.License == *req.License && s.ID != req.ID && s.DeletedAt == nil {
				return ErrDuplicateSubscription
			}
		}
		subscription.SubscriptionBase.License = *req.License
	}

	if req.Name != nil {
		subscription.SubscriptionBase.Name = *req.Name
	}
	if req.Tier != nil {
		subscription.SubscriptionBase.Tier = *req.Tier
	}
	if req.Features != nil {
		subscription.SubscriptionBase.Features = *req.Features
	}
	if req.Status != nil {
		subscription.SubscriptionBase.Status = *req.Status
	}
	if req.StartDate != nil {
		subscription.SubscriptionBase.StartDate = *req.StartDate
	}
	if req.EndDate != nil {
		subscription.SubscriptionBase.EndDate = req.EndDate
	}
	if req.AutoRenew != nil {
		subscription.SubscriptionBase.AutoRenew = *req.AutoRenew
	}
	subscription.UpdatedAt = updatingTime

	return nil
}

// Softdeletes a subscription
func (db *Database) DeleteSubscription(ctx context.Context, ID uuid.UUID, deletionTime time.Time) error {
	// Check for context cancellation
	if err := ctx.Err(); err != nil {
		return err
	}

	// proceed with the deletion of the company
	db.mu.Lock()
	defer db.mu.Unlock()
	subscription, ok := db.subscriptions[ID]
	if !ok {
		return ErrNotFound
	}

	subscription.DeletedAt = &deletionTime
	subscription.UpdatedAt = deletionTime

	return nil
}

// Retrieves a subscription by ID
func (db *Database) GetSubscription(ctx context.Context, ID uuid.UUID) (*domain.Subscription, error) {
	// Check for context cancellation
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// proceed with the retrieval of the company
	db.mu.RLock()
	defer db.mu.RUnlock()
	subscription, ok := db.subscriptions[ID]
	if !ok || subscription.DeletedAt != nil {
		return nil, ErrNotFound
	}

	return subscription, nil
}

// Retrieves all subscriptions
func (db *Database) ListSubscriptions(ctx context.Context) ([]*domain.Subscription, error) {
	// Check for context cancellation
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// proceed with the retrieval of the companies
	db.mu.RLock()
	defer db.mu.RUnlock()
	subscriptions := make([]*domain.Subscription, 0, len(db.subscriptions))
	for _, subscription := range db.subscriptions {
		if subscription.DeletedAt == nil {
			subscriptions = append(subscriptions, subscription)
		}
	}

	return subscriptions, nil
}
