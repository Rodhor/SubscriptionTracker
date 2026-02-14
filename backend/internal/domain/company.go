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
	Name      string      `json:"name"`
	Area      ProductArea `json:"area"`
	TANNSSID  string      `json:"tannssid"`
	Comments  string      `json:"comments"`
	IsActive  bool        `json:"is_active"`
	DeletedAt *time.Time  `json:"deleted_at"`
}

// ============================================================================
// 2. The Database Entity
// ============================================================================

// Company represents the full database row.
type Company struct {
	ID          uuid.UUID `json:"id"`
	CompanyBase CompanyBase
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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
	ID          uuid.UUID `json:"id"`
	CompanyBase CompanyBase
	CreatedAt   time.Time `json:"created_at"`
}

// CreateCompanyRequest is the payload for adding a new company.
type CreateCompanyRequest struct {
	Name     string      `json:"name,omitempty"`
	Area     ProductArea `json:"area,omitempty"`
	License  string      `json:"license,omitempty"`
	IsActive *bool       `json:"is_active,omitempty"`
	Comments *string     `json:"comments,omitempty"`
}

// UpdateCompanyRequest allows partial updates.
type UpdateCompanyRequest struct {
	ID       uuid.UUID    `json:"id"`
	Name     *string      `json:"name,omitempty"`
	TANSSID  *string      `json:"tanssid,omitempty"`
	Area     *ProductArea `json:"area,omitempty"`
	License  *string      `json:"license,omitempty"`
	IsActive *bool        `json:"is_active,omitempty"`
	Comments *string      `json:"comments,omitempty"`
}

// DeleteCompanyRequest for soft deleting a company.
type DeleteCompanyRequest struct {
	ID uuid.UUID `json:"id"`
}
