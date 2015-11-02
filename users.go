package weibo

import (
	"net/url"
	"strconv"
)

type UsersService struct{
	c *Client
}

//读取接口

//根据用户ID获取用户信息
func (uss *UsersService) Show(uid int64) *User {
	p := url.Values{}
	p.Set("uid", strconv.FormatInt(uid, 10))
	u := &User{}
	uss.c.Get("users/show", p, u)
	return u
}

//根据用户昵称获取用户信息
func (uss *UsersService) ScreenNameShow(screenName string) *User {
	p := url.Values{}
	p.Set("screen_name", screenName)
	u := &User{}
	uss.c.Get("users/show", p, u)
	return u
}

//通过个性化域名获取用户资料以及用户最新的一条微博
func (uss *UsersService) DomainShow(domain string) *User {
	p := url.Values{}
	p.Set("domain", domain)
	u := &User{}
	uss.c.Get("users/domain_show", p, u)
	return u
}

//批量获取用户的粉丝数、关注数、微博数
func (uss *UsersService) Counts(uids string) (us []*UserCounts) {
	p := url.Values{}
	p.Set("uids", uids)
	uss.c.Get("users/counts", p, us)
	return
}

//获取用户等级信息
func (uss *UsersService) ShowRank(uid int64) *UserRank {
	p := url.Values{}
	p.Set("uid", strconv.FormatInt(uid, 10))
	ur := &UserRank{}
	uss.c.Get("users/show_rank", p, ur)
	return ur
}
