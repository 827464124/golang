package main

import "./result"
func main(){
	var db  = result.ConnDB()
	result.QueryData(db,"000001")
}
