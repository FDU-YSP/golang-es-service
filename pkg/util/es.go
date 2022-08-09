package util

import (
	"context"
	"github.com/golang/glog"
	"github.com/google/uuid"
	"github.com/olivere/elastic/v7"
)

type Salary struct {
	ID             string
	Employee       EmployeeData
	Employer       EmployerData
	ActuallyIssued int
}

type EmployeeData struct {
	All            int
	BaseSalary     int
	MealSupplement int
	Other          map[string]int
	Deduct         DeductData
}

type DeductData struct {
	All                         int
	Pension_ee                  int
	MedicalInsurance_ee         int
	UnemploymentInsurance_ee    int
	HousingFound_ee             int
	SupplementaryHousingFund_ee int
	Tax                         int
	Other                       map[string]int
}

type EmployerData struct {
	All                         int
	BaseSalary                  int
	Pension_er                  int
	MedicalInsurance_er         int
	UnemploymentInsurance_er    int
	HousingFound_er             int
	SupplementaryHousingFund_er int
	WorkInjuryInsurance_er      int
	MaternityInsurance_er       int
}


type ESWorker struct {
	Client *elastic.Client
}

// CreateIndex create a index
func (esw *ESWorker) CreateIndex(indexName string) bool {
	ctx := context.Background()
	exist, err := esw.Client.IndexExists(indexName).Do(ctx)
	if err != nil {
		glog.Error("retrieve index error, ", err)
		return false
	}
	if !exist {
		createIndex, err := esw.Client.CreateIndex(indexName).Do(ctx)
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

// InsertData insert data to ES
func (esw *ESWorker) InsertData(text string) {
	uid := uuid.New().String()

	data := Salary{
		//id:   uid,
		Text: text,
	}
	insert, err :=  esw.Client.Index().Index("salary").Id(uid).BodyJson(data).Do(context.Background())
	if err != nil {
		glog.Error("inert data error, ", err)
	}
	glog.Info("insert successfully, ", insert.Index, insert.Id, insert.Type)
}

func (esw *ESWorker) DeleteData() {

}

func (esw *ESWorker) QueryData() {
	elastic.NewMatchQuery("keyword", [...]("Caused", "by"))
}

// DeleteIndex delete an index
func (esw *ESWorker) DeleteIndex(indexName string) {
	deleteIndex, err := esw.Client.DeleteIndex(indexName).Do(context.Background())
	if err != nil {
		glog.Error("error to delete index, ", err)
	}
	if deleteIndex.Acknowledged {
		glog.Info("delete index successfully !")
	}
}