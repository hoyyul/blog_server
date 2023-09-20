package es_service

import (
	"blog_server/global"
	"blog_server/models"
	"context"
	"encoding/json"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/olivere/elastic/v7"
	"github.com/russross/blackfriday"
	"github.com/sirupsen/logrus"
)

type SearchData struct {
	Key   string `json:"key"`
	Body  string `json:"body"`  // content
	Slug  string `json:"slug"`  // redirect
	Title string `json:"title"` // article title
}

func SynchroFullTextToES(id, title, content string) {
	// save record info to struct
	indexList := GetSearchIndexDataByContent(id, title, content)

	// add struct to full text search index
	bulk := global.ESClient.Bulk()
	for _, indexData := range indexList {
		req := elastic.NewBulkIndexRequest().Index(models.FullTextModel{}.Index()).Doc(indexData)
		bulk.Add(req)
	}
	result, err := bulk.Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	global.Logger.Info("Sychro %d indexs to es successfully", len(result.Succeeded()))
}

func DeleteFullTextByArticleID(id string) {
	boolSearch := elastic.NewTermQuery("key", id)
	res, _ := global.ESClient.
		DeleteByQuery().
		Index(models.FullTextModel{}.Index()).
		Query(boolSearch).
		Do(context.Background())
	logrus.Infof("Deleted %d records", res.Deleted)
}

// save and return markdown data to SearchData struct with plain text and title
func GetSearchIndexDataByContent(id, title, content string) (searchDataList []SearchData) {
	data := strings.Split(content, "\n")
	var isBody bool = false
	var titleList, bodyList []string
	var body string

	// add artitle title
	titleList = append(titleList, getTitle(title))

	// add header title and body to list
	for _, s := range data {
		// check if body
		if strings.HasPrefix(s, "```") {
			isBody = !isBody
		}
		if strings.HasPrefix(s, "#") && !isBody {
			titleList = append(titleList, getTitle(s))
			bodyList = append(bodyList, getBody(body)) // transfer markdown to plain text
			body = ""
			continue
		}
		body += s
	}
	bodyList = append(bodyList, getBody(body))
	ln := len(titleList)

	// wrap to searchData
	for i := 0; i < ln; i++ {
		searchDataList = append(searchDataList, SearchData{
			Title: titleList[i],
			Body:  bodyList[i],
			Slug:  id + getSlug(titleList[i]),
			Key:   id,
		})
	}
	json.Marshal(searchDataList)
	//fmt.Println(string(b))
	//fmt.Println(len(bodyList), len(titleList))
	return searchDataList
}

func getTitle(title string) string {
	title = strings.ReplaceAll(title, "#", "")
	title = strings.ReplaceAll(title, " ", "")
	return title
}

func getBody(body string) string {
	unsafe := blackfriday.MarkdownCommon([]byte(body))
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	return doc.Text()
}

func getSlug(slug string) string {
	return "#" + slug
}
