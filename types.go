package moex_iss

type MoexCandle struct {
	Open float64
	Close float64
	High float64
	Low float64
	Volume float64
	Date string
}

type MoexCandles struct {
	Data []([]interface{}) `json:"data"`
}

type MoexCandleResponse struct {
	Candles MoexCandles `json:"candles"`
}

