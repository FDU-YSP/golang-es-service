package main

import (
	"fduysp/go-es/pkg/util"
	"flag"
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
	esw := &util.ESWorker{
		Client: esClient,
	}
	flag := esw.CreateIndex("full-doc-test")
	if flag {
		text := "Caused by: org.springframework.beans.factory.BeanCreationException: Error creating bean with name " +
			"'mongoConnectionPoolTagsProvider' defined in class path resource [org/springframework/boot/actuate/autoconfigure/metrics/" +
			"mongo/MongoMetricsAutoConfiguration$MongoConnectionPoolMetricsConfiguration.class]: Post-processing of merged bean definition" +
			" failed; nested exception is java.lang.IllegalStateException: Failed to introspect Class [io.micrometer.core.instrument.binder" +
			".mongodb.DefaultMongoConnectionPoolTagsProvider] from ClassLoader [sun.misc.Launcher$AppClassLoader@18b4aac2]\n\tat org." +
			"springframework.beans.factory.support.AbstractAutowireCapableBeanFactory.doCreateBean(AbstractAutowireCapableBeanFactory.java:597)"
		esw.InsertData(text)
	}
	//info, code, err := esClient.Ping(HOST).Do(context.Background())
	//if err != nil {
	//	glog.Fatalln("error to ping HOST")
	//}
	//fmt.Println(code)
	//glog.Info("Ping es successful " + info.Version.Number)
	//
	//ver, _ := esClient.ElasticsearchVersion(HOST)
	//glog.Info("version: ", ver)

	defer glog.Flush()
}