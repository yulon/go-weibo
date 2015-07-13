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

func (s *Statuses) PublicTimeline(count int, page int, baseApp bool) (ret *StatusList) {
	p := &url.Values{}
	p.Set("count", strconv.Itoa(count))
	p.Set("page", strconv.Itoa(page))
	ba := 0
	if baseApp {
		ba = 1
	}
	p.Set("base_app", strconv.Itoa(ba))
	ret = &StatusList{}
	s.app.Get("statuses/public_timeline", p, ret)
	return
}
