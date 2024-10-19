package user

type registerRequest struct {
	AppUserName string
	Name        string
	FirstName   string
	Password    string
	BirthDate   string
	Residance   string
	Email       string
}
type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type updateRequest struct {
	AppUserName string `json:"AppUserName"`
	Residance   string `json:"residance"`
}
type SettingReq struct {
	RemoveAllEp     bool   `json:"rmEpargne"`
	DeleteMyAccount bool   `json:"rmAccount"`
	BlockAccount    string `json:"blockAcc"`
	UnBlockAccount  string `json:"unblockAcc"`
}
