package customer

import app "github.com/isaqueveraslabs/email-microservice/application/email"

// Server implements proto interface
type Server struct {
	app.UnimplementedEMailServer
}
