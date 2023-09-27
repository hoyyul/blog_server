package flag

import (
	"blog_server/global"
	"context"
	"encoding/json"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func LoadIndex(jsonName string) {
	byteData, err := os.ReadFile(jsonName)
	if err != nil {
		logrus.Error(err)
		return
	}
	var jsonData EsIndexResponse
	err = json.Unmarshal(byteData, &jsonData)
	if err != nil {
		logrus.Error(err)
		return
	}
	_list := strings.Split(jsonName, ".")
	if len(_list) != 2 {
		logrus.Error("Incorrect json file name")
		return
	}
	index := _list[0]
	bd, _ := json.Marshal(jsonData.Mapping)
	CreateIndex(index, string(bd))
	for _, data := range jsonData.Data {
		_, err := global.ESClient.Index().
			Index(index).Id(data.ID).
			BodyJson(data.Row).Do(context.Background())
		if err != nil {
			logrus.Error(err)
			continue
		}
		logrus.Infof("%s entry successful", data.ID)
	}
	logrus.Infof("Index %s entry successful", index)

}

func RemoveIndex(index string) {
	logrus.Info("Index exists, deleting index")
	indexDelete, err := global.ESClient.DeleteIndex(index).Do(context.Background())
	if err != nil {
		logrus.Error("Failed to delete index")
		logrus.Error(err.Error())
		return
	}
	if !indexDelete.Acknowledged {
		logrus.Error("Failed to delete index")
		return
	}
	logrus.Info("Index deleted successfully")
}

func CreateIndex(index string, mapping string) {
	if IndexExists(index) {
		RemoveIndex(index)
	}
	createIndex, err := global.ESClient.
		CreateIndex(index).
		BodyString(mapping).
		Do(context.Background())
	if err != nil {
		logrus.Error("Failed to create index")
		logrus.Error(err.Error())
		return
	}
	if !createIndex.Acknowledged {
		logrus.Error("Creation failed")
		return
	}
	logrus.Infof("Index %s created successfully", index)
}
func IndexExists(index string) bool {
	exists, err := global.ESClient.
		IndexExists(index).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return exists
	}
	return exists
}
