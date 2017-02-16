package pocket

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const CONSUMER_KEY = "dummy_consumer_key"
const ACCESS_TOKEN = "dummy_access_token"

type dummyDoer struct{}

func (d *dummyDoer) doRequest(req *http.Request) (string, error) {
	return `
{"status":1,"list":{"1234567890":{"item_id":"1234567890",
"resolved_id":"1234567890",
"given_url":"http://example.com/",
"given_title":"dummy title",
"favorite":"0",
"status":"0",
"resolved_title":"dummy title",
"resolved_url":"http://example.com/",
"excerpt":"dummy",
"is_article":"1",
"has_video":"1",
"has_image":"1",
"word_count":"3197",
"images":{"1":{"item_id":"1234567890","image_id":"1",
	"src":"http://example.com/dummy.jpg",
	"width":"0","height":"0","credit":"dummy","caption":""}},
"videos":{"1":{"item_id":"1234567890","video_id":"1",
	"src":"http://example.com/dymmy",
	"width":"420","height":"315","type":"1","vid":"abcdefghijk"}}}}}
`, nil
}

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

func TestRetrieve(t *testing.T) {
	client := &Client{
		Host:        "https://getpocket.com",
		ApiEndpoint: "/v3",
		ConsumerKey: CONSUMER_KEY,
		AccessToken: ACCESS_TOKEN,
		Pocketer:    new(dummyDoer),
	}

	items, err := client.Retrieve(&RetrieveOpts{})
	assert.Nil(t, err)
	assert.NotNil(t, items)
	for _, item := range items.List {
		assert.EqualValues(t, item.ResolvedTitle, "dummy title")
		assert.EqualValues(t, item.GivenURL, "http://example.com/")
		assert.EqualValues(t, item.ItemID, "1234567890")
	}
}
