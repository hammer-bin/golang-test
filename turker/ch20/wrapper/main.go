package main

import "fmt"

type Database interface {
	Get()
	Set()
}

type CDatabase struct {
}

func (c CDatabase) GetData() {
}

func (c CDatabase) SetData() {
}

type Wrapper struct {
	cdb CDatabase
}

func (w Wrapper) Get() {
	w.cdb.GetData()
}

func (w Wrapper) Set() {
	w.cdb.SetData()
}

func main() {
	var cdatabase CDatabase
	var database Database
	database = Wrapper{cdatabase}

	fmt.Println(database)

	database = Wrapper{}
}
