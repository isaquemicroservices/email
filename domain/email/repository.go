package email

import (
	"github.com/isaqueveraslabs/email-microservice/configuration/database"
	"github.com/isaqueveraslabs/email-microservice/infrastructure/persistence/email/postgres"
)

// repository is a base structure that implements methods specified by IEmail
type repository struct {
	pg *postgres.PGEmail
}

// New creates a new customer repository
func New(db *database.DBTransaction) *repository {
	return &repository{pg: &postgres.PGEmail{DB: db}}
}
