package routers

import "blog_server/api"

func (rg RouterGroup) StatisticRouter() {
	app := api.ApiGroupApp.StatisticApi
	rg.GET("data_login", app.SevenDayLoginView)
	rg.GET("data_sum", app.StatisticSumView)
}
