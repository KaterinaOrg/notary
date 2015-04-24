package auth

import (
	"errors"
	"net/http"
)

type User struct {
	Name string
}

type Resource struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

// Scope is an identifier scope
type Scope interface {
	ID() string
	Compare(Scope) bool
}

// Authorizer is an interfaces to authorize a scope
type Authorizer interface {
	// Authorize checks whether the credentials provided in the request
	// are sufficient to satisfy the required scopes.
	Authorize(*http.Request, ...Scope) (*User, error)
}

// Authorization is an interface to determine whether
// an object has a scope
type Authorization interface {
	HasScope(Scope) bool
}

// ### START INSECURE AUTHORIZATION TOOLS ###
// THESE ARE FOR DEV PURPOSES ONLY, DO NOT USE IN
// PRODUCTION

// InsecureAuthorizer is an insecure implementation of IAuthorizer.
// WARNING: DON'T USE THIS FOR ANYTHING, IT'S VERY INSECURE
type InsecureAuthorizer struct{}

// Authorize authorizes any scope
// WARNING: LIKE I SAID, VERY INSECURE
func (auth *InsecureAuthorizer) Authorize(r *http.Request, scopes ...Scope) (*User, error) {
	return &User{}, nil
}

// InsecureAuthorization is an implementation of IAuthorization
// which will consider any scope authorized.
// WARNING: ALSO DON'T USE THIS, IT'S ALSO VERY INSECURE
type InsecureAuthorization struct {
}

// HasScope always returns true for any scope
// WARNING: THIS IS JUST INCREDIBLY INSECURE
func (authzn *InsecureAuthorization) HasScope(scope Scope) bool {
	return true
}

// ### END INSECURE AUTHORIZATION TOOLS ###

// NoAuthorizer is a factory for NoAuthorization object
type NoAuthorizer struct{}

// Authorize implements the IAuthorizer interface
func (auth *NoAuthorizer) Authorize(r *http.Request, scopes ...Scope) (*User, error) {
	return nil, errors.New("User not authorized")
}

// NoAuthorization is an implementation of IAuthorization
// which never allows a scope to be valid.
type NoAuthorization struct{}

// HasScope returns false for any scope
func (authzn *NoAuthorization) HasScope(scope Scope) bool {
	return false
}

// SimpleScope is a simple scope represented by a string.
type SimpleScope string

// ID returns the string representing the scope.
func (ss SimpleScope) ID() string {
	return string(ss)
}

// TODO: possibly rename to Equals
// Compare compares to the given scope for equality.
func (ss SimpleScope) Compare(toCompare Scope) bool {
	return ss.ID() == toCompare.ID()
}

var _ Scope = SimpleScope("Test")

const (
	// SSNoAuth is the simple scope "NoAuth"
	SSNoAuth SimpleScope = SimpleScope("NoAuth")

	// SSCreate is the simple scope "Create"
	SSCreate = SimpleScope("Create")

	// SSRead is the simple scope "Read"
	SSRead = SimpleScope("Read")

	// SSUpdate is the simple scope "Update"
	SSUpdate = SimpleScope("Update")

	// SSDelete is the simple scope "Delete"
	SSDelete = SimpleScope("Delete")
)