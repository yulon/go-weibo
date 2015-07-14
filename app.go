package weibo

import (
	"net/http"
	"net/url"
	"encoding/json"
	"bytes"
	"mime/multipart"
)

const (
	apiVer = "https://api.weibo.com/2"
)

func MakeApiUrl(apiName string) string {
	return apiVer + "/" + apiName + ".json"
}

type App struct{
	accessToken string
}

func NewApp(accessToken string) *App {
	return &App{accessToken}
}

func (a *App) Get(apiName string, param url.Values, ret interface{}) {
	param.Set("access_token", a.accessToken)
	resp, _ := http.Get(MakeApiUrl(apiName) + "?" + param.Encode())
	d := json.NewDecoder(resp.Body)
	d.Decode(ret)
}

func (a *App) PostForm(apiName string, param url.Values, ret interface{}) {
	param.Set("access_token", a.accessToken)
	resp, _ := http.PostForm(MakeApiUrl(apiName), param)
	d := json.NewDecoder(resp.Body)
	d.Decode(ret)
}

func (a *App) Post(apiName string, data map[string][]byte, ret interface{}) {
	data["access_token"] = []byte(a.accessToken)

	buf := bytes.NewBuffer(make([]byte, 0, 256))
	mw := multipart.NewWriter(buf)

	for n, d := range data {
		p, _ := mw.CreateFormField(n)
		p.Write(d)
	}

	resp, _ := http.Post(MakeApiUrl(apiName), mw.FormDataContentType(), buf)

	mw.Close()

	d := json.NewDecoder(resp.Body)
	d.Decode(ret)
}
