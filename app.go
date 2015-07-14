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

type App struct{
	accessToken string
}

func NewApp(accessToken string) *App {
	return &App{accessToken}
}

func (a *App) Get(apiName string, params url.Values, ret interface{}) {
	params.Set("access_token", a.accessToken)
	resp, _ := http.Get(MakeApiUrl(apiName) + "?" + params.Encode())
	d := json.NewDecoder(resp.Body)
	d.Decode(ret)
}

func (a *App) PostForm(apiName string, params url.Values, ret interface{}) {
	params.Set("access_token", a.accessToken)
	resp, _ := http.PostForm(MakeApiUrl(apiName), params)
	d := json.NewDecoder(resp.Body)
	d.Decode(ret)
}

type ReadNamer interface{
	io.Reader
	Name()string
}

func (a *App) Post(apiName string, params url.Values, files map[string]ReadNamer, ret interface{}) {
	params.Set("access_token", a.accessToken)

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
