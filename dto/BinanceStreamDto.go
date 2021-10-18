package dto

type BinanceStreamDto struct {
	Symbol string		`json:"s"`
	Kline KlineDto		`json:"k"`
}

type KlineDto struct {
	StartTime 		int64		`json:"t"`
	CloseTime 		int64		`json:"T"`
	Symbol 			string		`json:"s"`
	ClosePrice 		string		`json:"c"`
	Closed 			bool		`json:"x"`
}
