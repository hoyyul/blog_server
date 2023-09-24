package statistic_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"time"

	"github.com/gin-gonic/gin"
)

type DateCount struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type DateCountResponse struct {
	DateList     []string `json:"date_list"`
	LoginData    []int    `json:"login_data"`
	RegisterData []int    `json:"sign_data"`
}

func (StatisticApi) SevenDayLoginView(c *gin.Context) {
	var loginDateCount, registerDateCount []DateCount

	global.DB.Model(models.LoginDataModel{}).
		Where("date_sub(curdate(), interval 7 day) <= created_at").
		Select("date_format(created_at, '%Y-%m-%d') as date", "count(id) as count").
		Group("date").
		Scan(&loginDateCount)
	global.DB.Model(models.UserModel{}).
		Where("date_sub(curdate(), interval 7 day) <= created_at").
		Select("date_format(created_at, '%Y-%m-%d') as date", "count(id) as count").
		Group("date").
		Scan(&registerDateCount)
	var loginDateCountMap = map[string]int{}
	var registerDateCountMap = map[string]int{}
	var loginCountList, registerCountList []int
	now := time.Now()
	for _, i2 := range loginDateCount {
		loginDateCountMap[i2.Date] = i2.Count
	}
	for _, i2 := range registerDateCount {
		registerDateCountMap[i2.Date] = i2.Count
	}
	var dateList []string
	for i := -6; i <= 0; i++ {
		day := now.AddDate(0, 0, i).Format("2006-01-02")
		loginCount := loginDateCountMap[day]
		registerCount := registerDateCountMap[day]
		dateList = append(dateList, day)
		loginCountList = append(loginCountList, loginCount)
		registerCountList = append(registerCountList, registerCount)
	}

	res.OkWithData(DateCountResponse{
		DateList:     dateList,
		LoginData:    loginCountList,
		RegisterData: registerCountList,
	}, c)

}
