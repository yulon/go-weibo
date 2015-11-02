package weibo

import (
	"net/http"
	"net/url"
	"encoding/json"
	"bytes"
	"mime/multipart"
	"io"
)

const (
	apiVer = "https://api.weibo.com/2"
)

func MakeApiUrl(apiName string) string {
	return apiVer + "/" + apiName + ".json"
}

type Client struct{
	accessToken string
}

func NewClient(accessToken string) *Client {
	return &Client{accessToken}
}

func (c *Client) Get(apiName string, params url.Values, ret interface{}) {
	params.Set("access_token", c.accessToken)
	resp, _ := http.Get(MakeApiUrl(apiName) + "?" + params.Encode())
	d := json.NewDecoder(resp.Body)
	d.Decode(ret)
}

func (c *Client) PostForm(apiName string, params url.Values, ret interface{}) {
	params.Set("access_token", c.accessToken)
	resp, _ := http.PostForm(MakeApiUrl(apiName), params)
	d := json.NewDecoder(resp.Body)
	d.Decode(ret)
}

type ReadNamer interface{
	io.Reader
	Name()string
}

func (c *Client) PostFormData(apiName string, params url.Values, files map[string]ReadNamer, ret interface{}) {
	params.Set("access_token", c.accessToken)

	buf := bytes.NewBuffer(make([]byte, 0, 256))
	mw := multipart.NewWriter(buf)

	for n, _ := range params {
		mw.WriteField(n, params.Get(n))
	}

	for n, f := range files {
		w, _ := mw.CreateFormFile(n, f.Name())
		io.Copy(w, f)
	}

	resp, _ := http.Post(MakeApiUrl(apiName), mw.FormDataContentType(), buf)

	mw.Close()

	d := json.NewDecoder(resp.Body)
	d.Decode(ret)
}

type TokenInfo struct{
	Uid int64 `json:"uid"`
	Appkey string `json:"appkey"`
	CreateAt int64 `json:"create_at"`
	ExpiresIn int64 `json:"expires_in"`
}

func (c *Client) GetTokenInfo() (ti *TokenInfo) {
	resp, err := http.PostForm("https://api.weibo.com/oauth2/get_token_info", url.Values{
		"access_token": {c.accessToken},
	})
	if err != nil {
		return
	}
	ti = &TokenInfo{}
	d := json.NewDecoder(resp.Body)
	d.Decode(ti)
	return
}

func (c *Client) StatusesService() *StatusesService {
	return &StatusesService{c}
}

func (c *Client) CommentsService() *CommentsService {
	return &CommentsService{c}
}

func (c *Client) UsersService() *UsersService {
	return &UsersService{c}
}

func (c *Client) SearchService() *SearchService {
	return &SearchService{c}
}

func (c *Client) RemindService() *RemindService {
	return &RemindService{c}
}
