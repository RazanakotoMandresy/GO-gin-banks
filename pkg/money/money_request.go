package money

type DepoRetraiReq struct {
	Value     int32  `json:"value"`
	Lieux     string `json:"lieux"`
	Passwords string `json:"passwords"`
}
type sendMoneyRequest struct {
	Value     int32  `json:"value"`
	Passwords string `json:"password"`
}
