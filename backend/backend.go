package Backend

import (
	CompanyCommand "Groundwork/backend/internal/commands/company"
	SubscriptionCommand "Groundwork/backend/internal/commands/subcription"
	"Groundwork/backend/internal/database"
	CompanyService "Groundwork/backend/internal/services/company"
	SubscriptionService "Groundwork/backend/internal/services/subscription"
)

type Backend struct {
	Company      CompanyCommand.CompanyCommand
	Subscription SubscriptionCommand.SubscriptionCommand
}

func NewBackend() *Backend {
	// 1. Initialize Database
	db := database.NewDB()

	// 2. Initialize services
	compSvc := CompanyService.NewCompanyService(db)
	subSvc := SubscriptionService.NewSubscriptionService(db)

	// 3. Initialize commands
	compCmd := CompanyCommand.NewCompanyCommand(compSvc)
	subCmd := SubscriptionCommand.NewSubscriptionCommand(subSvc)

	// 4. return the assembled backend
	return &Backend{
		Company:      compCmd,
		Subscription: subCmd,
	}
}
