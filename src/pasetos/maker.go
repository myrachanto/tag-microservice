package pasetos

import (
	"time"

	httperrors "github.com/myrachanto/erroring"
)

type Maker interface {
	//creates a new token
	CreateToken(data *Data, duration time.Duration) (string, httperrors.HttpErr)
	VerifyToken(token string) (*Payload, httperrors.HttpErr)
}
