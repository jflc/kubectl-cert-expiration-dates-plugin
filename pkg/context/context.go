package context

import (
	"time"
)

// Certificate : provides information about a certificate
type Certificate struct {
	Context   string
	Cluster   string
	User      string
	ValidFrom time.Time
	ValidTo   time.Time
}

// CertExpirationDatesContext : provides information about all configured certificates
type CertExpirationDatesContext struct {
	Certificates []Certificate
}
