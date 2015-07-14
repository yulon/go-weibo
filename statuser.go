package weibo

import (
	"net/url"
	"strconv"
)

type Statuser struct{
	app *App
}

func NewStatuser(app *App) *Statuser {
	return &Statuser{app}
}

//读取接口

func (s *Statuser) getStatuses(apiName string, param url.Values) []*Status {
	ret := &struct{
		Statuses []*Status `json: statuses`
	}{}
	s.app.Get(apiName, param, ret)
	return ret.Statuses
}

func (s *Statuser) getStatusesIds(apiName string, param url.Values) []string {
	ret := &struct{
		Statuses []string `json: statuses`
	}{}
	s.app.Get(apiName, param, ret)
	return ret.Statuses
}

//返回最新的公共微博
func (s *Statuser) PublicTimeline(count int, page int, baseApp int) []*Status {
	p := url.Values{}
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("base_app", strconv.Itoa(baseApp))
	return s.getStatuses("statuses/public_timeline", p)
}

//获取当前登录用户及其所关注用户的最新微博
func (s *Statuser) FriendsTimeline(sinceId int64, maxId int64, count int, page int, baseApp int, feature int, trimUser int) []*Status {
	p := url.Values{}
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("base_app", strconv.Itoa(baseApp))
	p.Set("feature", strconv.Itoa(feature))
	p.Set("trim_user", strconv.Itoa(trimUser))
	return s.getStatuses("statuses/friends_timeline", p)
}

//获取当前登录用户及其所关注用户的最新微博
func (s *Statuser) HomeTimeline(sinceId int64, maxId int64, count int, page int, baseApp int, feature int, trimUser int) []*Status {
	p := url.Values{}
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("base_app", strconv.Itoa(baseApp))
	p.Set("feature", strconv.Itoa(feature))
	p.Set("trim_user", strconv.Itoa(trimUser))
	return s.getStatuses("statuses/home_timeline", p)
}

//获取当前登录用户及其所关注用户的最新微博的ID
func (s *Statuser) FriendsTimelineIds(sinceId int64, maxId int64, count int, page int, baseApp int, feature int) []string {
	p := url.Values{}
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("base_app", strconv.Itoa(baseApp))
	p.Set("feature", strconv.Itoa(feature))
	return s.getStatusesIds("statuses/friends_timeline/ids", p)
}

//获取某个用户最新发表的微博列表
func (s *Statuser) UserTimeline(uid int64, screenName string, sinceId int64, maxId int64, count int, page int, baseApp int, feature int) []*Status {
	p := url.Values{}
	p.Set("uid", strconv.FormatInt(uid, 10))
	p.Set("screen_name", screenName)
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("base_app", strconv.Itoa(baseApp))
	p.Set("feature", strconv.Itoa(feature))
	return s.getStatuses("statuses/user_timeline", p)
}

//获取用户发布的微博的ID
func (s *Statuser) UserTimelineIds(uid int64, screenName string, sinceId int64, maxId int64, count int, page int, baseApp int, feature int) []string {
	p := url.Values{}
	p.Set("uid", strconv.FormatInt(uid, 10))
	p.Set("screen_name", screenName)
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("base_app", strconv.Itoa(baseApp))
	p.Set("feature", strconv.Itoa(feature))
	return s.getStatusesIds("statuses/user_timeline/ids", p)
}

//【高】批量获取指定的一批用户的微博列表
func (s *Statuser) TimelineBatch(uids string, screenNames string, count int, page int, baseApp int, feature int) []*Status {
	p := url.Values{}
	p.Set("uids", uids)
	p.Set("screen_names", screenNames)
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("base_app", strconv.Itoa(baseApp))
	p.Set("feature", strconv.Itoa(feature))
	return s.getStatuses("statuses/timeline_batch", p)
}

//获取指定微博的转发微博列表
func (s *Statuser) RepostTimeline(id int64, sinceId int64, maxId int64, count int, page int, filterByAuthor int) []*Status {
	p := url.Values{}
	p.Set("id", strconv.FormatInt(id, 10))
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("filter_by_author", strconv.Itoa(filterByAuthor))
	return s.getStatuses("statuses/repost_timeline", p)
}

