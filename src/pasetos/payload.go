package pasetos

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	httperrors "github.com/myrachanto/erroring"
)

var (
	ErrExpiredToken = "token has expired"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}
type Data struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewPayload(data *Data, duration time.Duration) (*Payload, httperrors.HttpErr) {
	tokenid, err := uuid.NewRandom()
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("error with uuid generation, %d", err))
	}
	return &Payload{
		ID:        tokenid,
		Username:  data.Username,
		Email:     data.Email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}, nil
}
func (payload *Payload) Valid() httperrors.HttpErr {
	if time.Now().After(payload.ExpiredAt) {
		return httperrors.NewBadRequestError(ErrExpiredToken)
	}
	return nil
}
