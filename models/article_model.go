package models

import (
	"blog_server/global"
	"blog_server/models/ctype"
	"context"
	"encoding/json"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type ArticleModel struct {
	ID        string `json:"id" structs:"id"`
	CreatedAt string `json:"created_at" structs:"created_at"`
	UpdatedAt string `json:"updated_at" structs:"updated_at"`

	Title    string `json:"title" structs:"title"`
	Keyword  string `json:"keyword,omit(list)" structs:"keyword"`
	Abstract string `json:"abstract" structs:"abstract"`
	Content  string `json:"content,omit(list)" structs:"content"`

	LookCount     int `json:"look_count" structs:"look_count"`
	CommentCount  int `json:"comment_count" structs:"comment_count"`
	DiggCount     int `json:"digg_count" structs:"digg_count"`
	CollectsCount int `json:"collects_count" structs:"collects_count"`

	UserID       uint   `json:"user_id" structs:"user_id"`
	UserNickName string `json:"user_nick_name" structs:"user_nick_name"`
	UserAvatar   string `json:"user_avatar" structs:"user_avatar"`

	Category string `json:"category" structs:"category"`
	Source   string `json:"source" structs:"source"`
	Link     string `json:"link" structs:"link"`

	BannerID  uint   `json:"banner_id" structs:"banner_id"`
	BannerUrl string `json:"banner_url" structs:"banner_url"`

	Tags ctype.Array `json:"tags" structs:"tags"`
}

func (ArticleModel) Mapping() string {
	return `
{
  "settings": {
    "index":{
      "max_result_window": "100000"
    }
  }, 
  "mappings": {
    "properties": {
      "title": { 
        "type": "text"
      },
	  "keyword": { 
        "type": "keyword"
      },
      "abstract": { 
        "type": "text"
      },
      "content": { 
        "type": "text"
      },
      "look_count": {
        "type": "integer"
      },
      "comment_count": {
        "type": "integer"
      },
      "digg_count": {
        "type": "integer"
      },
      "collects_count": {
        "type": "integer"
      },
      "user_id": {
        "type": "integer"
      },
      "user_nick_name": { 
        "type": "keyword"
      },
      "user_avatar": { 
        "type": "keyword"
      },
      "category": { 
        "type": "keyword"
      },
      "source": { 
        "type": "keyword"
      },
      "link": { 
        "type": "text"
      },
      "banner_id": {
        "type": "integer"
      },
      "banner_url": { 
        "type": "keyword"
	  },
	  "tags": { 
	    "type": "keyword"
	  },
      "created_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      },
      "updated_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}
`
}

func (ArticleModel) Index() string {
	return "article_index"
}

func (a ArticleModel) IndexExists() bool {
	exist, err := global.ESClient.
		IndexExists(a.Index()).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return exist
	}
	return exist
}

func (a ArticleModel) CreateIndex() error {
	if a.IndexExists() {
		a.RemoveIndex()
	}

	createIndex, err := global.ESClient.
		CreateIndex(a.Index()).
		BodyString(a.Mapping()).
		Do(context.Background())
	if err != nil {
		logrus.Error("Failed to create index")
		logrus.Error(err.Error())
		return err
	}
	if !createIndex.Acknowledged {
		logrus.Error("Failed to create index")
		return err
	}
	logrus.Infof("Create index %s successfully", a.Index())
	return nil
}

func (a ArticleModel) RemoveIndex() error {
	indexDelete, err := global.ESClient.DeleteIndex(a.Index()).Do(context.Background())
	if err != nil {
		logrus.Error("Failed to delete index")
		logrus.Error(err.Error())
		return err
	}
	if !indexDelete.Acknowledged {
		logrus.Error("Failed to delete index")
		return err
	}
	logrus.Info("Delete index successfully")
	return nil
}

// create document to a index
func (a *ArticleModel) Create() (err error) {
	indexResponse, err := global.ESClient.Index().
		Index(a.Index()).
		BodyJson(a).Refresh("true").Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	a.ID = indexResponse.Id
	return nil
}

// search article by title
func (a ArticleModel) ISExist() bool {
	res, err := global.ESClient.
		Search(a.Index()).
		Query(elastic.NewTermQuery("keyword", a.Title)).
		Size(1).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return false
	}
	if res.Hits.TotalHits.Value > 0 {
		return true
	}
	return false
}

func (a *ArticleModel) GetDataByID(id string) error {
	res, err := global.ESClient.
		Get().
		Index(a.Index()).
		Id(id).
		Do(context.Background())
	if err != nil {
		return err
	}
	err = json.Unmarshal(res.Source, a)
	return err
}
