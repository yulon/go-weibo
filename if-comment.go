package weibo

import (
	"net/url"
	"strconv"
)

type IfComment struct{
	app *App
}

func NewIfComment(app *App) *IfComment {
	return &IfComment{app}
}

//读取接口

func (ifc *IfComment) getComments(apiName string, param url.Values) []*Comment {
	ret := &struct{
		Comments []*Comment `json: comments`
	}{}
	ifc.app.Get(apiName, param, ret)
	return ret.Comments
}

//根据微博ID返回某条微博的评论列表
func (ifc *IfComment) Show(id int64, sinceId int64, maxId int64, count int, page int, filterByAuthor int) []*Comment {
	p := url.Values{}
	p.Set("id", strconv.FormatInt(id, 10))
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("filter_by_author", strconv.Itoa(filterByAuthor))
	return ifc.getComments("comments/show", p)
}

//获取当前登录用户所发出的评论列表
func (ifc *IfComment) ByMe(sinceId int64, maxId int64, count int, page int, filterBySource int) []*Comment {
	p := url.Values{}
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("filter_by_source", strconv.Itoa(filterBySource))
	return ifc.getComments("comments/by_me", p)
}

//获取当前登录用户所接收到的评论列表
func (ifc *IfComment) ToMe(sinceId int64, maxId int64, count int, page int, filterByAuthor int, filterBySource int) []*Comment {
	p := url.Values{}
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("filter_by_author", strconv.Itoa(filterByAuthor))
	p.Set("filter_by_source", strconv.Itoa(filterBySource))
	return ifc.getComments("comments/to_me", p)
}

//获取当前登录用户的最新评论包括接收到的与发出的
func (ifc *IfComment) Timeline(sinceId int64, maxId int64, count int, page int, trimUser int) []*Comment {
	p := url.Values{}
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("trim_user", strconv.Itoa(trimUser))
	return ifc.getComments("comments/timeline", p)
}

//获取最新的提到当前登录用户的评论，即@我的评论
func (ifc *IfComment) Mentions(sinceId int64, maxId int64, count int, page int, filterByAuthor int, filterBySource int) []*Comment {
	p := url.Values{}
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("filter_by_author", strconv.Itoa(filterByAuthor))
	p.Set("filter_by_source", strconv.Itoa(filterBySource))
	return ifc.getComments("comments/mentions", p)
}

//根据评论ID批量返回评论信息
func (ifc *IfComment) ShowBatch(cids int64) (cs []*Comment) {
	p := url.Values{}
	p.Set("cids", strconv.FormatInt(cids, 10))
	ifc.app.Get("comments/show_batch", p, cs)
	return
}

//写入接口

//对一条微博进行评论
func (ifc *IfComment) Create(comment string, id int64, commentOri int, rip string) (c *Comment) {
	p := url.Values{}
	p.Set("comment", comment)
	p.Set("id", strconv.FormatInt(id, 10))
	p.Set("comment_ori", strconv.Itoa(commentOri))
	p.Set("rip", rip)
	ifc.app.PostForm("comments/create", p, c)
	return
}

//删除一条评论
func (ifc *IfComment) Destroy(cid int64) (c *Comment) {
	p := url.Values{}
	p.Set("cid", strconv.FormatInt(cid, 10))
	ifc.app.PostForm("comments/destroy", p, c)
	return
}

//根据评论ID批量删除评论
func (ifc *IfComment) DestroyBatch(cids string) (cs []*Comment) {
	p := url.Values{}
	p.Set("cids", cids)
	ifc.app.PostForm("comments/destroy_batch", p, cs)
	return
}

//回复一条评论
func (ifc *IfComment) Reply(cid int64, id int64, comment string, withoutMention int, commentOri int, rip string) (cs []*Comment) {
	p := url.Values{}
	p.Set("cid", strconv.FormatInt(cid, 10))
	p.Set("id", strconv.FormatInt(id, 10))
	p.Set("comment", comment)
	p.Set("without_mention", strconv.Itoa(withoutMention))
	p.Set("comment_ori", strconv.Itoa(commentOri))
	p.Set("rip", rip)
	ifc.app.PostForm("comments/reply", p, cs)
	return
}