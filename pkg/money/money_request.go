package money

type topTrans struct {
	SentTo     string `json:"sentTo"`
	Totals     int    `json:"sommeTrans"`
	UserName   string `json:"userName"`
	ImageSento string `json:"SentToImg"`
	ImgSender  string `json:"SentByImg"`
}
type DepoRetraiReq struct {
	Value     int32  `json:"value"`
	Lieux     string `json:"lieux"`
	Passwords string `json:"passwords"`
}
type sendMoneyRequest struct {
	Value int32 `json:"value"`
}
