package result

import (
	"../struct2db"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"time"
	"../conf"
)

type Status struct {
	lowestHighAvg bool
	highestLowAvg bool
	hasBuyIn bool
	value float64
	buyPrice float64
	salePrice float64
	number int
}
func ConnDB() *sql.DB {

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",conf.USERNAME,conf.PASSWORD,conf.NETWORK,conf.SERVER,conf.PORT,conf.DATABASE)
	DB,err := sql.Open("mysql",dsn)
	if err != nil{
		fmt.Printf("Open mysql failed,err:%v\n",err)
		return nil
	}
	DB.SetConnMaxLifetime(100*time.Second)  //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)//设置最大连接数
	DB.SetMaxIdleConns(16) //设置闲置连接数
	return DB

}
func QueryData(db *sql.DB,code string)  {
	var skt = new(struct2db.StockInfo)
	rows, err := db.Query("select date,code,name,open,high,low,close ,ma20 ,ma10 from his_stock_info where code = ? order by date", code)
	defer func() {
		if rows != nil {
			rows.Close() //可以关闭掉未scan连接一直占用

		}
	}()
	if err != nil {
		fmt.Printf("Query failed,err:%v", err)
		return
	}
	var ST Status
	ST.number = 100
	for rows.Next() {
		err = rows.Scan(&skt.Date,&skt.Code,&skt.Name,&skt.Open,&skt.High,&skt.Low,&skt.Close,&skt.Ma20,&skt.Ma10) //不scan会导致连接不释放
		if err != nil {
			fmt.Printf("Scan failed,err:%v", err)
			return
		}
		if skt.Low > skt.Ma20 {
			ST.lowestHighAvg  = true
		}else  {
			ST.lowestHighAvg = false
		}
		if skt.High < skt.Ma20{
			ST.highestLowAvg = true
		}else{
			ST.highestLowAvg = false
		}
		if ST.lowestHighAvg && ! ST.hasBuyIn {
			ST.hasBuyIn = true
			ST.buyPrice = skt.Close
			fmt.Println("buy in ",skt.Date,"  " ,ST.buyPrice)
		}
		if !ST.lowestHighAvg && ST.hasBuyIn{
			ST.hasBuyIn = false
			ST.salePrice = skt.Ma20
			ST.value += (ST.salePrice - ST.buyPrice) * 100 *(1-0.0002)
			fmt.Println("sale out ",skt.Date,"  " ,ST.salePrice)
		}
		//fmt.Print(skt)

	}
	fmt.Println(ST.value)
}



