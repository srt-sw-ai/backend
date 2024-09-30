package dto

type SignupDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	NickName string `json:"nickName"`
}

type AuthDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EditProfileDto struct {
	NickName string `json:"nickName"`
	Email    string `json:"email"`
}
