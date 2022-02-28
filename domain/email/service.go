package email

import "github.com/isaqueveraslabs/email-microservice/configuration/database"

// Service models a service base struct
type Service struct {
	repo IEmail
}

// GetService retrieves a service type
func GetService(r IEmail) *Service {
	return &Service{repo: r}
}

// GetEmailRepository retrieve repository for access to email data
func GetEmailRepository(db *database.DBTransaction) IEmail {
	return New(db)
}
