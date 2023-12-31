package article_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"context"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type CalendarResponse struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type Bucket struct {
	KeyAsString string `json:"key_as_string"`
	Key         int64  `json:"key"`
	DocCount    int    `json:"doc_count"`
}

type BucketsResponse struct {
	Buckets []Bucket `json:"buckets"`
}

var DateCount = map[string]int{}

func (ArticleApi) ArticleCalendarCountView(c *gin.Context) {
	// 1.date aggregation search
	agg := elastic.NewDateHistogramAggregation().Field("created_at").CalendarInterval("day")
	now := time.Now()
	aYearAgo := now.AddDate(-1, 0, 0)
	format := "2006-01-02 15:04:05"
	query := elastic.NewRangeQuery("created_at").
		Gte(aYearAgo.Format(format)).
		Lte(now.Format(format))

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("calendar", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Failed to get calendars", c)
		return
	}

	// 2.precess buckets(bucket with article count for each interval)
	var buckets BucketsResponse
	_ = json.Unmarshal(result.Aggregations["calendar"], &buckets)
	var resList = make([]CalendarResponse, 0)
	for _, bucket := range buckets.Buckets { // len(buckets.Buckets) <= 365
		_time, _ := time.Parse(format, bucket.KeyAsString)      // get time type by given format; bucket.KeyAsString is like 2006-01-02 00:00:00
		DateCount[_time.Format("2006-01-02")] = bucket.DocCount // _time.Format("2006-01-02") transfer the time to string
	}

	// 3.save to resList
	days := int(now.Sub(aYearAgo).Hours() / 24) // 365 or 366
	for i := 0; i <= days; i++ {
		day := aYearAgo.AddDate(0, 0, i).Format("2006-01-02")
		count := DateCount[day]
		resList = append(resList, CalendarResponse{
			Date:  day,
			Count: count,
		})
	}

	res.OkWithData(resList, c)

}
