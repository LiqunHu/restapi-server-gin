package auth

type SigninIN struct {
	LoginType    string `json:"LoginType" binding:"required,oneof=WEB MOBILE"`
	Username     string `json:"Username" binding:"required,max=100"`
	IdentifyCode string `json:"IdentifyCode" binding:"required,max=100"`
	MagicNo      string `json:"MagicNo" binding:"required,max=100"`
}
