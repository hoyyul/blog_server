package models

import (
	"blog_server/global"
	"blog_server/models/ctype"
	"context"

	"github.com/sirupsen/logrus"
)

type ArticleModel struct {
	ID        string `json:"id"` // id for es
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`

	Title    string `gorm:"size:32" json:"title"`
	Abstract string `json:"abstract"`
	Content  string `json:"content,omit(list)"`

	LookCount     int `json:"look_count"`
	CommentCount  int `json:"comment_count"`
	DiggCount     int `json:"digg_count"`
	CollectsCount int `json:"collects_count"`

	UserID       uint   `json:"user_id"`
	UserNickName string `json:"user_nick_name"`
	UserAvatar   string `json:"user_avatar"`

	Category string `gorm:"size:20" json:"category"`
	Source   string `json:"source,omit(list)"`
	Link     string `json:"link,omit(list)"`

	BannerID  uint   `json:"cover_id"`
	BannerUrl string `json:"banner_url"`

	Tags ctype.Array `gorm:"type:string;size:64" json:"tags"`
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
        "type": "text"
      },
      "user_avatar": { 
        "type": "text"
      },
      "category": { 
        "type": "text"
      },
      "source": { 
        "type": "text"
      },
      "link": { 
        "type": "text"
      },
      "banner_id": {
        "type": "integer"
      },
      "banner_url": { 
        "type": "text"
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
		BodyJson(a).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	a.ID = indexResponse.Id
	return nil
}