//获取一条原创微博的最新转发微博的ID
func (s *Statuser) RepostTimelineIds(id int64, sinceId int64, maxId int64, count int, page int, filterByAuthor int) []string {
	p := url.Values{}
	p.Set("id", strconv.FormatInt(id, 10))
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("filter_by_author", strconv.Itoa(filterByAuthor))
	return s.getStatusesIds("statuses/repost_timeline/ids", p)
}

//获取最新的提到登录用户的微博列表，即@我的微博
func (s *Statuser) Mentions(sinceId int64, maxId int64, count int, page int, filterByAuthor int, filterBySource int, filterByType int) []*Status {
	p := url.Values{}
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("filter_by_author", strconv.Itoa(filterByAuthor))
	p.Set("filter_by_source", strconv.Itoa(filterBySource))
	p.Set("filter_by_type", strconv.Itoa(filterByType))
	return s.getStatuses("statuses/mentions", p)
}

//获取@当前用户的最新微博的ID
func (s *Statuser) MentionsIds(sinceId int64, maxId int64, count int, page int, filterByAuthor int, filterBySource int, filterByType int) []string {
	p := url.Values{}
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("filter_by_author", strconv.Itoa(filterByAuthor))
	p.Set("filter_by_source", strconv.Itoa(filterBySource))
	p.Set("filter_by_type", strconv.Itoa(filterByType))
	return s.getStatusesIds("statuses/mentions/ids", p)
}

//获取双向关注用户的最新微博
func (s *Statuser) BilateralTimeline(sinceId int64, maxId int64, count int, page int, baseApp int, feature int, trimUser int) []*Status {
	p := url.Values{}
	p.Set("since_id", strconv.FormatInt(sinceId, 10))
	p.Set("max_id", strconv.FormatInt(maxId, 10))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	p.Set("base_app", strconv.Itoa(baseApp))
	p.Set("feature", strconv.Itoa(feature))
	p.Set("trim_user", strconv.Itoa(trimUser))
	return s.getStatuses("statuses/bilateral_timeline", p)
}

//根据微博ID获取单条微博内容
func (s *Statuser) Show(id int64) (ret *Status) {
	p := url.Values{}
	p.Set("id", strconv.FormatInt(id, 10))
	s.app.Get("statuses/show", p, ret)
	return
}

//【高】根据微博ID获取单条微博内容
func (s *Statuser) ShowBatch(ids string, trimUser int) []*Status {
	p := url.Values{}
	p.Set("ids", ids)
	p.Set("trim_user", strconv.Itoa(trimUser))
	return s.getStatuses("statuses/show_batch", p)
}

//通过微博（评论、私信）ID获取其MID
func (s *Statuser) Querymid(id int64, typ int, isBatch int) string {
	p := url.Values{}
	p.Set("id", strconv.FormatInt(id, 10))
	p.Set("type", strconv.Itoa(typ))
	p.Set("is_batch", strconv.Itoa(isBatch))
	ret := &struct{
		Mid string `json: mid`
	}{}
	s.app.Get("statuses/querymid", p, ret)
	return ret.Mid
}

//通过微博（评论、私信）MID获取其ID
func (s *Statuser) Queryid(mid string, typ int, isBatch int, inbox int, isBase62 int) string {
	p := url.Values{}
	p.Set("mid", mid)
	p.Set("type", strconv.Itoa(typ))
	p.Set("is_batch", strconv.Itoa(isBatch))
	p.Set("inbox", strconv.Itoa(inbox))
	p.Set("isBase62", strconv.Itoa(isBase62))
	ret := &struct{
		Id string `json: id`
	}{}
	s.app.Get("statuses/queryid", p, ret)
	return ret.Id
}

//批量获取指定微博的转发数评论数
func (s *Statuser) Count(ids string) (scs []*StatusCount) {
	p := url.Values{}
	p.Set("ids", ids)
	s.app.Get("statuses/count", p, scs)
	return
}

//生成根据ID跳转到单条微博页的URL
func (s *Statuser) Go(uid int64, id int64) string {
	p := url.Values{}
	p.Set("uid", strconv.FormatInt(uid, 10))
	p.Set("id", strconv.FormatInt(id, 10))
	return "http://api.weibo.com/2/statuses/go?" + p.Encode()
}

