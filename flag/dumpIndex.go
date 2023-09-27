package flag

import (
	"blog_server/global"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type EsRawMessage struct {
	Row json.RawMessage `json:"row"`
	ID  string          `json:"id"`
}

type EsIndexResponse struct {
	Data    []EsRawMessage `json:"data"`
	Mapping interface{}    `json:"mapping"`
}

func DumpIndex(index string) {
	result, err := global.ESClient.
		Search(index).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	mapping, err := global.ESClient.GetMapping().Index("").Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}

	var jsonList []EsRawMessage
	for _, hit := range result.Hits.Hits {
		var jsonData EsRawMessage
		jsonData.Row = hit.Source
		jsonData.ID = hit.Id
		jsonList = append(jsonList, jsonData)
	}
	if len(jsonList) == 0 {
		logrus.Infof("No data under %s index", index)
		return
	}

	indexMapping, ok := mapping[index]
	if !ok {
		logrus.Errorf("No such index under this mapping")
		return
	}

	esIndexResponse := EsIndexResponse{
		Data:    jsonList,
		Mapping: indexMapping,
	}

	file, err := os.Create(fmt.Sprintf("%s.json", index))
	defer file.Close()
	if err != nil {
		logrus.Error(err)
		return
	}
	byteData, _ := json.Marshal(esIndexResponse)
	file.Write(byteData)
	logrus.Infof("Index %s successfully imported, total %d records", index, len(jsonList))

}
