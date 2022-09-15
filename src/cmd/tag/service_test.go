package tag

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var jsondata = `{"name":"name","title":"this title","description":"description"}`

func TestServiceCreatecategory(t *testing.T) {
	// fmt.Println(">>>>>>>>>", Bizname)
	// Bizname = "test"
	tag := &Tag{}
	if err := json.Unmarshal([]byte(jsondata), &tag); err != nil {
		t.Errorf("failed to unmarshal user data %v", err.Error())
	}
	service := NewtagService(NewtagRepo())
	Bizname = "alias4567"
	u, err := service.Create(tag)
	// fmt.Println(">>>>>>>>>", u)
	assert.EqualValues(t, "name", u.Name, "failed to validate create method")
	assert.Nil(t, err)
	service.Delete(u.Code)
}
func TestServiceGetAlltag(t *testing.T) {

	tag := &Tag{}
	if err := json.Unmarshal([]byte(jsondata), &tag); err != nil {
		t.Errorf("failed to unmarshal user data %v", err.Error())
	}
	service := NewtagService(NewtagRepo())
	Bizname = "alias4567"
	_, _ = service.Create(tag)
	_, err := service.GetAll()
	assert.Nil(t, err, "Something went wrong testing with the Getting all method")
	service.Delete(tag.Code)
}
func TestServiceGetOnetag(t *testing.T) {

	tag := &Tag{}
	if err := json.Unmarshal([]byte(jsondata), &tag); err != nil {
		t.Errorf("failed to unmarshal user data %v", err.Error())
	}
	service := NewtagService(NewtagRepo())
	Bizname = "alias4567"
	tag, err := service.Create(tag)
	assert.EqualValues(t, "name", tag.Name, "Something went wrong testing with the Getting all method")
	assert.Nil(t, err, "Something went wrong testing with the Getting all method")
	service.Delete(tag.Code)
}

func TestServiceUpdatetag(t *testing.T) {
	tag := &Tag{}
	if err := json.Unmarshal([]byte(jsondata), &tag); err != nil {
		t.Errorf("failed to unmarshal user data %v", err.Error())
	}
	service := NewtagService(NewtagRepo())
	Bizname = "alias4567"
	tag1, _ := service.Create(tag)
	tag1.Name = "update"
	u, err := service.Update(tag.Code, tag1)
	assert.EqualValues(t, tag1.Name, u.Name, "Something went wrong testing with the Getting one method")
	assert.Nil(t, err)
	service.Delete(tag1.Code)
}
func TestServiceDeletetag(t *testing.T) {

	tag1 := &Tag{}
	if err := json.Unmarshal([]byte(jsondata), &tag1); err != nil {
		t.Errorf("failed to unmarshal user data %v", err.Error())
	}
	service := NewtagService(NewtagRepo())
	Bizname = "alias4567"
	tag, _ := service.Create(tag1)
	res, err := service.Delete(tag.Code)
	expected := "deleted successfully"
	assert.EqualValues(t, expected, res, "Something went wrong testing with the Deleting method")
	assert.Nil(t, err, "Something went wrong testing with the Deleting method")
}
