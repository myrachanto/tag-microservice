package tag

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUserInputRequiredFields(t *testing.T) {
	jsondata := `{"name":"Tag name","title":"tag title","description":"desscription","shopalias": "alias"
	}`
	tag := &Tag{}
	if err := json.Unmarshal([]byte(jsondata), &tag); err != nil {
		t.Errorf("failed to unmarshal tag data %v", err.Error())
	}
	// fmt.Println("------------------", tag)
	expected := ""
	if err := tag.Validate(); err != nil {
		expected = "Invalid Name"
		if err.Message() == expected {
			assert.EqualValues(t, "", err.Message(), "Erro valivating Name")
		}
		expected = "Invalid title"
		if err.Message() == expected {
			assert.EqualValues(t, "", err.Message(), "Error validating title")
		}
		expected = "Invalid Description"
		if err.Message() == expected {
			assert.EqualValues(t, "", err.Message(), "Error validating Description")
		}
		expected = "Invalid Shopalias"
		if err.Message() == expected {
			assert.EqualValues(t, "", err.Message(), "Error validating Shopalias")
		}
	}
}
