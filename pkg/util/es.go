package util

import (
	"context"
	"github.com/golang/glog"
	"github.com/olivere/elastic/v7"
)

type Salary struct {
	id   string
	name string
}

type ESWorker struct {
	Client *elastic.Client
}

func (esw *ESWorker) CreateIndex() bool {
	ctx := context.Background()
	exist, err := esw.Client.IndexExists("salary").Do(ctx)
	if err != nil {
		glog.Error("retrieve index error, ", err)
		return false
	}
	if !exist {
		// create es index
		createIndex, err := esw.Client.CreateIndex("salary").Do(ctx)
		if err != nil {
			glog.Error("create es index error, ", err)
			return false
		}
		if createIndex.Acknowledged {
			glog.Info("create es index successfully !")
			return true
		}
	}
	return true
}

func (esw *ESWorker) InsertData() {
	data := Salary{
		id:   "123456",
		name: "first",
	}
	insert, err :=  esw.Client.Index().Index("salary").Id("111").BodyJson(data).Do(context.Background())
	if err != nil {
		glog.Error("inert data error, ", err)
	}
	glog.Info("insert successfully, ", insert.Index, insert.Id, insert.Type)
}