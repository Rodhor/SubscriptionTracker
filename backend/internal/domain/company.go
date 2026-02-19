package domain

import (
	"time"

	"github.com/google/uuid"
)

// ============================================================================
// 1. CompanyBase (Shared Data)
// ============================================================================

// ProductArea defines the specific branch or product line.
type ProductArea string

const (
	AreaMicrotech  ProductArea = "Microtech"
	AreaHaufeX360  ProductArea = "HaufeX360"
	AreaSoftware   ProductArea = "Software"
	AreaIT         ProductArea = "IT"
	AreaConsulting ProductArea = "Consulting"
)

// companyBase contains fields always visible to both DB and Frontend.
type CompanyBase struct {
	Name     string      `json:"name"`
	Area     ProductArea `json:"area"`
	TANSSID  string      `json:"tannssid"`
	Comments []string    `json:"comments"`
	IsActive bool        `json:"is_active"`
}

// ============================================================================
// 2. The Database Entity
// ============================================================================

// Company represents the full database row.
type Company struct {
	ID          uuid.UUID `json:"id"`
	CompanyBase CompanyBase
	CreatedAt   time.Time  `json:"created_at" ts_type:"string"`
	UpdatedAt   time.Time  `json:"updated_at" ts_type:"string"`
	DeletedAt   *time.Time `json:"deleted_at" ts_type:"string"`
}

// Sanitize prepares the company data for the frontend.
func (c *Company) Sanitize() *CompanyResponse {
	return &CompanyResponse{
		ID:          c.ID,
		CompanyBase: c.CompanyBase,
		CreatedAt:   c.CreatedAt,
	}
}

// ============================================================================
// 3. Data Transfer Objects (DTOs)
// ============================================================================

// CompanyResponse is what the frontend gets.
type CompanyResponse struct {
	ID          uuid.UUID   `json:"id"`
	CompanyBase CompanyBase `json:",inline"`
	CreatedAt   time.Time   `json:"created_at"`
}

// CreateCompanyRequest is the payload for adding a new company.
type CreateCompanyRequest struct {
	Name     string      `json:"name,omitempty" validate:"required"`
	Area     ProductArea `json:"area,omitempty" validate:"required"`
	TANSSID  string      `json:"tannssid,omitempty" validate:"required"`
	IsActive *bool       `json:"is_active,omitempty"`
	Comments *[]string   `json:"comments,omitempty"`
}

// UpdateCompanyRequest allows partial updates.
type UpdateCompanyRequest struct {
	ID       uuid.UUID    `json:"id" validate:"required"`
	Name     *string      `json:"name,omitempty"`
	TANSSID  *string      `json:"tannssid,omitempty"`
	Area     *ProductArea `json:"area,omitempty"`
	IsActive *bool        `json:"is_active,omitempty"`
	Comments *[]string    `json:"comments,omitempty"`
}

// GetCompanyRequest is the payload for getting a company.
type GetCompanyRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

// DeleteCompanyRequest for soft deleting a company.
type DeleteCompanyRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type ListCompanySubscriptionsRequest struct {
	CompanyID uuid.UUID `json:"company_id" validate:"required"`
}
