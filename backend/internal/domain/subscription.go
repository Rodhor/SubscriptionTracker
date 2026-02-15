package domain

import (
	"time"

	"github.com/google/uuid"
)

// ============================================================================
// 1. SubscriptionBase (Shared Data)
// ============================================================================

type SubscriptionTier string

const (
	TierCore          SubscriptionTier = "Core"
	TierEnterprise    SubscriptionTier = "Enterprise"
	TierEnterpriseAPI SubscriptionTier = "Enterprise API"
)

type SubscriptionStatus string

const (
	StatusActive    SubscriptionStatus = "Active"
	StatusInactive  SubscriptionStatus = "Inactive"
	StatusCancelled SubscriptionStatus = "Cancelled"
	StatusPending   SubscriptionStatus = "Pending"
	StatusExpired   SubscriptionStatus = "Expired"
)

// subscriptionBase contains the core details of the plan.
type SubscriptionBase struct {
	Name      string             `json:"name"`
	License   string             `json:"license"`
	Tier      SubscriptionTier   `json:"tier"`
	Features  []string           `json:"features"`
	Status    SubscriptionStatus `json:"status"`
	StartDate time.Time          `json:"start_date"`
	EndDate   *time.Time         `json:"end_date"`
	AutoRenew bool               `json:"auto_renew"`
}

// ============================================================================
// 2. The Database Entity
// ============================================================================

type Subscription struct {
	ID               uuid.UUID `json:"id"`
	CompanyID        uuid.UUID `json:"company_id"`
	SubscriptionBase SubscriptionBase
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
}

func (s *Subscription) Sanitize() *SubscriptionResponse {
	return &SubscriptionResponse{
		ID:               s.ID,
		CompanyID:        s.CompanyID,
		SubscriptionBase: s.SubscriptionBase,
	}
}

// ============================================================================
// 3. Data Transfer Objects (DTOs)
// ============================================================================

// SubscriptionResponse for representing a subscription to the frontend.
type SubscriptionResponse struct {
	ID               uuid.UUID `json:"id"`
	CompanyID        uuid.UUID `json:"company_id"`
	SubscriptionBase SubscriptionBase
}

// CreateSubscriptionRequest for creating a new subscription.
type CreateSubscriptionRequest struct {
	CompanyID uuid.UUID           `json:"company_id"`
	Name      string              `json:"name"`
	License   string              `json:"license"`
	Tier      *SubscriptionTier   `json:"tier,omitempty"`
	Features  *[]string           `json:"features,omitempty"`
	Status    *SubscriptionStatus `json:"status,omitempty"`
	StartDate *time.Time          `json:"start_date,omitempty"`
	EndDate   *time.Time          `json:"end_date,omitempty"`
	AutoRenew *bool               `json:"auto_renew,omitempty"`
}

// UpdateSubscriptionRequest for updating a subscription.
type UpdateSubscriptionRequest struct {
	ID        uuid.UUID           `json:"id"`
	Name      *string             `json:"name,omitempty"`
	License   *string             `json:"license,omitempty"`
	Tier      *SubscriptionTier   `json:"tier,omitempty"`
	Features  *[]string           `json:"features,omitempty"`
	Status    *SubscriptionStatus `json:"status,omitempty"`
	StartDate *time.Time          `json:"start,omitempty"`
	EndDate   *time.Time          `json:"end_date,omitempty"`
	AutoRenew *bool               `json:"auto_renew,omitempty"`
}

// DeleteSubscriptionRequest for soft deleting a subscription.
type DeleteSubscriptionRequest struct {
	ID uuid.UUID `json:"id"`
}

// GetSubscriptionRequest for retrieving a singular subscription
type GetSubscriptionRequest struct {
	ID uuid.UUID `json:"id"`
}
