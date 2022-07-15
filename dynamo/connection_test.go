package dynamo_test

import (
	"testing"

	"github.com/crud/dynamo"
	"github.com/stretchr/testify/assert"
)

func init() {
	dynamo.Dynamo = dynamo.ConnectDynamo()
}

func TestCreateTablePeople(t *testing.T) {
	err := dynamo.CreateTablePeople()

	assert.Nil(t, err)
}

func TestCreateTableHistory(t *testing.T) {
	err := dynamo.CreateTableHistory()

	assert.Nil(t, err)
}

func TestInjectItem(t *testing.T) {
	err := dynamo.CreateItem(dynamo.User)

	assert.Nil(t, err)
}

// func TestCreateAndRead(t *testing.T) {
// 	DB.ResetTables(UserProfile{})

// 	azer := UserProfile{
// 		Name:  "Azer",
// 		Bio:   "I like photography",
// 		Email: "azer@roadbeats.com",
// 	}

// 	assert.Equal(t, azer.Id, 0)
// 	err := DB.CreateAndRead(&azer)
// 	assert.Nil(t, err)
// 	assert.Equal(t, azer.Id, 1)

// 	DB.DropTables(UserProfile{})
// }

// func TestCreateEmpty(t *testing.T) {
// 	DB.ResetTables(UserProfile{})

// 	azer := UserProfile{
// 		Name: "Azer",
// 	}

// 	err := DB.Create(azer)
// 	assert.Nil(t, err)

// 	DB.DropTables(UserProfile{})
// }

// func TestEmbedding(t *testing.T) {
// 	DB.ResetTables(EmbeddedFoo{})

// 	foo := EmbeddedFoo{
// 		Foo: Foo{
// 			APIKey: "hi",
// 			YOLO:   true,
// 			Beast:  "span eggs",
// 		},
// 		Span: 123,
// 		Eggs: "yoo",
// 	}

// 	assert.Equal(t, foo.Id, 0)
// 	err := DB.CreateAndRead(&foo)
// 	assert.Nil(t, err)
// 	assert.Equal(t, foo.Id, 1)
// 	assert.Equal(t, foo.APIKey, "hi")
// 	assert.Equal(t, foo.Beast, "span eggs")

// 	DB.DropTables(EmbeddedFoo{})
// }

// func TestCreatingRenamedTableRow(t *testing.T) {
// 	DB.ResetTables(Post{})

// 	p := Post{
// 		Title: "Foo",
// 		Text:  "bar",
// 	}

// 	assert.Equal(t, p.Id, 0)
// 	err := DB.CreateAndRead(&p)
// 	assert.Nil(t, err)
// 	assert.Equal(t, p.Id, 1)

// 	DB.DropTables(Post{})
// }
