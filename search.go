package weibo

import (
	"net/url"
	"strconv"
)

type SearchService struct{
	app *App
}

//搜索联想接口

//搜索用户时的联想搜索建议
func (ss *SearchService) SuggestionsUsers(q string, count int) (ius []*InfoUser) {
	p := url.Values{}
	p.Set("q", url.QueryEscape(q))
	p.Set("count", strconv.Itoa(count))
	ss.app.Get("search/suggestions/users", p, ius)
	return
}

//搜索学校时的联想搜索建议
func (ss *SearchService) SuggestionsSchools(q string, count int, typ int) (iss []*InfoSchool) {
	p := url.Values{}
	p.Set("q", url.QueryEscape(q))
	p.Set("count", strconv.Itoa(count))
	p.Set("type", strconv.Itoa(typ))
	ss.app.Get("search/suggestions/schools", p, iss)
	return
}

//搜索公司时的联想搜索建议
func (ss *SearchService) SuggestionsCompanies(q string, count int) (ics []*InfoCompanie) {
	p := url.Values{}
	p.Set("q", url.QueryEscape(q))
	p.Set("count", strconv.Itoa(count))
	ss.app.Get("search/suggestions/companies", p, ics)
	return
}

//搜索应用时的联想搜索建议
func (ss *SearchService) SuggestionsApps(q string, count int) (ias []*InfoApp) {
	p := url.Values{}
	p.Set("q", url.QueryEscape(q))
	p.Set("count", strconv.Itoa(count))
	ss.app.Get("search/suggestions/apps", p, ias)
	return
}

//@用户时的联想建议
func (ss *SearchService) SuggestionsAtUsers(q string, count int, typ int, rang int) (iaus []*InfoAtUser) {
	p := url.Values{}
	p.Set("q", url.QueryEscape(q))
	p.Set("count", strconv.Itoa(count))
	p.Set("type", strconv.Itoa(typ))
	p.Set("range", strconv.Itoa(typ))
	ss.app.Get("search/suggestions/at_users", p, iaus)
	return
}

//搜索话题接口

//搜索某一话题下的微博
func (ss *SearchService) Topics(q string, count int, page int) []*Status {
	p := url.Values{}
	p.Set("q", url.QueryEscape(q))
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	ret := &struct{
		Statuses []*Status `json:"statuses"`
	}{}
	ss.app.Get("search/topics", p, ret)
	return ret.Statuses
}
