package plaid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInfo(t *testing.T) {

	test_id := "test_id"
	test_secret := "test_secret"
	environment := "https://tartan.plaid.com"

	accessToken := "test_chase"

	c := NewClient(test_id, test_secret, environment)

	t.Log("Expected PostResponse")
	postRes, mfaRes, err := c.InfoGet(accessToken)

	assert.Nil(t, mfaRes)
	assert.Nil(t, err)
	assert.NotNil(t, postRes)
	assert.Equal(t, len(postRes.Accounts) > 0, true)
	assert.NotNil(t, postRes.Info)
}
