package httpsession

import (
	"github.com/google/uuid"
	"net/http"
)

// createUUID generates a random UUID and returns UUID and error back to the calling function
func CreateUUID() (uuid.UUID, error) {
	uid := uuid.UUID{}
	uid, err := uuid.NewRandom()
	if err != nil {
		return uid, err
	}
	return uid, nil
}

// CreateCookie creates an HTTP cookie and returns pointer to HTTP cookie and error
func CreateCookie(name string, value string, httpOnly bool, secure bool, path string, domain string) *http.Cookie {
	// Initialize http.Cookie
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		HttpOnly: httpOnly,
		Secure:   secure,
		Path:     path,
		Domain:   domain,
	}
	return &cookie
}
