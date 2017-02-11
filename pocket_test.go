package pocket

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const CONSUMER_KEY = "dummy_consumer_key"
const ACCESS_TOKEN = "dummy_access_token"

func TestNewClient1(t *testing.T) {
	c, err := NewClient(CONSUMER_KEY, ACCESS_TOKEN)
	assert.Nil(t, err)
	assert.EqualValues(t, c.ConsumerKey, CONSUMER_KEY)
	assert.EqualValues(t, c.AccessToken, ACCESS_TOKEN)
}

func TestNewClient2(t *testing.T) {
	c, err := NewClient("", ACCESS_TOKEN)
	assert.NotNil(t, err)
	assert.Nil(t, c)
	assert.EqualValues(t, err.Error(), "Missing ConsumerKey")

}

func TestNewClient3(t *testing.T) {
	c, err := NewClient(CONSUMER_KEY, "")
	assert.NotNil(t, err)
	assert.Nil(t, c)
	assert.EqualValues(t, err.Error(), "Missing AccessToken")
}

func TestAdd(t *testing.T) {}

func TestModify(t *testing.T) {}

func TestRetrieve(t *testing.T) {}
