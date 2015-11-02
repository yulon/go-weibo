package weibo

import (
	"net/http"
	"net/url"
	"encoding/json"
)

type App struct{
	key string
	secret string
	cb string
}

func NewApp(appKey string, appSecret string, callbackURL string) *App {
	return &App{
		key: appKey,
		secret: appSecret,
		cb: callbackURL,
	}
}

func (a *App) MakeAuthorizeURL(state string, mobile bool) string {
	p := url.Values{
		"client_id": {a.key},
		"redirect_uri": {a.cb},
		"state": {state},
	}

	var preURL string
	if mobile {
		preURL = "https://open.weibo.cn/oauth2/authorize?"
		p.Set("display", "mobile")
	}else{
		preURL = "https://api.weibo.com/oauth2/authorize?"
	}

	return preURL + p.Encode()
}

type AccessToken struct{
	AccessToken string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
	Uid string `json:"uid"`
}

func (a *App) GetAccessToken(code string) (at *AccessToken) {
	resp, err := http.PostForm("https://api.weibo.com/oauth2/access_token", url.Values{
		"client_id": {a.key},
		"client_secret": {a.secret},
		"grant_type": {"authorization_code"},
		"code": {code},
		"redirect_uri": {a.cb},
	})

	if err != nil {
		return
	}

	at = &AccessToken{}
	d := json.NewDecoder(resp.Body)
	d.Decode(at)
	return
}
