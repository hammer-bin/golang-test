package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang-test/test01/example01/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var databaseAddr string
var databasePort string
var databaseName string
var databaseId string
var databasePassword string
var db *gorm.DB

func init() {
	db = connect()
}

func getDataSource() string {
	databaseAddr = common.ConfInfo["database.mariadb.address"]
	databasePort = common.ConfInfo["database.mariadb.port"]
	databaseName = common.ConfInfo["database.mariadb.name"]
	databaseId = common.ConfInfo["database.mariadb.id"]
	databasePassword = common.ConfInfo["database.mariadb.password"]

	dataSource := fmt.Sprint(databaseId, ":", databasePassword, "@(", databaseAddr, ":", databasePort, ")/", databaseName, "?parseTime=true")
	fmt.Println("dataSource :: ", dataSource)
	return dataSource
}

func connect() *gorm.DB {
	dsn := getDataSource()
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic("Db 연결에 실패하였습니다.")
	}

	return db
}

type result struct {
	Name     string
	cpuRatio float64
}

type CpClusters struct {
	ClusterId string `gorm:"column:cluster_id"`
	Name      string `gorm:"column:name"`
}

type CpMetricClusterStatus struct {
	ClusterId string  `gorm:"column:cluster_id"`
	cpuRatio  float64 `gorm:"column:cpu_ratio"`
}

func main() {

	//db.Model(&CpClusters{}).Select("cp_clusters.name, cp_metric_cluster_status.cpu_ratio").
	//	Joins("left join cp_metric_cluster_status on cp_metric_cluster_status.cluster_id = cp_clusters.cluster_id").Scan(&result{})
	//// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id
	//
	//
	//rows, err := db.Table("cp_clusters").Select("cp_clusters.name, cp_metric_cluster_status.cpu_ratio").
	//	Joins("left join cp_metric_cluster_status on cp_metric_cluster_status.cluster_id = cp_clusters.cluster_id").Rows()
	//
	//if err != nil {
	//
	//}

	var cluster CpClusters

	//row1 := db.Joins("cp_metric_cluster_status").Find(&cluster)

	db.Preload("cp_metric_cluster_status").First(&cluster)
	data, _ := json.Marshal(cluster)
	var prettyJSON bytes.Buffer
	_ = json.Indent(&prettyJSON, data, "", "\t")
	fmt.Println(prettyJSON.String())

	//fmt.Println(row1)

}
