package weibo

import (
	"net/url"
	"strconv"
)

type Reminder struct{
	app *App
}

func NewReminder(app *App) *Reminder {
	return &Reminder{app}
}

//读取接口

//获取某个用户的各种消息未读数
func (r *Reminder) UnreadCount(uid int64) (ret *Remind) {
	p := url.Values{}
	p.Set("uid", strconv.FormatInt(uid, 10))
	r.app.Get("remind/unread_count", p, ret)
	return
}

//写入接口

//对当前登录用户某一种消息未读数进行清零
func (r *Reminder) SetCount(typ string) bool {
	p := url.Values{}
	p.Set("type", typ)
	ret := &struct{
		Result bool `json: result`
	}{}
	r.app.PostForm("remind/set_count", p, ret)
	return ret.Result
}
