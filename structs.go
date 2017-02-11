package pocket

import (
	"net/http"
)

type Client struct {
	HTTPClient *http.Client

	Host        string
	ApiEndpoint string

	ConsumerKey, AccessToken string
}

type OauthInput struct {
	ConsumerKey string `json:"consumer_key"`
	RedirectURI string `json:"redirect_uri"`
}

type AuthorizeInput struct {
	ConsumerKey string `json:"consumer_key"`
	Code        string `json:"code"`
}
type OauthOutput struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

type AuthorizeOutput struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
}

type RetrieveInput struct {
	ConsumerKey string `json:"consumer_key"`
	AccessToken string `json:"access_token"`
	RetrieveOpts
}

type RetrieveOpts struct {
	State       string `json:"state,omitempty"`
	Favorite    bool   `json:"favorite,omitempty"`
	Tag         string `json:"tag,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	Sort        string `json:"sort,omitempty"`
	DetailType  string `json:"detailType,omitempty"`
	Search      string `json:"search,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Since       string `json:"since,omitempty"`
	Count       int    `json:"count,omitempty"`
	Offset      int    `json:"offset,omitempty"`
}

type RetrieveOutput struct {
	Status int `json:status`
	List   map[string]Item
}

type Item struct {
	ItemID        string `json:"item_id"`
	ResolvedID    string `json:"resolved_id"`
	GivenURL      string `json:"given_url"`
	ResolvedURL   string `json:"resolved_url"`
	GivenTitle    string `json:"given_title"`
	ResolvedTitle string `json:"resolved_title"`
	Favorite      string `json:"favorite"`
	Status        string `json:"status"`
	Excerpt       string `json:"excerpt"`
	IsArticle     string `json:"is_article"`
	HasImage      string `json:"has_image"`
	HasVideo      string `json:"has_video"`
	WordCount     string `json:"word_count"`
	Images        map[string]Image
	Videos        map[string]Video
}

type Image struct {
	ItemID  string `json:"item_id"`
	ImageID string `json:"image_id"`
	Src     string `json:"src"`
	Width   string `json:"width"`
	Height  string `json:"height"`
	Credit  string `json:"credit"`
	Caption string `json:"caption"`
}

type Video struct {
	ItemID  string `json:"item_id"`
	VideoID string `json:"video_id"`
	Src     string `json:"src"`
	Width   string `json:"width"`
	Height  string `json:"height"`
	Type    string `json:"type"`
	Vid     string `json:"vid"`
}

type AddInput struct {
	ConsumerKey string `json:"consumer_key"`
	AccessToken string `json:"access_token"`
	URL         string `json:"url"`
	AddOpts
}

type AddOpts struct {
	Title   string `json:"title,omitempty"`
	Tags    string `json:"tags,omitempty"`
	TweetID int    `json:"tweet_id,omitempty"`
}

type AddOutput struct {
	Status int `json:status`
}

type ModifyInput struct {
	ConsumerKey string    `json:"consumer_key"`
	AccessToken string    `json:"access_token"`
	Actions     []*Action `json:"actions"`
}

type ModifyOutput struct {
	Status        int      `json:"status"`
	ActionResults []string `json"action_results"`
}

type Action struct {
	Action string `json:"action"`
	ItemID string `json:"item_id"`
}
