package mailing

// Email is the data structure that represents an email.
type Email struct {
	idCliente *int
	email     string
	verified  bool
}

// NewEmail creates a new email.
func NewEmail(idCliente *int, email string, verified bool) Email {
	return Email{
		idCliente: idCliente,
		email:     email,
		verified:  verified,
	}
}

// IDCliente returns the email's client ID.
func (e Email) IDCliente() *int {
	return e.idCliente
}

// EmailAddress returns the email address.
func (e Email) EmailAddress() string {
	return e.email
}

// IsVerified returns whether the email is verified.
func (e Email) IsVerified() bool {
	return e.verified
}