//获取微博官方表情的详细信息
func (s *Statuser) Emotions(typ string, language string) (e []*Emotion) {
	p := url.Values{}
	p.Set("type", typ)
	p.Set("language", language)
	s.app.Get("emotions", p, e)
	return
}

//写入接口

//转发一条微博
func (s *Statuser) Repost(id int64, status string, isComment int, rip string) (ret *Status) {
	p := url.Values{}
	p.Set("id", strconv.FormatInt(id, 10))
	p.Set("status", status)
	p.Set("is_comment", strconv.Itoa(isComment))
	p.Set("rip", rip)
	s.app.PostForm("statuses/repost", p, ret)
	return
}

//根据微博ID删除指定微博
func (s *Statuser) Destroy(id int64) (ret *Status) {
	p := url.Values{}
	p.Set("id", strconv.FormatInt(id, 10))
	s.app.PostForm("statuses/destroy", p, ret)
	return
}

//发布一条新微博
func (s *Statuser) Update(status string, visible int, listId string, lat float64, long float64, annotations string, rip string) (ret *Status) {
	p := url.Values{}
	p.Set("status", status)
	p.Set("visible", strconv.Itoa(visible))
	p.Set("list_id", listId)
	p.Set("lat", strconv.FormatFloat(lat, 'f', 10, 64))
	p.Set("long", strconv.FormatFloat(long, 'f', 10, 64))
	p.Set("annotations", annotations)
	p.Set("rip", rip)
	s.app.PostForm("statuses/update", p, ret)
	return
}

//上传图片并发布一条新微博
func (s *Statuser) Upload(status string, visible int, listId string, pic []byte, lat float64, long float64, annotations string, rip string) (ret *Status) {
	p := map[string][]byte{
		"status": []byte(status),
		"visible": []byte(strconv.Itoa(visible)),
		"list_id": []byte(listId),
		"pic": pic,
		"lat": []byte(strconv.FormatFloat(lat, 'f', 10, 64)),
		"long": []byte(strconv.FormatFloat(long, 'f', 10, 64)),
		"annotations": []byte(annotations),
		"rip": []byte(rip),
	}
	s.app.Post("statuses/upload", p, ret)
	return
}

//【高】指定一个图片URL地址抓取后上传并同时发布一条新微博
func (s *Statuser) UploadUrlText(status string, visible int, listId string, urlText string, lat float64, long float64, annotations string, rip string) (ret *Status) {
	p := url.Values{}
	p.Set("status", status)
	p.Set("visible", strconv.Itoa(visible))
	p.Set("list_id", listId)
	p.Set("url", urlText)
	p.Set("lat", strconv.FormatFloat(lat, 'f', 10, 64))
	p.Set("long", strconv.FormatFloat(long, 'f', 10, 64))
	p.Set("annotations", annotations)
	p.Set("rip", rip)
	s.app.PostForm("statuses/upload_url_text", p, ret)
	return
}

//【高】使用已上传的图片发布一条新微博
func (s *Statuser) UploadPicId(status string, visible int, listId string, picId string, lat float64, long float64, annotations string, rip string) (ret *Status) {
	p := url.Values{}
	p.Set("status", status)
	p.Set("visible", strconv.Itoa(visible))
	p.Set("list_id", listId)
	p.Set("pic_id", picId)
	p.Set("lat", strconv.FormatFloat(lat, 'f', 10, 64))
	p.Set("long", strconv.FormatFloat(long, 'f', 10, 64))
	p.Set("annotations", annotations)
	p.Set("rip", rip)
	s.app.PostForm("statuses/upload_url_text", p, ret)
	return
}

//【高】屏蔽某条微博
func (s *Statuser) FilterCreate(id int64) (ret *Status) {
	p := url.Values{}
	p.Set("id", strconv.FormatInt(id, 10))
	s.app.PostForm("statuses/filter/create", p, ret)
	return
}

//【高】屏蔽某个@到我的微博以及后续由对其转发引起的@提及
func (s *Statuser) MentionsShield(id int64, followUp int) (ret *Status) {
	p := url.Values{}
	p.Set("id", strconv.FormatInt(id, 10))
	p.Set("follow_up", strconv.Itoa(followUp))
	s.app.PostForm("statuses/mentions/shield", p, ret)
	return
}
