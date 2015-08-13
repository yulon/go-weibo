package weibo

import (
	"net/url"
	"strconv"
)

type IfUser struct{
	app *App
}

func NewIfUser(app *App) *IfUser {
	return &IfUser{app}
}

//读取接口

//根据用户ID获取用户信息
func (ifu *IfUser) Show(uid int64) *User {
	p := url.Values{}
	p.Set("uid", strconv.FormatInt(uid, 10))
	u := &User{}
	ifu.app.Get("users/show", p, u)
	return u
}

//根据用户昵称获取用户信息
func (ifu *IfUser) ScreenNameShow(screenName string) *User {
	p := url.Values{}
	p.Set("screen_name", screenName)
	u := &User{}
	ifu.app.Get("users/show", p, u)
	return u
}

//通过个性化域名获取用户资料以及用户最新的一条微博
func (ifu *IfUser) DomainShow(domain string) (u *User) {
	p := url.Values{}
	p.Set("domain", domain)
	u = &User{}
	ifu.app.Get("users/domain_show", p, u)
	return u
}

//批量获取用户的粉丝数、关注数、微博数
func (ifu *IfUser) Counts(uids string) (us []*UserCounts) {
	p := url.Values{}
	p.Set("uids", uids)
	ifu.app.Get("users/counts", p, us)
	return
}

//获取用户等级信息
func (ifu *IfUser) ShowRank(uid int64) *UserRank {
	p := url.Values{}
	p.Set("uid", strconv.FormatInt(uid, 10))
	ur := &UserRank{}
	ifu.app.Get("users/show_rank", p, ur)
	return ur
}
