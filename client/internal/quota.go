package mailing

// Quota is the data structure that represents a quota.
type Quota struct {
	idCliente  int
	month      int
	year       int
	day        int
	sentEmails int
}

// NewQuota creates a new quota.
func NewQuota(idCliente, month, year, day, sentEmails int) Quota {
	return Quota{
		idCliente:  idCliente,
		month:      month,
		year:       year,
		day:        day,
		sentEmails: sentEmails,
	}
}

// IDCliente returns the quota's client ID.
func (q Quota) IDCliente() int {
	return q.idCliente
}

// Month returns the month of the quota.
func (q Quota) Month() int {
	return q.month
}

// Year returns the year of the quota.
func (q Quota) Year() int {
	return q.year
}

// Day returns the day of the quota.
func (q Quota) Day() int {
	return q.day
}

// SentEmails returns the number of sent emails in the quota.
func (q Quota) SentEmails() int {
	return q.sentEmails
}
