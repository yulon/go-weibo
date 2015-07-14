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
