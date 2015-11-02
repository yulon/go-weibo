package weibo

import (
	"net/url"
	"strconv"
)

type RemindService struct{
	c *Client
}

//读取接口

//获取某个用户的各种消息未读数
func (rs *RemindService) UnreadCount(uid int64) (r *Remind) {
	p := url.Values{}
	p.Set("uid", strconv.FormatInt(uid, 10))
	r = &Remind{}
	rs.c.Get("remind/unread_count", p, r)
	return
}

//写入接口

//对当前登录用户某一种消息未读数进行清零
func (rs *RemindService) SetCount(typ string) bool {
	p := url.Values{}
	p.Set("type", typ)
	ret := &struct{
		Result bool `json: result`
	}{}
	rs.c.PostForm("remind/set_count", p, ret)
	return ret.Result
}
