package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbstractPagination(t *testing.T) {

	pg := &Pagination{}

	assert.NotNil(t, pg)
	assert.Empty(t, pg.Size)
	assert.Empty(t, pg.Limit)
	assert.Empty(t, pg.IsLastPage)
	assert.Empty(t, pg.NextPageStart)
}

func TestPrPagination(t *testing.T) {

	prPage := &PRPagination{}

	assert.NotNil(t, prPage)
	assert.Empty(t, prPage.Size)
	assert.Empty(t, prPage.Limit)
	assert.Empty(t, prPage.IsLastPage)
	assert.Empty(t, prPage.NextPageStart)

	assert.Empty(t, prPage.Values)
}

func TestUserPagination(t *testing.T) {

	urPage := &UserPagination{}

	assert.NotNil(t, urPage)
	assert.Empty(t, urPage.Size)
	assert.Empty(t, urPage.Limit)
	assert.Empty(t, urPage.IsLastPage)
	assert.Empty(t, urPage.NextPageStart)

	assert.Empty(t, urPage.Values)
	assert.Empty(t, urPage.GetUsers())

}

func TestUserValues(t *testing.T) {
	uv := &UserValues{}

	assert.NotNil(t, uv)
	assert.Empty(t, uv.User)
	assert.Empty(t, uv.User.Name)
	assert.Empty(t, uv.User.DisplayName)
}

func TestUserPagination_GetUsers_ExpectGiven1UserAsResult(t *testing.T) {
	pg := &UserPagination{}
	pg.Size = 10
	pg.Limit = 15
	pg.IsLastPage = true
	pg.NextPageStart = 0

	userValues := []UserValues{{User{Name: "atiba", DisplayName: "Atiba Hutchinson"}}}
	pg.Values = userValues

	assert.NotEmpty(t, pg.GetUsers())
	assert.Equal(t, 1, len(pg.GetUsers()))
	assert.Equal(t, "atiba", pg.GetUsers()[0].Name)
	assert.Equal(t, "Atiba Hutchinson", pg.GetUsers()[0].DisplayName)
}

func TestUserPagination_GetUsers_ExpectGiven3UserAsResult(t *testing.T) {
	pg := &UserPagination{}
	pg.Size = 10
	pg.Limit = 15
	pg.IsLastPage = true
	pg.NextPageStart = 0

	usr1 := User{Name: "atiba", DisplayName: "Atiba Hutchinson"}
	usr2 := User{Name: "dorukhan", DisplayName: "Dorukhan Toküz"}
	usr3 := User{Name: "ersin", DisplayName: "Destanoğlu"}

	userValues := []UserValues{{usr1}, {usr2}, {usr3}}
	pg.Values = userValues

	assert.NotEmpty(t, pg.GetUsers())
	assert.Equal(t, 3, len(pg.GetUsers()))
}
