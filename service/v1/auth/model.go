package auth

type SigninIN struct {
	LoginType    string `json:"login_type" binding:"required,oneof=ADMIN WEB MOBILE"`
	Username     string `json:"username" binding:"required,max=100"`
	IdentifyCode string `json:"identify_code" binding:"required,max=100"`
}
