package struct2db
type StockInfo struct {
	Code string `db:"code"`;
	Name string `db:"name"`;
	Date string `db:"date"`;
	Open float64 `db:"open"`;
	High float64 `db:"high"`;
	Close float64 `db:"close"`;
	Low float64 `db:"low"`;
	Volume float64 `db:"volume"`;
	Price_change float64 `db:"price_change"`;
	P_change float64 `db:"p_change"`;
	Ma5 float64 `db:"ma5"`;
	Ma10 float64 `db:"ma10"`;
	Ma20 float64 `db:"ma20"`;
	V_ma5 float64 `db:"v_ma5"`;
	V_ma10 float64 `db:"v_ma10"`;
	V_ma20 float64 `db:"v_ma20"`;

}

type RealTimeData struct{
	Code string
	Name string
	Changepercent float64
	Trade float64
	Open float64
	High float64
	Low float64
	Close float64
	Volume float64
	Turnoverratio float64
	Amount float64
	Per float64
}
