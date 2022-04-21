package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/aobt/sqlmapper"
	_ "github.com/go-sql-driver/mysql" //묵시적 사용
	"log"
	"os"
)

var table = "iaas_info"

type TableRow struct {
	No          int64  `sql:"no"`
	Iaas        string `sql:"iaas"`
	Credentials string `sql:"credentials"`
	Project     string `sql:"project"`
	Region      string `sql:"region"`
}

// Query by primary key (field[0])
func QueryByKey(ctx context.Context, tx *sql.Tx, db *sql.DB, fieldKey int64) (
	*TableRow, error) {

	var row TableRow
	row.No = fieldKey
	fm, err := sqlmapper.NewFieldsMap(table, &row)
	if err != nil {
		return nil, err
	}

	objptr, err := fm.SQLSelectByPriKey(ctx, tx, db)
	if err != nil {
		return nil, err
	}

	return objptr.(*TableRow), nil
}

func getDatabaseData() {
	db, err := sql.Open("mysql", "root:PaaS-TA@2022@(localhost:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	data, _ := QueryByKey(ctx, nil, db, 1)
	log.Printf("\nNo = %d\n IaaS = %s\n Credentials = %s\n Project = %s\n Region = %s\n ", data.No, data.Iaas, data.Credentials, data.Project, data.Region)

}
func main() {

	path, _ := os.Getwd()
	fmt.Println(path)
	file, err := os.Create("provider.tf")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer file.Close()
	str := `provider "google" {` + "\n\t" + `credentials = "${file("paasta-cp-3728429.json")}"` + "\n\t" + `project = "paasta-cp"` + "\n\t" + `region = "northamerica-northeast2"` + "\n}"
	fmt.Fprintln(file, str)

	getDatabaseData()
}
