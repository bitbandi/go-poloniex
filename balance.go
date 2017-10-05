package poloniex

type Balance struct {
	Available float64 `json:"available,string"`
	BtcValue  float64 `json:"btcValue,string"`
	OnOrders  float64 `json:"onOrders,string"`
}
