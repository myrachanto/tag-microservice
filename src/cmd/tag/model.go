package tag

import (
	httperrors "github.com/myrachanto/erroring"
	"github.com/myrachanto/microservice/tag/src/support"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tag struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Code        string             `json:"code"`
	Url         string             `json:"url"`
	Used        bool               `json:"used"`
	Picture     string             `json:"picture"`
	Featured    bool               `json:"featured"`
	Base        support.Base
}
type Majorcat struct {
	Code string `json:"code"`
}

func (t Tag) Validate() httperrors.HttpErr {
	if t.Name == "" {
		return httperrors.NewBadRequestError("Invalid Name")
	}
	if t.Title == "" {
		return httperrors.NewBadRequestError("Invalid title")
	}
	if t.Description == "" {
		return httperrors.NewBadRequestError("Invalid Description")
	}
	// if t.Shopalias == "" {
	// 	return httperrors.NewNotFoundError("Invalid Shopalias")
	// }
	return nil
}
