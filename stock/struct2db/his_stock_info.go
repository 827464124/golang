package struct2db
type StockInfo struct {
	Code string `db:"code"`;
	Name string `db:"name"`;
	Date string `db:"date"`;
	Open string `db:"open"`;
	High string `db:"high"`;
	Close string `db:"close"`;
	Low string `db:"low"`;
	Volume string `db:"volume"`;
	Price_change string `db:"price_change"`;
	P_change string `db:"p_change"`;
	Ma5 string `db:"ma5"`;
	Ma10 string `db:"ma10"`;
	Ma20 string `db:"ma20"`;
	V_ma5 string `db:"v_ma5"`;
	V_ma10 string `db:"v_ma10"`;
	V_ma20 string `db:"v_ma20"`;

}
