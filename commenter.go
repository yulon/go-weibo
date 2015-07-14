package weibo

import (
	"net/url"
	"strconv"
)

type Commenter struct{
	app *App
}

func NewCommenter(app *App) *Commenter {
	return &Commenter{app}
}

//读取接口

func (c *Commenter) getComments(apiName string, param url.Values) []*Comment {
	ret := &struct{
		Comments []*Comment `json: comments`
	}{}
	c.app.Get(apiName, param, ret)
	return ret.Comments
}

//根据微博ID返回某条微博的评论列表
func (c *Commenter) Show(id int64, sinceId int64, maxId int64, count int, page int, filterByAuthor int) []*Comment {
	p := url.Values{}
	p.Set("id", strconv.FormatInt(id, 10))
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("filter_by_author", strconv.Itoa(filterByAuthor))
	return c.getComments("comments/show", p)
}

//获取当前登录用户所发出的评论列表
func (c *Commenter) ByMe(sinceId int64, maxId int64, count int, page int, filterBySource int) []*Comment {
	p := url.Values{}
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("filter_by_source", strconv.Itoa(filterBySource))
	return c.getComments("comments/by_me", p)
}

//获取当前登录用户所接收到的评论列表
func (c *Commenter) ToMe(sinceId int64, maxId int64, count int, page int, filterByAuthor int, filterBySource int) []*Comment {
	p := url.Values{}
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("filter_by_author", strconv.Itoa(filterByAuthor))
	p.Set("filter_by_source", strconv.Itoa(filterBySource))
	return c.getComments("comments/to_me", p)
}

//获取当前登录用户的最新评论包括接收到的与发出的
func (c *Commenter) Timeline(sinceId int64, maxId int64, count int, page int, trimUser int) []*Comment {
	p := url.Values{}
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("trim_user", strconv.Itoa(trimUser))
	return c.getComments("comments/timeline", p)
}

//获取最新的提到当前登录用户的评论，即@我的评论
func (c *Commenter) Mentions(sinceId int64, maxId int64, count int, page int, filterByAuthor int, filterBySource int) []*Comment {
	p := url.Values{}
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("filter_by_author", strconv.Itoa(filterByAuthor))
	p.Set("filter_by_source", strconv.Itoa(filterBySource))
	return c.getComments("comments/mentions", p)
}

//根据评论ID批量返回评论信息
func (c *Commenter) ShowBatch(cids int64) (cs []*Comment) {
	p := url.Values{}
	p.Set("cids", strconv.FormatInt(cids, 10))
	c.app.Get("comments/show_batch", p, cs)
	return
}
