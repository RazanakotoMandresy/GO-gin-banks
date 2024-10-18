package user

type registerRequest struct {
	AppUserName       string
	Name              string
	FirstName         string
	Moneys            uint
	Password          string
	Date_de_naissance string
	Residance         string
	Email             string
}
type loginRequest struct {
	Email    string `json:"Email"`
	Password string `json:"password"`
}
type updateRequest struct {
	AppUserName string `json:"AppUserName"`
	Residance   string `json:"residance"`
}