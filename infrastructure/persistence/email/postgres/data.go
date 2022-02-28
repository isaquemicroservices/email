package postgres

import "github.com/isaqueveraslabs/email-microservice/configuration/database"

// PGEmail implements methods for postgres query execution
type PGEmail struct {
	DB *database.DBTransaction
}
