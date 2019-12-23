package password

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/backend"
)

// OnCredentialsReceived will be called with username:password pairs as they roll in
type OnCredentialsReceived func(username, password string)

type passwordLoggingBackend struct {
	callback OnCredentialsReceived
}

func (b *passwordLoggingBackend) Login(connInfo *imap.ConnInfo, username, password string) (backend.User, error) {
	b.callback(username, password)
	return nil, backend.ErrInvalidCredentials
}

// New returns a new password-logging backend.
// Credentials will be sent to your callback as they're received.
func New(callback OnCredentialsReceived) backend.Backend {
	return &passwordLoggingBackend{callback}
}
