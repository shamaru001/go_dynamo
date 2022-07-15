package dynamo_test

import (
	"testing"

	"github.com/crud/dynamo"
	"github.com/stretchr/testify/assert"
)

func TestCreateItem(t *testing.T) {
	err := dynamo.CreateItem(dynamo.User)

	assert.Nil(t, err)
}

func TestUpdateItem(t *testing.T) {

	dynamo.User.Name = "Testing"
	err := dynamo.UpdateItem(dynamo.User)

	assert.Nil(t, err)
}

func TestGetItem(t *testing.T) {

	user, err := dynamo.GetItem(1)

	assert.Nil(t, err)
	assert.Equal(t, user.Id, 1)
}
