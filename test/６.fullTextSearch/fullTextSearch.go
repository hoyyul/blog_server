package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
)

func main() {
	var data = "## 环境搭建\n\n拉取镜像\n\n```Python\ndocker pull elasticsearch:7.12.0\n```\n\n\n\n创建docker容器挂在的目录:\n\n```Python\nmkdir -p /opt/elasticsearch/config & mkdir -p /opt/elasticsearch/data & mkdir -p /opt/elasticsearch/plugins\n\nchmod 777 /opt/elasticsearch/data\n\n```\n\n配置文件\n\n```Python\necho \"http.host: 0.0.0.0\" >> /opt/elasticsearch/config/elasticsearch.yml\n```\n\n\n\n创建容器\n\n```Python\n# linux\ndocker run --name es -p 9200:9200  -p 9300:9300 -e \"discovery.type=single-node\" -e ES_JAVA_OPTS=\"-Xms84m -Xmx512m\" -v /opt/elasticsearch/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml -v /opt/elasticsearch/data:/usr/share/elasticsearch/data -v /opt/elasticsearch/plugins:/usr/share/elasticsearch/plugins -d elasticsearch:7.12.0\n```\n\n\n\n访问ip:9200能看到东西\n\n![](http://python.fengfengzhidao.com/pic/20230129212040.png)\n\n就说明安装成功了\n\n\n\n浏览器可以下载一个 `Multi Elasticsearch Head` es插件\n\n\n\n第三方库\n\n```Go\ngithub.com/olivere/elastic/v7\n```\n\n## es连接\n\n```Go\nfunc EsConnect() *elastic.Client  {\n  var err error\n  sniffOpt := elastic.SetSniff(false)\n  host := \"http://127.0.0.1:9200\"\n  c, err := elastic.NewClient(\n    elastic.SetURL(host),\n    sniffOpt,\n    elastic.SetBasicAuth(\"\", \"\"),\n  )\n  if err != nil {\n    logrus.Fatalf(\"es连接失败 %s\", err.Error())\n  }\n  return c\n}\n```"
	GetSearchIndexDataByContent("/article/hd893bxGHD84", "es的环境搭建", data)
}

type SearchData struct {
	Body  string `json:"body"`  // content
	Slug  string `json:"slug"`  // redirect
	Title string `json:"title"` // article title
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
		})
	}
	b, _ := json.Marshal(searchDataList)
	fmt.Println(string(b))
	fmt.Println(len(bodyList), len(titleList))
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
