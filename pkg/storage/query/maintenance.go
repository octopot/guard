package query

import (
	"time"

	domain "github.com/kamilsk/guard/pkg/service/types"
)

// RegisterAccount TODO issue#docs
type RegisterAccount struct {
	ID   *domain.ID
	Name string

	users []*RegisterUser
}

// AddUser TODO issue#docs
func (q *RegisterAccount) AddUser(user *RegisterUser) *RegisterAccount {
	if user != nil {
		q.users = append(q.users, user)
	}
	return q
}

// Users TODO issue#docs
func (q *RegisterAccount) Users() []*RegisterUser {
	return q.users
}

// RegisterUser TODO issue#docs
type RegisterUser struct {
	ID   *domain.ID
	Name string

	tokens []*RegisterToken
}

// AddToken TODO issue#docs
func (q *RegisterUser) AddToken(token *RegisterToken) *RegisterUser {
	if token != nil {
		q.tokens = append(q.tokens, token)
	}
	return q
}

// Tokens TODO issue#docs
func (q *RegisterUser) Tokens() []*RegisterToken {
	return q.tokens
}

// RegisterToken TODO issue#docs
type RegisterToken struct {
	ID        *domain.Token
	ExpiredAt *time.Time
}
