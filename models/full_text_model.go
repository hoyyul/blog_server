package models

import (
	"blog_server/global"
	"context"

	"github.com/sirupsen/logrus"
)

type FullTextModel struct {
	ID    string `json:"id" structs:"id"`       // es id
	Key   string `json:"key"`                   // article id
	Title string `json:"title" structs:"title"` // article title
	Slug  string `json:"slug" structs:"slug"`   // redirect
	Body  string `json:"body" structs:"body"`   // body
}

func (FullTextModel) Index() string {
	return "full_text_index"
}

func (FullTextModel) Mapping() string {
	return `
{
  "settings": {
    "index":{
      "max_result_window": "100000"
    }
  }, 
  "mappings": {
    "properties": {
	  "key": {
        "type": "keyword"
      },
      "title": { 
        "type": "text"
      },
      "slug": { 
        "type": "keyword"
      },
      "body": { 
        "type": "text"
      }
    }
  }
}
`
}

func (a FullTextModel) IndexExists() bool {
	exists, err := global.ESClient.
		IndexExists(a.Index()).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return exists
	}
	return exists
}

func (a FullTextModel) CreateIndex() error {
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

func (a FullTextModel) RemoveIndex() error {
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
