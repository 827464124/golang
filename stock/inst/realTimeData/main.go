package main

import (
	"../../result"
	"../../struct2db"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)
func main(){
	db := result.ConnDB()
	QueryData(db,"600070")
}
func QueryData(db *sql.DB,code string)  {
	var rtd = new(struct2db.RealTimeData)
	rows, err := db.Query("select code,name,open,high,low,close ,turnoverratio from real_time_data where code = ? order by date", code)
	defer func() {
		if rows != nil {
			rows.Close() //可以关闭掉未scan连接一直占用

		}
	}()
	if err != nil {
		fmt.Printf("Query failed,err:%v", err)
		return
	}


	for rows.Next() {
		err = rows.Scan(&rtd.Code,&rtd.Name,&rtd.Open,&rtd.High,&rtd.Low,&rtd.Turnoverratio) //不scan会导致连接不释放
		if err != nil {
			fmt.Printf("Scan failed,err:%v", err)
			return
		}
		fmt.Println("名称： ",rtd.Name)
		fmt.Println("开盘价： ",rtd.Open)
		fmt.Println("收盘价： ",rtd.Close)
		fmt.Println("最高价： ",rtd.High)
		fmt.Println("最低价： ",rtd.Low)
		fmt.Println("换手率： ",rtd.Turnoverratio)

		//fmt.Print(skt)

	}

}