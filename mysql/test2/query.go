package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type cp struct {
	clusterId    string
	name         string
	providerType string
}

func main() {
	db, err := sql.Open("mysql", "terraman:Paasta!2022@(15.164.195.107:31306)/cpdev?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//var name string

	//err = db.QueryRow("SELECT name FROM test.user WHERE no = 1").Scan(&name)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(name)

	rows, err := db.Query("SELECT a.cluster_id, a.name, a.provider_type FROM cp_clusters a, cp_metric_cluster_status b\n WHERE a.cluster_id = b.cluster_id\n   AND a.status = ?", "A")
	if err != nil {
		log.Fatal(err)
	}

	//var clusterId string
	//var name string
	//var providerType string

	var cpa cp

	for rows.Next() {
		err := rows.Scan(&cpa)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(cpa.clusterId, cpa.name, cpa.clusterId)
	}

	result, err := db.Exec("INSERT INTO test.user (name, age) VALUES (?, ?)", "jimmy2", 32)
	if err != nil {
		log.Fatal(err)
	}

	n, err := result.RowsAffected()
	if n == 1 {
		n, _ := result.LastInsertId()
		fmt.Println("1 row inserted. no : ", n)
	}

	stmt, _ := db.Prepare("UPDATE test.user SET name=?, age=? WHERE no=?")
	defer stmt.Close()
	_, _ = stmt.Exec("Tom", "34", "4")

	//_ = db.QueryRow("SELECT no, name, age FROM test.user WHERE no=?", 4).Scan(&no, &name2, &age)
	//fmt.Println(no, name2, age)

}
