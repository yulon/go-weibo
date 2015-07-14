package weibo

import (
	"net/url"
	"strconv"
)

type Statuses struct{
	app *App
}

func NewStatuses(app *App) *Statuses {
	return &Statuses{app}
}

//读取接口

func (s *Statuses) PublicTimeline(count int, page int, baseApp int) []*Status { //返回最新的公共微博
	p := &url.Values{}
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("base_app", strconv.Itoa(baseApp))
	return s.app.GetStatuses("statuses/public_timeline", p)
}

func (s *Statuses) FriendsTimeline(sinceId int64, maxId int64, count int, page int, baseApp int, feature int, trimUser int) []*Status { //获取当前登录用户及其所关注用户的最新微博
	p := &url.Values{}
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("base_app", strconv.Itoa(baseApp))
	p.Set("feature", strconv.Itoa(feature))
	p.Set("trim_user", strconv.Itoa(trimUser))
	return s.app.GetStatuses("statuses/friends_timeline", p)
}

func (s *Statuses) HomeTimeline(sinceId int64, maxId int64, count int, page int, baseApp int, feature int, trimUser int) []*Status { //获取当前登录用户及其所关注用户的最新微博
	p := &url.Values{}
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("base_app", strconv.Itoa(baseApp))
	p.Set("feature", strconv.Itoa(feature))
	p.Set("trim_user", strconv.Itoa(trimUser))
	return s.app.GetStatuses("statuses/home_timeline", p)
}

func (s *Statuses) FriendsTimelineIds(sinceId int64, maxId int64, count int, page int, baseApp int, feature int) []string { //获取当前登录用户及其所关注用户的最新微博的ID
	p := &url.Values{}
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("base_app", strconv.Itoa(baseApp))
	p.Set("feature", strconv.Itoa(feature))
	return s.app.GetStatusesIds("statuses/friends_timeline/ids", p)
}

func (s *Statuses) UserTimeline(uid int64, screenName string, sinceId int64, maxId int64, count int, page int, baseApp int, feature int) []*Status { //获取某个用户最新发表的微博列表
	p := &url.Values{}
	p.Set("uid", strconv.FormatInt(uid, 10))
	p.Set("screen_name", screenName)
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("base_app", strconv.Itoa(baseApp))
	p.Set("feature", strconv.Itoa(feature))
	return s.app.GetStatuses("statuses/user_timeline", p)
}

func (s *Statuses) UserTimelineIds(uid int64, screenName string, sinceId int64, maxId int64, count int, page int, baseApp int, feature int) []string { //获取用户发布的微博的ID
	p := &url.Values{}
	p.Set("uid", strconv.FormatInt(uid, 10))
	p.Set("screen_name", screenName)
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("base_app", strconv.Itoa(baseApp))
	p.Set("feature", strconv.Itoa(feature))
	return s.app.GetStatusesIds("statuses/user_timeline/ids", p)
}

func (s *Statuses) TimelineBatch(uids string, screenNames string, count int, page int, baseApp int, feature int) []*Status { //【高】批量获取指定的一批用户的微博列表
	p := &url.Values{}
	p.Set("uids", uids)
	p.Set("screen_names", screenNames)
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("base_app", strconv.Itoa(baseApp))
	p.Set("feature", strconv.Itoa(feature))
	return s.app.GetStatuses("statuses/timeline_batch", p)
}

func (s *Statuses) RepostTimeline(id int64, sinceId int64, maxId int64, count int, page int, filterByAuthor int) []*Status { //获取指定微博的转发微博列表
	p := &url.Values{}
	p.Set("id", strconv.FormatInt(id, 10))
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("filter_by_author", strconv.Itoa(filterByAuthor))
	return s.app.GetStatuses("statuses/repost_timeline", p)
}

func (s *Statuses) RepostTimelineIds(id int64, sinceId int64, maxId int64, count int, page int, filterByAuthor int) []string { //获取一条原创微博的最新转发微博的ID
	p := &url.Values{}
	p.Set("id", strconv.FormatInt(id, 10))
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("filter_by_author", strconv.Itoa(filterByAuthor))
	return s.app.GetStatusesIds("statuses/repost_timeline/ids", p)
}

func (s *Statuses) Mentions(sinceId int64, maxId int64, count int, page int, filterByAuthor int, filterBySource int, filterByType int) []*Status { //获取最新的提到登录用户的微博列表，即@我的微博
	p := &url.Values{}
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("filter_by_author", strconv.Itoa(filterByAuthor))
	p.Set("filter_by_source", strconv.Itoa(filterBySource))
	p.Set("filter_by_type", strconv.Itoa(filterByType))
	return s.app.GetStatuses("statuses/mentions", p)
}

func (s *Statuses) MentionsIds(sinceId int64, maxId int64, count int, page int, filterByAuthor int, filterBySource int, filterByType int) []string { //获取@当前用户的最新微博的ID
	p := &url.Values{}
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("filter_by_author", strconv.Itoa(filterByAuthor))
	p.Set("filter_by_source", strconv.Itoa(filterBySource))
	p.Set("filter_by_type", strconv.Itoa(filterByType))
	return s.app.GetStatusesIds("statuses/mentions/ids", p)
}

func (s *Statuses) BilateralTimeline(sinceId int64, maxId int64, count int, page int, baseApp int, feature int, trimUser int) []*Status { //获取双向关注用户的最新微博
	p := &url.Values{}
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("base_app", strconv.Itoa(baseApp))
	p.Set("feature", strconv.Itoa(feature))
	p.Set("trim_user", strconv.Itoa(trimUser))
	return s.app.GetStatuses("statuses/bilateral_timeline", p)
}

func (s *Statuses) Show(id int64) *Status { //根据微博ID获取单条微博内容
	p := &url.Values{}
	p.Set("id", strconv.FormatInt(id, 10))
	return s.app.GetStatus("statuses/show", p)
}

func (s *Statuses) ShowBatch(ids string, trimUser int) []*Status { //【高】根据微博ID获取单条微博内容
	p := &url.Values{}
	p.Set("ids", ids)
	p.Set("trim_user", strconv.Itoa(trimUser))
	return s.app.GetStatuses("statuses/show_batch", p)
}

func (s *Statuses) Querymid(id int64, typ int, isBatch int) string { //通过微博（评论、私信）ID获取其MID
	p := &url.Values{}
	p.Set("id", strconv.FormatInt(id, 10))
	p.Set("type", strconv.Itoa(typ))
	p.Set("is_batch", strconv.Itoa(isBatch))
	return s.app.GetMid("statuses/querymid", p)
}

func (s *Statuses) Queryid(mid string, typ int, isBatch int, inbox int, isBase62 int) string { //通过微博（评论、私信）MID获取其ID
	p := &url.Values{}
	p.Set("mid", mid)
	p.Set("type", strconv.Itoa(typ))
	p.Set("is_batch", strconv.Itoa(isBatch))
	p.Set("inbox", strconv.Itoa(inbox))
	p.Set("isBase62", strconv.Itoa(isBase62))
	return s.app.GetId("statuses/queryid", p)
}

func (s *Statuses) Count(ids string) (scs []*StatusCount) { //批量获取指定微博的转发数评论数
	p := &url.Values{}
	p.Set("ids", ids)
	s.app.Get("statuses/count", p, scs)
	return
}

func (s *Statuses) Go(uid int64, id int64) string { //生成根据ID跳转到单条微博页的URL
	p := &url.Values{}
	p.Set("uid", strconv.FormatInt(uid, 10))
	p.Set("id", strconv.FormatInt(id, 10))
	return "http://api.weibo.com/2/statuses/go?" + p.Encode()
}
