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
		user.account, q.users = q, append(q.users, user)
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

	account *RegisterAccount
	tokens  []*RegisterToken
}

// AddToken TODO issue#docs
func (q *RegisterUser) AddToken(token *RegisterToken) *RegisterUser {
	if token != nil {
		token.user, q.tokens = q, append(q.tokens, token)
	}
	return q
}

// Account TODO issue#docs
func (q *RegisterUser) Account() *RegisterAccount {
	return q.account
}

// Tokens TODO issue#docs
func (q *RegisterUser) Tokens() []*RegisterToken {
	return q.tokens
}

// WithDefaultToken TODO issue#docs
func (q *RegisterUser) WithDefaultToken() *RegisterUser {
	return q.AddToken(&RegisterToken{})
}

// RegisterToken TODO issue#docs
type RegisterToken struct {
	ID        *domain.Token
	ExpiredAt *time.Time

	user *RegisterUser
}

// User TODO issue#docs
func (q *RegisterToken) User() *RegisterUser {
	return q.user
}
