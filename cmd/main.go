package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/olivere/elastic/v7"
)

const (
	HOST = "http://192.168.234.4:9200/"
)

func main() {
	flag.Parse()
	flag.Set("alsologtostderr", "true")
	// if es is not a public network ip, please set sniff = false
	esClient, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(HOST))
	if err != nil {
		glog.Fatalln("error to create es client, ", err)
	}
	info, code, err := esClient.Ping(HOST).Do(context.Background())
	if err != nil {
		glog.Fatalln("error to ping HOST")
	}
	fmt.Println(code)
	glog.Info("Ping es successful " + info.Version.Number)
	defer glog.Flush()
}