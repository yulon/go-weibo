package weibo

import (
	"net/http"
	"net/url"
	"encoding/json"
)

const (
	api = "https://api.weibo.com/2"
)

type App struct{
	accessToken string
}

func NewApp(accessToken string) *App {
	return &App{accessToken}
}

func (a *App) Get(name string, param *url.Values, ret interface{}) {
	param.Set("access_token", a.accessToken)
	resp, _ := http.Get(api + "/" + name + ".json?" + param.Encode())
	d := json.NewDecoder(resp.Body)
	d.Decode(ret)
}

func (a *App) GetStatus(name string, param *url.Values) (ret *Status) {
	ret = &Status{}
	a.Get(name, param, ret)
	return
}

func (a *App) GetStatuses(name string, param *url.Values) []*Status {
	ret := &struct{
		Statuses []*Status `json: statuses`
	}{}
	a.Get(name, param, ret)
	return ret.Statuses
}

func (a *App) GetStatusesIds(name string, param *url.Values) []string {
	ret := &struct{
		Statuses []string `json: statuses`
	}{}
	a.Get(name, param, ret)
	return ret.Statuses
}
