package weibo

import (
	"net/url"
	"strconv"
)

type IfRemind struct{
	app *App
}

func NewIfRemind(app *App) *IfRemind {
	return &IfRemind{app}
}

//读取接口

//获取某个用户的各种消息未读数
func (ifr *IfRemind) UnreadCount(uid int64) (r *Remind) {
	p := url.Values{}
	p.Set("uid", strconv.FormatInt(uid, 10))
	ifr.app.Get("remind/unread_count", p, r)
	return
}

//写入接口

//对当前登录用户某一种消息未读数进行清零
func (ifr *IfRemind) SetCount(typ string) bool {
	p := url.Values{}
	p.Set("type", typ)
	ret := &struct{
		Result bool `json: result`
	}{}
	ifr.app.PostForm("remind/set_count", p, ret)
	return ret.Result
}
