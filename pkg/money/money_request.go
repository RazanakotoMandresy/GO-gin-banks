package money

type DepoRetraiReq struct {
	Value     int32  `json:"value"`
	Lieux     string `json:"place"`
	Passwords string `json:"password"`
}
type sendMoneyRequest struct {
	Value     int32  `json:"value"`
	Passwords string `json:"password"`
}
type historicRequest struct {
	Order string `json:"order_by"`
	Limit int    `json:"limit"`
}
