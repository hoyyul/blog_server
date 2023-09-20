package main

import (
	"blog_server/global"
	"blog_server/initialization"
	"blog_server/models"
	"blog_server/service/es_service"
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func main() {
	initialization.InitConf()
	initialization.InitLogger()
	global.ESClient = initialization.EsConnect()

	// get every record in article index
	boolSearch := elastic.NewMatchAllQuery()
	res, _ := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())

	// handle each record
	for _, hit := range res.Hits.Hits {
		var article models.ArticleModel
		_ = json.Unmarshal(hit.Source, &article)

		// save record info to struct
		indexList := es_service.GetSearchIndexDataByContent(hit.Id, article.Title, article.Content)

		// add struct to full text search index
		bulk := global.ESClient.Bulk()
		for _, indexData := range indexList { // use bulk add here because one id may have two records
			req := elastic.NewBulkIndexRequest().Index(models.FullTextModel{}.Index()).Doc(indexData)
			bulk.Add(req)
		}
		result, err := bulk.Do(context.Background())
		if err != nil {
			logrus.Error(err)
			continue
		}
		fmt.Printf("Sychro %d indexs to es successfully", len(result.Succeeded()))
	}

}
